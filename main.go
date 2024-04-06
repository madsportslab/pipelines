package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/madsportslab/nbalake"
	"github.com/stephenhu/stats"
)

var rawBucket, analyticsBucket string

var syncTicker *time.Ticker


func initRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/analytics", analyticsHandler)
	router.HandleFunc("/api/v1/analytics/{year:[0-9]+}", analyticsHandler)
	router.HandleFunc("/api/v1/games", gameHandler)
	router.HandleFunc("/api/v1/legacy/{year:[0-9]+}", legacyHandler)
	router.HandleFunc("/api/v1/schedule", scheduleHandler)
	router.HandleFunc("/api/v1/version", versionHandler)

	return router

} // initRouter


func initLake() {

	nbalake.ConnectionNew()

	currentSeason = stats.GetCurrentSeason()

	nbalake.InitBuckets([]string{"2023.nba.raw", "2023.nba.analytics"})

	rawBucket = nbalake.BucketName(currentSeason,
		nbalake.BUCKET_RAW)

	analyticsBucket = nbalake.BucketName(currentSeason,
		nbalake.BUCKET_ANALYTICS)

	loadSchedule()

} // initLake


func initSyncJob() {

	syncTicker := time.NewTicker(
		APP_SYNC_PERIOD * time.Hour)
	
	done := make(chan bool)

	go func() {
		
		for {

			select {
			case <-done:
				return
			case <-syncTicker.C:

				resumeGamesDownload()

				generateData()

			}

		}

	}()

} // initSyncJob


func main() {

	fmt.Printf("Starting %s v%s...\n", APP_NAME, APP_VERSION)

	initLake()

	initSyncJob()

	log.Fatal(http.ListenAndServe(":8686", initRouter()))

} // main
