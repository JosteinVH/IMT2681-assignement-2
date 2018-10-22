package mongodb

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	. "IMT2681-assignement-2/data"
)

var Global TrackStorage

type TracksMongoDB struct {
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

func (db *TracksMongoDB) Get(keyID string) (Tracks, bool) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	track := Tracks{}
	checkOk := true

	err = session.DB(db.DatabaseName).C(db.DatabaseCol).Find(bson.M{"ID":keyID}).One(&track)
	if err != nil  {
		checkOk = false
		fmt.Println("ERROR: %v", err)
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