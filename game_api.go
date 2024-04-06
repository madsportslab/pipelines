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


func pullGames(bucket string, games []string) {

	// check valid year

	for _, g := range games {

		fn := fmt.Sprintf("%s%s", g, EXT_JSON)
    fp := fmt.Sprintf("%s%s%s", g, EXT_PBP, EXT_JSON)

		score := stats.NbaGetBoxscore(g)

		if score.Game.Status == GAME_FINAL {

			if !nbalake.Exists(bucket, fn) {

				log.Printf("Storing object %s into %s\n", fn, bucket)

				storeJson(bucket, fn, score)

			}
	
			if !nbalake.Exists(bucket, fp) {
				
				plays := stats.NbaGetPlays(g)
	
				log.Printf("Storing object %s into %s\n", fp, bucket)

				storeJson(bucket, fp, plays)
		
			}
	
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


	default:
	}

} // gameHandler
