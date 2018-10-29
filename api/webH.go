package api

import (
	."IMT2681-assignement-2/data"
	"IMT2681-assignement-2/mongodb"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func RegWebH(w http.ResponseWriter, r *http.Request) {
	var webURL Webhook

	if err := json.NewDecoder(r.Body).Decode(&webURL); err != nil {
		http.Error(w, "Check body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()


	// Set triggervalue if none is provided
	if webURL.TriggerValue < 1 {
		webURL.TriggerValue = 1
	}
	webURL.Id = strconv.Itoa(ider)
	// Register webhookUR
	for _, wH := range mongodb.G_Webhook.GetAllWebH() {
		if wH.WebhookUrl == webURL.WebhookUrl {
			http.Error(w,"Already exists", http.StatusConflict)
			return
		}
	}
	mongodb.G_Webhook.Add(webURL)
	fmt.Fprintf(w, webURL.Id)
	ider++
}

func calcProcTime(id int) {
	test = append(test, strconv.Itoa(id))

	url := WebhookInfo{}
	startTime := time.Now()
	text, webURL := NyFunc()
	if text == "" && webURL == "" {
		fmt.Printf("No such url")
	}
	processing := (int((time.Now().Sub(startTime)))/ 1000000)

	url.Text = text + strconv.Itoa(processing)+"ms"

	b, err := json.Marshal(url)
	if err != nil {
		fmt.Printf("Error in Count(): %v", err.Error())
	}
	http.Post(webURL, "application-json", bytes.NewBuffer((b)))
}

func NyFunc() (string,string){
	totTracks := mongodb.Global.GetAllTracks()
	count := mongodb.Global.Count()
	allWebH := mongodb.G_Webhook.GetAllWebH()

	for _, wH := range allWebH {
		if len(test) % wH.TriggerValue == 0 {
			url := WebhookInfo{}
			javel := Convertion(wH.TriggerValue)
			url.Text = "Latest timestamp: " + strconv.Itoa(int(totTracks[count-1].Timestamp)) + "\n" + strconv.Itoa(wH.TriggerValue) + " new tracks are ID: " + javel + "\n" + "Processing: "

			return url.Text, wH.WebhookUrl
		}
	}
	return "",""
}

func GetWebH(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	webID:= vars["id"]

	webH, ok := mongodb.G_Webhook.GetWebhook(webID)
	if !ok {
		http.Error(w, "No such ID '"+webID+"'",http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(webH); err != nil {
		http.Error(w,"Could not encode", http.StatusInternalServerError)
	}

}

func DelWebH(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	webID:= vars["id"]

	webH, check := mongodb.G_Webhook.GetWebhook(webID)

	if !check {
		http.Error(w,"Could not get webhook", http.StatusConflict)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(webH); err != nil {
		http.Error(w,"Failed to decode", http.StatusNotFound)
	}

	ok := mongodb.G_Webhook.DelWebhook(webID)
	if !ok {
		http.Error(w,"Failed to delete", http.StatusInternalServerError)
	}
}

func Convertion(count int) string{
	var testing []string
	for i := len(test)-count; i<len(test);i++  {
		sjekk,_:= strconv.Atoi(test[i])
		if i <= sjekk {
			testing = append(testing,test[i])
		}
	}
	k := strings.Join(testing, ",")
	return k
}


