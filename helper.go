package main

import (
	"strings"
	"time"
)


func parseList(s string, d string) []string {

	var ret []string

	if len(s) == 0 {
		return nil
	}

	tokens := strings.Split(s, d)

	for _, t := range tokens {
		ret = append(ret, strings.TrimSpace(t))
	}

	return ret

} // parseList


func percentage(attempted int, made int) float32 {

	if attempted == 0 {
		return 0.0
	} else {
		return float32(made)/float32(attempted)
	}

} // percentage


func perGamePercentage(games int, d int) float32 {

	if games == 0 {
		return 0.0
	} else {
		return float32(d)/float32(games)
	}

} // perGamePercentage


func perGamePercentageFp(games int, d float32) float32 {

	if games == 0 {
		return 0.0
	} else {
		return d/float32(games)
	}

} // perGamePercentageFp


func playedGame(mins int) int {

	if mins > 0 {
		return 1
	} else {
		return 0
	}

} // playedGame


func getNowStamp() string {
  
	now := time.Now()
	
	return now.Format(NBA_DATE_FORMAT)

} // getNowStamp
