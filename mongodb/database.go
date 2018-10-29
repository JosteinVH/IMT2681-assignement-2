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

	fmt.Printf("Connection...")
	defer session.Close()
}

/*
Add adds new tracks to the db.
*/
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

/*
Count returns the current count of the tracks in in-memory storage.
*/
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

/*
Get returns a track with a given ID or empty track struct.
*/

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
		fmt.Printf("Error %v", err.Error())
		checkOk = false
		return Tracks{},checkOk
	}

	return track, checkOk
}


/*
GetAll returns a slice with all the tracks.
*/
func (db *TracksMongoDB) GetAllTracks() []Tracks{
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil{
		panic(err)
	}
	defer session.Close()

	var all []Tracks

	err = session.DB(db.DatabaseName).C(db.DatabaseCol).Find(bson.M{}).All(&all)
	if err != nil {
		fmt.Printf("Error %v", err.Error())
		return []Tracks{}
	}

	return all
}

/*
Delete every track in database
*/
func (db *TracksMongoDB) DelAll() bool {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		fmt.Printf("Couldn't get last: %v", err)
	}
	defer session.Close()

	_,err = session.DB(db.DatabaseName).C(db.DatabaseCol).RemoveAll(bson.M{})
	ok := true

	if err != nil{
		fmt.Printf("Error in LastTrack: %v", err)
		ok = false
		return ok
	}
	return ok
}

/*
For testing purpose*
Init initializes the mongo storage.
 */
func (dbWB *WebhookMongoDB) Init() {

	session, err := mgo.Dial(dbWB.DatabaseURL)
	if err != nil {
		//handle error
		panic(err)
	}

	fmt.Printf("Connection...")
	defer session.Close()
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


/*
GetAll returns a slice with all the tracks.
*/
func (dbWB *WebhookMongoDB) GetAllWebH() []Webhook{
	session, err := mgo.Dial(dbWB.DatabaseURL)
	if err != nil{
		panic(err)
	}
	defer session.Close()

	var all []Webhook

	err = session.DB(dbWB.DatabaseName).C(dbWB.DatabaseCol).Find(bson.M{}).All(&all)
	if err != nil {
		fmt.Printf("Error in getallWebH(): %v", err.Error())
		return []Webhook{}
	}

	return all
}

/*
Get returns a webhook with a given ID or empty webhook struct.
*/
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

/*
Delete specific webhook in database
*/
func (dbWB *WebhookMongoDB) DelWebhook(keyID string) bool{
	session, err := mgo.Dial(dbWB.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	ok := true

	err = session.DB(dbWB.DatabaseName).C(dbWB.DatabaseCol).Remove(bson.M{"id": keyID})
	if err != nil {
		fmt.Printf("Remove failed: %v",err)
		ok = false
		return ok

	}

	return ok
}
/*
Count returns the current count of the tracks in in-memory storage.
Made it for testing**
*/

func (dbWB *WebhookMongoDB) Count() int {
	session, err := mgo.Dial(dbWB.DatabaseURL)
	if err != nil{
		panic(err)
	}
	defer session.Close()

	count, err := session.DB(dbWB.DatabaseName).C(dbWB.DatabaseCol).Count()
	if err != nil{
		fmt.Printf("Error in Count(): %v", err.Error())
		return -1
	}

	return count
}
