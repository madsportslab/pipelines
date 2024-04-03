package main

import (
	"encoding/json"
	"log"
	"net/http"

	//"github.com/gorilla/mux"
	"github.com/madsportslab/nbalake"
	"github.com/stephenhu/stats"
)

var schedule stats.NbaSchedule


func loadSchedule() {

	getSchedule()
	
	j := nbalake.Get(rawBucket, SCHEDULE_JSON)

	err := json.Unmarshal(j, &schedule)

	if err != nil {
		log.Println(err)
	}

} // loadSchedule


func getSchedule() {

	if !nbalake.Exists(rawBucket, SCHEDULE_JSON) {
			
		j := stats.NbaGetScheduleJson()

		nbalake.Put(rawBucket, SCHEDULE_JSON, j)

	}

} // getSchedule


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


func scheduleHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		loadSchedule()

		j, err := json.MarshalIndent(schedule, STR_EMPTY, STR_TAB)

		if err != nil {
			log.Println(err)
		} else {
			w.Write(j)
		}

	default:
	}

} // scheduleHandler
