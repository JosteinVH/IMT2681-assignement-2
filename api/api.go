package api

import (
	. "IMT2681-assignement-2/data"
	"IMT2681-assignement-2/mongodb"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/marni/goigc"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// The time
var timer = time.Now()

var ider int = 1

const (
	FIRST = 1
	CAP   = 5
)

//TODO Ticker in memory
//TODO Webhook in memory
var Urls[] Webhook

func conver(d time.Duration) string {
	// For string manipulation
	var felles []string
	sec := d.Seconds()

	const (
		mins   = 60       // Minutes in seconds
		hours  = 3600     // Hours in seconds
		days   = 86400    // Days in seconds2
		months = 2629746  // Months in seconds
		years  = 31556952 // Years in seconds
	)

	felles = append(felles, "P")

	// Divide seconds with years in seconds to find number of current years
	year := int(sec / years)
	if year >= 1 {
		felles = append(felles, strconv.Itoa(year))
		felles = append(felles, "Y")
		// Subtracting the number of years in seconds - to provide right amount of seconds
		sec -= float64(years * year)
	}
	// Divide seconds with months in seconds to find number of current months
	month := int(sec / months)
	if month >= 1 {
		felles = append(felles, strconv.Itoa(month))
		felles = append(felles, "M")
		// Subtracting the number of months in seconds - to provide right amount of seconds
		sec -= float64(months * month)
	}

	// Divide seconds with days in seconds to find number of current days
	day := int(sec / days) // Days in seconds
	if day >= 1 {
		felles = append(felles, strconv.Itoa(day))
		felles = append(felles, "D")
		// Subtracting the number of days in seconds - to provide right amount of seconds
		sec -= float64(86400 * day)
	}

	felles = append(felles, "T")

	// Divide seconds with hours in seconds to find number of current hours
	hour := int(sec / hours) // Hours in seconds
	if hour >= 1 {
		felles = append(felles, strconv.Itoa(hour))
		felles = append(felles, "H")
		// Subtracting the number of hours in seconds - to provide right amount of seconds
		sec -= float64(hours * hour)

	}

	// Divide seconds with minutes in seconds to find number of current minutes
	min := int(sec / mins) // Minutes in seconds
	if min >= 1 {
		felles = append(felles, strconv.Itoa(min))
		felles = append(felles, "M")
		sec -= float64(mins * min)

	}

	if sec >= 0 {
		felles = append(felles, strconv.Itoa(int(sec)))
		felles = append(felles, "S")
	}

	// Joins the part of the slice to one string
	k := strings.Join(felles, "")
	// Returns string with corresponding timestamp
	return k
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	// Time since application started
	uptime := time.Since(timer)
	iso := conver(uptime)
	infoApi := Info{
		iso,
		"Service for IGC tracks.",
		"v1",
	}

	// Set the header to json
	w.Header().Set("Content-Type", "application/json")
	// Encodes information to user
	if err := json.NewEncoder(w).Encode(infoApi); err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
}

func GetAllId(w http.ResponseWriter, r *http.Request) {
	var ids = make([]int, 0)
	for _, tr := range mongodb.Global.GetAllTracks() {
		ids = append(ids, tr.Id)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(ids); err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
}

func AddTrack(w http.ResponseWriter, r *http.Request) {
	var igcUrl Url
	// If sent data is actual json
	if err := json.NewDecoder(r.Body).Decode(&igcUrl); err != nil {
		http.Error(w, "Check body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	track, err := igc.ParseLocation(igcUrl.Url)
	// Checks for valid URL sent in body
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {

		// Finds total track_length
		totalDistance := 0.0
		for i := 0; i < len(track.Points)-1; i++ {
			totalDistance += track.Points[i].Distance(track.Points[i+1])
		}

		// SLICE OF INT TO KEEP TRACK OF THE POST ID'S
		idCount := mongodb.Global.Count()
		idCount++
		// Unique time
		now := time.Now()
		unixNano := now.UnixNano()
		start := unixNano / 1000000

		var t Tracks = Tracks{
			idCount,
			start,
			track.Date.String(),
			track.Pilot,
			track.GliderType,
			track.GliderID,
			totalDistance,
			igcUrl.Url,
		}

		mongodb.Global.Add(t)
		//TODO Call webhook (Send the track?)
		//RegWebH(w,r)
		w.Header().Set("Content-Type", "application/json")
		// Encodes unique id in json - back to user
		if err := json.NewEncoder(w).Encode(TrackId{Id: idCount}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	trackId, _ := strconv.Atoi(id)

	t, ok := mongodb.Global.Get(trackId)

	if !ok {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Encodes information for a specific track in json back to user
	if err := json.NewEncoder(w).Encode(t); err != nil {
		http.Error(w, "Could not encode", http.StatusInternalServerError)
		return
	}
}

func GetTrackProp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	prop := vars["prop"]

	trackId, _ := strconv.Atoi(id)

	ts, ok := mongodb.Global.Get(trackId)

	if !ok {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	switch prop {
	case "pilot":
		fmt.Fprint(w, ts.Pilot)
		return
	case "glider":
		fmt.Fprint(w, ts.Glider)
		return
	case "glider_id":
		fmt.Fprint(w, ts.GliderId)
		return
	case "track_length":
		fmt.Fprint(w, ts.Track_length)
		return
	case "H_date":
		fmt.Fprint(w, ts.H_date)
		return
	default:
		// Returns 404 with empty body - when not known method is provided
		http.Error(w, "", http.StatusNotFound)
		return
	}
}

func InfoTicker(w http.ResponseWriter, r *http.Request) Ticker {
	//LAST := mongodb.Global.Count()
	// TODO Almost same as "GetTicker" --> See if they can be combined
	tick := Ticker{}

	allTrack := mongodb.Global.GetAllTracks()
	tot := mongodb.Global.Count()
	// TODO Check for empty DB
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
}

func GetInfoTicker(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	tick := InfoTicker(w, r)
	tick.Processing = int64((time.Now().Sub(startTime)) / 1000000)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tick); err != nil {
		fmt.Printf("PROBLEM: %v", err)
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

func GetTicker(timestamp int) Ticker {

	tick := Ticker{}

	allTrack := mongodb.Global.GetAllTracks()
	tot := mongodb.Global.Count()

	// Get the last timestamp in db
	tick.T_latest = allTrack[tot-1].Timestamp
	// TODO Check for empty DB

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

func CalcTime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	timestamp, err := strconv.Atoi((vars["time"]))

	if err != nil {
		fmt.Println("Could not convert: %v", err)
	}

	startTime := time.Now()
	tick := GetTicker(timestamp)
	tick.Processing = int64((time.Now().Sub(startTime)) / 1000000)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tick); err != nil {
		http.Error(w, "Could not encode", http.StatusInternalServerError)
		return
	}
}

func RegWebH(w http.ResponseWriter, r *http.Request) {
	var webURL Webhook

	// Register webhookUR
	if err := json.NewDecoder(r.Body).Decode(&webURL); err != nil {
		http.Error(w, "Check body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Set triggervalue if none is provided
	if webURL.TriggerValue < 1 {
		webURL.TriggerValue = 1
	}

	Urls = append(Urls, webURL)
	webURL.Id = strconv.Itoa(ider)
	fmt.Fprintf(w, webURL.Id)

	ider++

}