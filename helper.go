package main

import (
	"strings"
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
