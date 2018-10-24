package data

type TrackStorage interface {
	Init()
	Add(t Tracks) error
	Count() int
	GetAllTracks() []Tracks
	Get(keyID int) (Tracks, bool)
	GetLastTrack() Tracks
}
/*
type TickerStorage interface {
	//Init()
	Add(ti Ticker) error
}*/

type Tracks struct {
	Id 			 int `json:"id"`
	Timestamp	 int64   `json:"timestamp"`
	H_date       string  `json:"H_date"`
	Pilot        string  `json:"pilot"`
	Glider       string  `json:"glider"`
	GliderId     string  `json:"glider_id"`
	Track_length float64 `json:"track_length"`
	Url          string  `json:"track_src_url"`
}

// Igcinfo
type Info struct {
	Uptime  string `json:"uptime"`
	Info    string `json:"info"`
	Version string `json:"version"`
}

// Track ids
type TrackId struct {
	Id int `json:"id"`
}

// POST URL
type Url struct {
	Url string `json:"url"`
}

type Ticker struct {
	T_latest   int64 `json:"t_latest"`
	T_start    int64 `json:"t_start"`
	T_stop     int64 `json:"t_stop"`
	Tracks     []int  `json:"tracks"`
	Processing int64 `json:"processing"`
}
