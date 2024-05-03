package main

import (
	"fmt"
	"testing"
)

func CheckList(list []string, count int, length int) {

	if len(list) != count {
		fmt.Printf("Number of items in list incorrect, received %d " +
		  "should have %d\n", len(list), count)
	}

	for _, v := range list {

		if len(v) != length {
			fmt.Printf("Parsed item contains invalid characters, " +
			  "should be %d, but found %d", length, len(v))
		}

	}

} // CheckList


func TestParseList(t *testing.T) {

	ids := parseList("0042000402, 0042000403, 0042000404, 0042000405",
	  DELIMITER_COMMA)

	CheckList(ids, 4, 10)


} // TestParseList


func TestParseListMixed(t *testing.T) {

	ids := parseList("0042000402,0042000403,  0042000404 , 0042000405",
	  DELIMITER_COMMA)

	CheckList(ids, 4, 10)


} // TestParseListMixed


func TestStripGameIdValid(t *testing.T) {

	id := stripGameId("0042000402.json")

	if id != "0042000402" {
		t.Error("game id invalid")
	}

} // TestStripGameIdValid


func TestStripGameIdAlpha(t *testing.T) {

	id := stripGameId("aabcd")

	if id != "aabcd" {
		t.Error("game id invalid, should return empty string")
	}

} // TestStripGameIdAlpha


func TestStripGameIdEmpty(t *testing.T) {

	id := stripGameId(STR_EMPTY)

	if id != STR_EMPTY {
		t.Log(id)
		t.Error("game id invalid, should return empty string")
	}

} // TestStripGameIdEmpty
