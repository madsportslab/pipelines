package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/madsportslab/nbalake"
	"github.com/stephenhu/stats"
)


func initRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/games", gameHandler)
	router.HandleFunc("/api/v1/version", versionHandler)

	return router

} // initRouter


func initLake() {

	nbalake.ConnectionNew()

	currentSeason = stats.GetCurrentSeason()

	nbalake.InitBuckets([]string{"2023.nba.raw", "2023.nba.analytics"})

} // initLake


func main() {

	fmt.Printf("Starting %s v%s...\n", APP_NAME, APP_VERSION)

	initLake()

	log.Fatal(http.ListenAndServe(":8686", initRouter()))

} // main
