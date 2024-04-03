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


func storeJson(fn string, data interface{}) {

	j, err := json.Marshal(data)

	if err != nil {
		log.Println(err)
	} else {
		nbalake.Put(rawBucket, fn, j)
	}

} // storeJson


func pullGames(games []string) {

	for _, g := range games {

		fn := fmt.Sprintf("%s%s", g, EXT_JSON)
    fp := fmt.Sprintf("%s%s%s", g, EXT_PBP, EXT_JSON)

		score := stats.NbaGetBoxscore(g)

		if score.Game.Status == GAME_FINAL {

			if !nbalake.Exists(rawBucket, fn) {
				storeJson(fn, score)
			}
	
			if !nbalake.Exists(rawBucket, fp) {
				
				plays := stats.NbaGetPlays(g)
	
				storeJson(fp, plays)
		
			}
	
		}

	}

} // pullGames


func gameHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

		ids := r.FormValue(FORM_PARAM_IDS)

		// r.FormValue(FORM_PARAM_SEASON)
		// r.FormValue(FORM_PARAM_LEAGUE)
		
		games := parseList(ids, DELIMITER_COMMA)

		pullGames(games)


	case http.MethodPut:

		games := getGamesToDownload()

		log.Println(games)

		pullGames(games)

	case http.MethodGet:


	default:
	}

} // gameHandler
