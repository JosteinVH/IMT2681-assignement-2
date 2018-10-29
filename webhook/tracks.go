package main

import (
	."IMT2681-assignement-2/data"
	."IMT2681-assignement-2/mongodb"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)


const theDiscordWebhook = "https://hooks.slack.com/services/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"


func sendDiscordLogEntry(now int,prev int) {

	if now != prev {
		info := WebhookInfo{}
		info.Text = "Number of track: "+strconv.Itoa(now)+" previous: "+strconv.Itoa(now)+"\n"
		raw, _ := json.Marshal(info)
		resp, err := http.Post(theDiscordWebhook, "application/json", bytes.NewBuffer(raw))

		if err != nil {
			fmt.Println(err)
			fmt.Println(ioutil.ReadAll(resp.Body))
		}
	}
}

//TODO MAKE IT WORK WITH THE WHOLE PROGRAM
func main() {
	for {
		delay := time.Minute * 10
		prev := Global.Count()
		time.Sleep(delay)
		now := Global.Count()
		sendDiscordLogEntry(now, prev)
	}
}