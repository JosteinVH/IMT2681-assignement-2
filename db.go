package main

import (
	"github.com/globalsign/mgo"
	"fmt"
)


func main(){

	session, err := mgo.Dial("mongodb://heihade:heihade123@ds131983.mlab.com:31983/josteivhdb")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	c := session.DB("josteivhdb").C("track")

	fmt.Println(c)
}

