package mongodb

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	. "IMT2681-assignement-2/data"
)

var Global TrackStorage
var G_Webhook WebhookStorage
//var GTicker TickerStorage

type TracksMongoDB struct {
	DatabaseURL  string
	DatabaseName string
	DatabaseCol  string
}

type WebhookMongoDB struct {
	DatabaseURL  string
	DatabaseName string
	DatabaseCol  string
}


func (db *TracksMongoDB) Init() {
	// Make sure we can connect to database
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		//handle error
		panic(err)
	}

	fmt.Println("Connection...")
	defer session.Close()
}

func (db *TracksMongoDB) Add(t Tracks) error {
	// Make sure we can connect to database
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.DatabaseCol).Insert(t)
	if err != nil {
		fmt.Printf("Error in insert(): %v", err.Error())
		return err
	}

	return nil
}

func (db *TracksMongoDB) Count() int {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil{
		panic(err)
	}
	defer session.Close()

	count, err := session.DB(db.DatabaseName).C(db.DatabaseCol).Count()
	if err != nil{
		fmt.Printf("Error in Count(): %v", err.Error())
		return -1
	}

	return count
}

func (db *TracksMongoDB) Get(keyID int) (Tracks, bool) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	track := Tracks{}
	checkOk := true

	err = session.DB(db.DatabaseName).C(db.DatabaseCol).Find(bson.M{"id":keyID}).One(&track)
	if err != nil  {
		checkOk = false
		return Tracks{},checkOk
		//fmt.Println("ERROR: %v", err)
	}

	return track, checkOk
}

func (db *TracksMongoDB) GetAllTracks() []Tracks{
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil{
		panic(err)
	}
	defer session.Close()

	var all []Tracks

	err = session.DB(db.DatabaseName).C(db.DatabaseCol).Find(bson.M{}).All(&all)
	if err != nil {
		fmt.Println("%v",err)
		return []Tracks{}
	}

	return all
}


func(db *TracksMongoDB) AddTicker(ti Ticker) error {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		fmt.Printf("ERROR TICKER: %v", err)
		//panic(err)
	}
	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.DatabaseCol).Insert(ti)
	if err != nil{
		fmt.Printf("Error in insert: %v", err)
	}

	return nil
}

func (db *TracksMongoDB) DelAll() {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		fmt.Printf("Couldn't get last: %v", err)
	}
	defer session.Close()

	_,err = session.DB(db.DatabaseName).C(db.DatabaseCol).RemoveAll(bson.M{})

	if err != nil{
		fmt.Printf("Error in LastTrack: %v", err)
	}
}


func (dbWB *WebhookMongoDB) Add(w Webhook) error {
	// Make sure we can connect to database
	session, err := mgo.Dial(dbWB.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(dbWB.DatabaseName).C(dbWB.DatabaseCol).Insert(w)
	if err != nil {
		fmt.Printf("Error in insert(): %v", err.Error())
		return err
	}

	return nil
}


func (dbWB *WebhookMongoDB) GetAllWebH() []Webhook{
	session, err := mgo.Dial(dbWB.DatabaseURL)
	if err != nil{
		panic(err)
	}
	defer session.Close()

	var all []Webhook

	err = session.DB(dbWB.DatabaseName).C(dbWB.DatabaseCol).Find(bson.M{}).All(&all)
	if err != nil {
		fmt.Println("%v",err)
		return []Webhook{}
	}

	return all
}


func (dbWB *WebhookMongoDB)	UpdateW(url string,count int) {
	colQuerier := bson.M{"webhookurl": url}
	change := bson.M{"$set": bson.M{"count": count}}
	session, err := mgo.Dial(dbWB.DatabaseURL)
	if err != nil{
		panic(err)
	}
	defer session.Close()

	err = session.DB(dbWB.DatabaseName).C(dbWB.DatabaseCol).Update(colQuerier, change)
	if err != nil {
		fmt.Println("%v",err)
	}
}

func (dbWB *WebhookMongoDB) GetWebhook(keyID string) (Webhook, bool) {
	session, err := mgo.Dial(dbWB.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	webH := Webhook{}
	checkOk := true

	err = session.DB(dbWB.DatabaseName).C(dbWB.DatabaseCol).Find(bson.M{"id":keyID}).One(&webH)
	if err != nil  {
		checkOk = false
		return Webhook{},checkOk
	}

	return webH, checkOk
}


func (dbWB *WebhookMongoDB) DelWebhook(keyID string) bool{
	session, err := mgo.Dial(dbWB.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	ok := true

	err = session.DB(dbWB.DatabaseName).C(dbWB.DatabaseCol).Remove(bson.M{"id": keyID})
	if err != nil {
		fmt.Println("Remove failed: %v",err)
		ok = false
		return ok

	}

	return ok
}


