package mongodb

import (
	. "IMT2681-assignement-2/data"
	"github.com/globalsign/mgo"
	"os"
	"testing"
)

func setupDB(t *testing.T) *TracksMongoDB {
	db := TracksMongoDB{
		os.Getenv("DB_URL"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_COL_T"),
	}

	session, err := mgo.Dial(db.DatabaseURL)
	defer session.Close()

	if err != nil {
		t.Error(err)
	}
	return &db
}

func dropDB(t *testing.T, db *TracksMongoDB) {
	session, err := mgo.Dial(db.DatabaseURL)
	defer session.Close()
	if err != nil {
		t.Error(err)
	}
	err = session.DB(db.DatabaseName).DropDatabase()
	if err != nil {
		t.Error(err)
	}
}

func TestTrackMongoDB_Add(t *testing.T) {
	db := setupDB(t)
	defer dropDB(t, db)

	db.Init()
	if db.Count() != 0 {
		t.Error("Database not properly initialized, Track count should be 0")
	}
	tracks := Tracks{
		Id:1,
		Timestamp:1540810212915,
		H_date:"2016-02-19 00:00:00 +0000 UTC",
		Pilot:"Miguel Angel Gordillo",
		Glider:"RV8",
		GliderId:"EC-XLL",
		Track_length:443.2573603705269,
		Url:"http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc",
	}

	db.Add(tracks)
	if db.Count() != 1 {
		t.Error("Adding new Track failed.")
	}
}

func TestTrackMongoDB_Get(t *testing.T) {
	db := setupDB(t)
	defer dropDB(t, db)

	db.Init()
	if db.Count() != 0 {
		t.Error("database not properly initialized. Track count() should be 0.")
	}

	track := Tracks{
		Id:1,
		Timestamp:1540810212915,
		H_date:"2016-02-19 00:00:00 +0000 UTC",
		Pilot:"Miguel Angel Gordillo",
		Glider:"RV8",
		GliderId:"EC-XLL",
		Track_length:443.2573603705269,
		Url:"http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc",
	}

	db.Add(track)

	if db.Count() != 1 {
		t.Error("adding new Track failed.")
	}

	newTrack, ok := db.Get(track.Id)
	if !ok {
		t.Error("couldn't find track id 1",)
	}

	if newTrack.Id != track.Id ||
		newTrack.Timestamp != track.Timestamp ||
		newTrack.H_date != track.H_date ||
		newTrack.Pilot != track.Pilot ||
		newTrack.Glider != track.Glider ||
		newTrack.GliderId != track.GliderId ||
		newTrack.Track_length != track.Track_length||
		newTrack.Url != track.Url {
		t.Error("tracks do not match")
	}

	all := db.GetAllTracks()
	if len(all) != 1 || all[0].Id != track.Id {
		t.Error("GetAllTracks() doesn't return proper slice of all the items")
	}
}

func TestMongoDB_Delete(t *testing.T) {
	db := setupDB(t)
	defer dropDB(t, db)

	db.Init()
	if db.Count() != 0 {
		t.Error("Database not properly initialized, Track count should be 0")
	}
	track := Tracks{
		Id:1,
		Timestamp:1540810212915,
		H_date:"2016-02-19 00:00:00 +0000 UTC",
		Pilot:"Miguel Angel Gordillo",
		Glider:"RV8",
		GliderId:"EC-XLL",
		Track_length:443.2573603705269,
		Url:"http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc",
	}

	err := db.Add(track)
	if err != nil {
		t.Error("Adding new Track failed")
	}

	count := db.Count()


	ok := db.DelAll()

	if !ok {
		t.Error("Deleting the Track failed")
	}

	if db.Count() == count {
		t.Error("Deleting the Track failed")
	}

}


func setupDB_WebH(t *testing.T) *WebhookMongoDB {
	dbWeb := WebhookMongoDB{
		os.Getenv("DB_URL"),
		os.Getenv("DB_NAME"),
		"webhook",
	}

	session, err := mgo.Dial(dbWeb.DatabaseURL)
	defer session.Close()

	if err != nil {
		t.Error(err)
	}
	return &dbWeb
}

func dropDB_WebH(t *testing.T, dbWeb *WebhookMongoDB) {
	session, err := mgo.Dial(dbWeb.DatabaseURL)
	defer session.Close()
	if err != nil {
		t.Error(err)
	}
	err = session.DB(dbWeb.DatabaseName).DropDatabase()
	if err != nil {
		t.Error(err)
	}
}

func  TestWebhookMongoDB_Add(t *testing.T) {
	dbWeb := setupDB_WebH(t)
	defer dropDB_WebH(t, dbWeb)

	dbWeb.Init()
	if dbWeb.Count() != 0 {
		t.Error("Database not properly initialized, webhook count should be 0")
	}
	webH := Webhook{
		Id:"1",
		WebhookUrl:"https://hooks.slack.com/services/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",
		TriggerValue:2,
	}

	dbWeb.Add(webH)
	if dbWeb.Count() != 1 {
		t.Error("Adding new webhook failed.")
	}
}


func TestWebhookMongoDB_GetWebhook(t *testing.T) {
	dbWeb := setupDB_WebH(t)
	defer dropDB_WebH(t, dbWeb)

	dbWeb.Init()
	if dbWeb.Count() != 0 {
		t.Error("database not properly initialized. student count() should be 0.")
	}

	webH := Webhook{
		Id:"1",
		WebhookUrl:"https://hooks.slack.com/services/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",
		TriggerValue:2,
	}

	dbWeb.Add(webH)

	if dbWeb.Count() != 1 {
		t.Error("adding new webhook failed.")
	}

	newWebH, ok := dbWeb.GetWebhook(webH.Id)
	if !ok {
		t.Error("couldn't find webhook id 1",)
	}

	if newWebH.Id != webH.Id ||
		newWebH.TriggerValue != webH.TriggerValue ||
		newWebH.WebhookUrl != webH.WebhookUrl {
		t.Error("tracks do not match")
	}

	all := dbWeb.GetAllWebH()
	if len(all) != 1 || all[0].Id != webH.Id {
		t.Error("GetAllWebH() doesn't return proper slice of all the items")
	}
}


func TestWebhookMongoDB_DelWebhook(t *testing.T) {
	dbWeb := setupDB_WebH(t)
	defer dropDB_WebH(t, dbWeb)

	dbWeb.Init()
	if dbWeb.Count() != 0 {
		t.Error("Database not properly initialized, webhook count should be 0")
	}
	webH := Webhook{
		Id:"1",
		WebhookUrl:"https://hooks.slack.com/services/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",
		TriggerValue:2,
	}
	err := dbWeb.Add(webH)
	if err != nil {
		t.Error("Adding new webhook failed")
	}

	count := dbWeb.Count()

	ok := dbWeb.DelWebhook(webH.Id)

	if !ok {
		t.Error("Deleting the webhook failed")
	}

	if dbWeb.Count() == count {
		t.Error("Deleting the webhook failed")
	}

}