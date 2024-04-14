package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"github.com/gorilla/mux"
	"github.com/madsportslab/nbalake"
	"github.com/stephenhu/stats"
)


func storeJson(bucket string, fn string, data interface{}) {

	j, err := json.Marshal(data)

	if err != nil {
		log.Println(err)
	} else {
		nbalake.Put(bucket, fn, j)
	}

} // storeJson


func storeGame(bucket string, id string) {

	fn := fmt.Sprintf("%s%s", id, EXT_JSON)

	score := stats.NbaGetBoxscore(id)

	if score.Game.Status == GAME_FINAL {

		if !nbalake.Exists(bucket, fn) {

			log.Printf("Storing object %s into %s\n", fn, bucket)

			storeJson(bucket, fn, score)

			jobMu.Lock()
			jobState[jobName(bucket, id, JOB_BOXSCORE)] = true
			jobMu.Unlock()

		}

		storePlays(bucket, id)

	}

} // storeBoxscore


func storePlays(bucket string, id string) {

	fn := fmt.Sprintf("%s%s%s", id, EXT_PBP, EXT_JSON)

	if !nbalake.Exists(bucket, fn) {
			
		plays := stats.NbaGetPlays(id)

		log.Printf("Storing object %s into %s\n", fn, bucket)

		storeJson(bucket, fn, plays)

		jobMu.Lock()
		jobState[jobName(bucket, id, JOB_PLAYBYPLAY)] = true
		jobMu.Unlock()

	}

} // storePlays



func getGamesToDownload() []string {

	games := []string{}

	for _, dates := range schedule.LeagueSchedule.GameDates {

		if !stats.IsFutureGame(dates.GameDate) {

			for _, game := range dates.Games {
				games = append(games, game.ID)
			}

		}

	}

	return games

} // getGamesToDownload


func pullGames(bucket string, games []string) {

	for _, g := range games {

		jobState[jobName(bucket, g, JOB_BOXSCORE)] = false
		jobState[jobName(bucket, g, JOB_PLAYBYPLAY)] = false

		downloads <- Download{
			GameID: g,
			Bucket: bucket,
		}

	}

} // pullGames


func resumeGamesDownload() {

	games := getGamesToDownload()

	pullGames(rawBucket, games)

} // resumeGamesDownload


func gameHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

		ids := r.FormValue(FORM_PARAM_IDS)
		
		games := parseList(ids, DELIMITER_COMMA)

		pullGames(rawBucket, games)


	case http.MethodPut:

		resumeGamesDownload()

	case http.MethodGet:

		j, err := json.Marshal(jobState)

		if err != nil {
			log.Println(err)
		} else {
			w.Write(j)
		}

	default:
	}

} // gameHandler
