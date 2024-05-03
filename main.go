package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/madsportslab/nbalake"
	"github.com/stephenhu/stats"
)

var rawBucket, analyticsBucket string


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

	rawBucket = nbalake.BucketName(currentSeason,
		nbalake.BUCKET_RAW)

	analyticsBucket = nbalake.BucketName(currentSeason,
		nbalake.BUCKET_ANALYTICS)

	nbalake.InitBuckets([]string{
		rawBucket,
		analyticsBucket})

	loadSchedule()

	initGameMap()

} // initLake


func main() {

	fmt.Printf("Starting %s v%s...\n", APP_NAME, APP_VERSION)

	initLake()

	initSyncJob()

	initJobListener()

	go resumeGamesDownload()

	log.Fatal(http.ListenAndServe(":8686", initRouter()))

} // main
