package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Download struct {
	GameID					string				`json:"gameId"`
	Bucket					string				`json:"bucket"`
}


type Etl struct {
	SeasonID				string				`json:"seasonId"`
	Bucket					string				`json:"bucket"`
	Name						string				`json:"name"`
}


var syncTicker *time.Ticker
var downloads chan Download
var analytics chan Etl
var jobState map[string] bool  // raw|analytics:gameId|seasonId:box|pbp
var jobMu sync.Mutex


func jobName(b string, id string, st string) string {
	return fmt.Sprintf("%s:%s:%s", b, id, st)
} // jobName


func initJobListener() {

	fmt.Println("\t+ initiating job listener...")

	downloads = make(chan Download)
	analytics = make(chan Etl)
	jobState	= make(map[string] bool)

	go func() {

		for {

			select {
			case job1 := <-downloads:
				storeGame(job1.Bucket, job1.GameID)
			case job2 := <-analytics:
				log.Println(job2)
			}
		
		}

	}()


} // initJobListener


func initSyncJob() {

	fmt.Printf("\t+ initiating sync job, every %dh...\n", APP_SYNC_PERIOD)
	
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
