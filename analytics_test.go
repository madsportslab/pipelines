package main

import (
	"testing"
)


func TestIsGameTypePlayoff(t *testing.T) {

	if !isGameType(GAME_TYPE_PLAYOFF) {
		t.Error("Should be a valid game type")
	}

} // TestIsGameTypePlayoff


func TestIsGameTypePreseason(t *testing.T) {

	if !isGameType(GAME_TYPE_PRESEASON) {
		t.Error("Should be a valid game type")
	}

} // TestIsGameTypePreseason


func TestIsGameTypeInvalid(t *testing.T) {

	if isGameType(606) {
		t.Error("Should be a valid game type")
	}

} // TestIsGameTypeInvalid
