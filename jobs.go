package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Download struct {
	GameID					string				`json:"gameId"`
	Resource        int           `json:"resource"`
	Bucket					string				`json:"bucket"`
	Created         string        `json:"created"`
	Completed      	string        `json:"completed"`
}


type Etl struct {
	SeasonID				string				`json:"seasonId"`
	Bucket					string				`json:"bucket"`
	Name						string				`json:"name"`
}


type State struct {
	ID              string        `json:"id"`
	Created         string        `json:"created"`
	Completed      	string        `json:"completed"`
}


var syncTicker *time.Ticker
var downloads chan Download
var analytics chan Etl
var jobState map[string] *State  // raw|analytics:gameId|seasonId:box|pbp
var jobMu sync.Mutex


func jobName(b string, id string, st string) string {
	return fmt.Sprintf("%s:%s:%s", b, id, st)
} // jobName


func initJobListener() {

	fmt.Println("\t+ initiating job listener...")

	downloads = make(chan Download)
	analytics = make(chan Etl)
	jobState	= make(map[string] *State)

	go func() {

		for {

			select {
			case job1 := <-downloads:

				if job1.Resource == RESOURCE_BOXSCORE {
					storeBoxscore(job1.Bucket, job1.GameID)
				} else if job1.Resource == RESOURCE_PLAYBYPLAY {
					storePlays(job1.Bucket, job1.GameID)
				} else {
					log.Println(job1.Resource)
					log.Println("Error: unknown download resource type")
				}

			case <-analytics:
				
				generateData()

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
