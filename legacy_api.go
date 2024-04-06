package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/madsportslab/nbalake"
	"github.com/stephenhu/stats"
)


var validSeasons = map[string] bool{
	"2015": true,
	"2016": true,
	"2017": true,
	"2018": true,
	"2019": true,
	"2020": true,
	"2021": true,
	"2022": true,
	"2023": true,
}


func getLegacyGames(season string) []string {

	var games []string

	_, ok := validSeasons[season]

	if ok {

		sched := stats.NbaGetLegacySchedule(season)
	
		for _, month := range sched.Lscd {

			for _, g := range month.Mscd.Games {
				games = append(games, g.ID)
			}

		}
	
	}

	return games

} // getLegacyGames


func legacyHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

		vars := mux.Vars(r)

		year := vars[FORM_PARAM_YEAR]

		_, ok := validSeasons[year]

		if ok {
		
			games := getLegacyGames(year)

			raw := nbalake.BucketName(year, nbalake.BUCKET_RAW)
			
			analytics :=	nbalake.BucketName(year,
				nbalake.BUCKET_ANALYTICS)

			nbalake.InitBuckets([]string{raw, analytics})
	
			pullGames(raw, games)
			
		} else {
			
			log.Println("Error: invalid season " + year +
			  ".  Valid seasons: 2015-2023")

		}

		
	case http.MethodPut:


	case http.MethodGet:


	default:
	}

} // legacyHandler
