package api

import (
	"IMT2681-assignement-2/mongodb"
	. "IMT2681-assignement-2/data"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func InfoTicker(w http.ResponseWriter, r *http.Request) Ticker {
	tick := Ticker{}

	allTrack := mongodb.Global.GetAllTracks()
	tot := mongodb.Global.Count()
	if tot != 0 {
		for _, track := range allTrack {
			if track.Id <= CAP {
				// Amount of tracks equal to N-page element
				// or
				// CAP is greater than total amount of tracks
				if track.Id == CAP || CAP >= tot {
					tick.T_stop = track.Timestamp
				}
				// comment
				if tick.T_start == 0 {
					tick.T_start = track.Timestamp
				}
				tick.Tracks = append(tick.Tracks, track.Id)
			}
		}
		tick.T_latest = allTrack[tot-1].Timestamp
		return tick
	} else {
		return tick
	}
}

func GetLatest(w http.ResponseWriter, r *http.Request) {
	// Get the last track in db
	track, ok := mongodb.Global.Get(mongodb.Global.Count())

	if track.Id == mongodb.Global.Count() && ok {
		fmt.Fprint(w, track.Timestamp)
		return
	} else {
		// No tracks in db
		http.Error(w, "", http.StatusNoContent)
	}
}
func GetInfoTicker(w http.ResponseWriter, r *http.Request) {
	test := Ticker{}
	startTime := time.Now()
	tick := InfoTicker(w, r)
	if tick.T_start == test.T_start{
		http.Error(w,"db is empty", http.StatusNotFound)
	} else {
		tick.Processing = int64((time.Now().Sub(startTime)) / 1000000)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(tick); err != nil {
			fmt.Printf("PROBLEM: %v", err)
		}
	}
}
func GetTicker(timestamp int) Ticker {

	tick := Ticker{}

	allTrack := mongodb.Global.GetAllTracks()
	tot := mongodb.Global.Count()

	if tot != 0 {
		// Get the last timestamp in db
		tick.T_latest = allTrack[tot-1].Timestamp

		for _, track := range allTrack {
			if track.Id <= CAP { //track.Id%CAP != 0 &&
				// Amount of tracks equal to N-page element
				// or
				// CAP is greater than total amount of tracks
				if track.Id == CAP || CAP >= tot {
					tick.T_stop = track.Timestamp
				}
				// comment
				if track.Timestamp > int64(timestamp) && tick.T_start == 0 {
					tick.T_start = track.Timestamp
				}
				tick.Tracks = append(tick.Tracks, track.Id)
			}
		}
		return tick
	}

	return tick
}

func CalcTime(w http.ResponseWriter, r *http.Request) {
	test := Ticker{}
	vars := mux.Vars(r)
	timestamp, err := strconv.Atoi((vars["time"]))

	if err != nil {
		fmt.Printf("Error in Count(): %v", err.Error())
		return
	}

	startTime := time.Now()
	tick := GetTicker(timestamp)
	if tick.T_start == test.T_start {
		http.Error(w,"db is empty", http.StatusNotFound)
		return
	}
	tick.Processing = int64((time.Now().Sub(startTime)) / 1000000)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tick); err != nil {
		http.Error(w, "Could not encode", http.StatusInternalServerError)
		return
	}
}
