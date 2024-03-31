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


func pullGames(ids string) {

	games := parseList(ids, DELIMITER_COMMA)

	raw := nbalake.BucketName(currentSeason,
		nbalake.BUCKET_RAW)

	for _, g := range games {

		fn := fmt.Sprintf("%s%s", g, EXT_JSON)

		if !nbalake.Exists(raw, fn) {

			score := stats.NbaGetBoxscore(g)

			if score.Game.Status == GAME_FINAL {

				j, err := json.Marshal(score)

				if err != nil {
					log.Println(err)
				} else {
					nbalake.Put(raw, fn, j)
				}

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
		
		pullGames(ids)


	case http.MethodGet:


	default:
	}

} // gameHandler
