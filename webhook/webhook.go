package webhook

import (
	"IMT2681-assignement-2/data"
	"IMT2681-assignement-2/mongodb"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)


const TheDiscordWebhook = "https://hooks.slack.com/services/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"
func SendDiscordLogEntry(now int,prev int) {

	if now != prev {
		info := data.WebhookInfo{}
		diff := now - prev
		info.Text = "Behind  "+strconv.Itoa(diff)+" tracks"
		raw, _ := json.Marshal(info)
		resp, err := http.Post(TheDiscordWebhook, "application/json", bytes.NewBuffer(raw))

		if err != nil {
			fmt.Println(err)
			fmt.Println(ioutil.ReadAll(resp.Body))
		}
	}
}

func Tracks() {
	for {
		delay := time.Minute * 10
		prev := mongodb.Global.Count()
		time.Sleep(delay)
		now := mongodb.Global.Count()
		SendDiscordLogEntry(now, prev)
	}
}