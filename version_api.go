package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func versionHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		s := struct {
			Version string
		}{
			Version: APP_VERSION,
		}

		j, err := json.Marshal(s)

		if err != nil {
			log.Println(err)
		} else {
			w.Write(j)
		}

	default:
	}
	
} // versionHandler
