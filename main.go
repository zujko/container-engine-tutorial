package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Time struct {
	Timezone string    `json:"timezone"`
	Time     time.Time `json:"time,string"`
}

func main() {
	r := httprouter.New()

	r.GET("/time/:continent/:place", GetTime)
	r.GET("/time", GetTime)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func GetTime(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	country := p.ByName("continent")
	place := p.ByName("place")
	var currentTime time.Time
	tz := fmt.Sprintf("%s/%s", country, place)

	if country == "" || place == "" {
		currentTime = time.Now().UTC()
		tz = "UTC"
	} else {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		currentTime = time.Now().In(loc)
	}
	w.Header().Set("Content-Type", "application/json")

	timeResponse, err := json.Marshal(&Time{tz, currentTime})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(timeResponse)
}
