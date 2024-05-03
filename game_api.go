package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	//"github.com/gorilla/mux"
	"github.com/madsportslab/nbalake"
	"github.com/stephenhu/stats"
)

type GameDownload struct {
  GameDate        string      `json:"gameDate"` 		// YYYYMMDD
	Downloaded    	string      `json:"downloaded"`
	GameType        int         `json:"gameType"`
	Boxscore				bool				`json:"boxscore"`
	PlayByPlay      bool        `json:"playByPlay"`
	Final        		bool				`json:"final"`
}


var (
	gamesMap				map[string] *GameDownload
)


func gameFinal(g stats.NbaGame) bool {

	if strings.Contains(g.Status, GAME_STATUS_FINAL) {
		return true
	} else {
		return false
	}

} // gameFinal


func gameType(g stats.NbaGame) int {

  if g.WeekName == STR_EMPTY && g.GameLabel == GAME_LABEL_PRESEASON {
		return GAME_TYPE_PRESEASON
	} else if g.WeekNumber != 0 && g.WeekName != STR_EMPTY &&
	  g.GameLabel == GAME_LABEL_INSEASON {
		return GAME_TYPE_INSEASON
	} else if g.WeekNumber != 0 && g.WeekName != STR_EMPTY &&
	  (g.GameLabel == STR_EMPTY || strings.Contains(
		g.GameLabel, GAME_LABEL_INTL)) {
		return GAME_TYPE_REGULAR
	} else if g.WeekNumber != 0 && g.WeekName == WEEK_NAME_ALLSTAR &&
	  strings.Contains(g.GameLabel, GAME_LABEL_RISING_STARS) {
		return GAME_TYPE_RISING_STARS
  } else if g.WeekNumber != 0 && g.WeekName == WEEK_NAME_ALLSTAR &&
	  g.GameLabel == GAME_LABEL_ALLSTAR {
		return GAME_TYPE_ALLSTAR
	} else if g.WeekName == STR_EMPTY && g.WeekNumber == 0 &&
	  strings.Contains(g.GameLabel, GAME_LABEL_PLAYIN) {
		return GAME_TYPE_PLAYIN
	} else if g.WeekName == STR_EMPTY && g.WeekNumber == 0 &&
	  g.GameLabel != STR_EMPTY {
		return GAME_TYPE_PLAYOFF
	}

	return GAME_TYPE_UNKNOWN

} // gameType


func isGameType(g int) bool {

	switch(g) {
	case GAME_TYPE_PRESEASON:
		return true
	case GAME_TYPE_REGULAR:
		return true
	case GAME_TYPE_INSEASON:
		return true
	case GAME_TYPE_ALLSTAR:
		return true
	case GAME_TYPE_PLAYIN:
		return true
	case GAME_TYPE_PLAYOFF:
		return true
	default:
		return false
	}

} // isGameType


func storeJson(bucket string, fn string, data interface{}) {

	j, err := json.Marshal(data)

	if err != nil {
		log.Println(err)
	} else {
		nbalake.Put(bucket, fn, j)
	}

} // storeJson


func storeBoxscore(bucket string, id string) {

	fn := fmt.Sprintf("%s%s", id, EXT_JSON)

	score := stats.NbaGetBoxscore(id)

	if gamesMap[id].Final {

		if !nbalake.Exists(bucket, fn) {

			log.Printf("Storing object %s into %s\n", fn, bucket)

			storeJson(bucket, fn, score)

		}

		jobMu.Lock()

		//jobState[jobName(bucket, id,
		//	JOB_BOXSCORE)].Completed = time.Now().String()

		//gamesMap[id].Downloaded
		gamesMap[id].Boxscore = true

		jobMu.Unlock()

		//storePlays(bucket, id)

	}

} // storeBoxscore


func storePlays(bucket string, id string) {

	fn := fmt.Sprintf("%s%s%s", id, EXT_PBP, EXT_JSON)

	if !nbalake.Exists(bucket, fn) {
			
		plays := stats.NbaGetPlays(id)

		log.Printf("Storing object %s into %s\n", fn, bucket)

		storeJson(bucket, fn, plays)

	}

	jobMu.Lock()

	//jobState[jobName(bucket, id,
	//	JOB_PLAYBYPLAY)].Completed = time.Now().String()
	
	gamesMap[id].PlayByPlay = true
	gamesMap[id].Downloaded =
	  time.Now().Format(stats.DATE_FORMAT)
	
	jobMu.Unlock()

} // storePlays


func getGamesToDownload() []string {

	games := []string{}

	for id, info := range gamesMap {

		if (!info.Boxscore || !info.PlayByPlay) &&
		  !stats.IsAfterNow(info.GameDate) {
			games = append(games, id)
		}

	}

	log.Println(games)

	return games

} // getGamesToDownload


func pullGames(bucket string, games []string) {

	for _, g := range games {

		if gamesMap[g].Final {

			downloads <- Download{
				GameID: g,
				Resource: RESOURCE_BOXSCORE,
				Bucket: bucket,
			}
	
			downloads <- Download{
				GameID: g,
				Resource: RESOURCE_PLAYBYPLAY,
				Bucket: bucket,
			}
	
		}
	
	}

	analytics <- Etl{
		SeasonID: currentSeason,
		Bucket: analyticsBucket,
	}

} // pullGames


func resumeGamesDownload() {

	games := getGamesToDownload()

	pullGames(rawBucket, games)

} // resumeGamesDownload


func compareDownloaded(b string) {

	objects := nbalake.List(b)

	for o := range objects {
		
		if o.Key == SCHEDULE_JSON {
			continue
		}

		id := stripGameId(o.Key)

		gd, ok := gamesMap[id]

		if ok {

			if strings.Contains(o.Key, EXT_PBP) {
				gd.PlayByPlay = true
			} else {
				gd.Boxscore = true
			}
	
			d := o.LastModified.Format(stats.DATE_FORMAT)
	
			gd.Downloaded = d
	
		}

	}

} // compareDownloaded


func initGameMap() {

	gamesMap = make(map[string]*GameDownload)

	for _, dates := range schedule.LeagueSchedule.GameDates {

		gds := stats.GameDateToString(dates.GameDate)

		for _, game := range dates.Games {

			gamesMap[game.ID] = &GameDownload{
				GameDate:  gds,
				GameType: gameType(game),
				Final: gameFinal(game),
			}

		}

	}
	
	compareDownloaded(rawBucket)

} // initGameMap


func gameHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

		ids := r.FormValue(FORM_PARAM_IDS)
		
		games := parseList(ids, DELIMITER_COMMA)

		pullGames(rawBucket, games)


	case http.MethodPut:

		resumeGamesDownload()

	case http.MethodGet:

		initGameMap()

		j, err := json.Marshal(gamesMap)

		if err != nil {
			log.Println(err)
		} else {
			w.Write(j)
		}

	default:
	}

} // gameHandler
