package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const MaxMsgNum = 20 // 0...19, in total 20

func Main(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	channelID := r.URL.Query().Get("channelID")
	if channelID == "" {
		channelID = "cxumolket" // "durov"
	}
	log.Println("channelID:", channelID)

	msgNum, err := strconv.Atoi(r.URL.Query().Get("msgNum"))
	// msgNum = MaxMsgNum - msgNum
	if err != nil || msgNum < 1 || msgNum > MaxMsgNum {
		msgNum = 1
		if err != nil {
			log.Println("msgNum err:", err)
		}
	}
	log.Println("msgNum:", msgNum)

	messages_structured := getTGchannel(channelID)

	jsonResp, err := json.Marshal(messages_structured[msgNum-1])
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	} else {
		w.Write(jsonResp)
	}
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getTGchannel(channelID string) []interface{} {
	url := fmt.Sprintf("https://t.me/s/%s", channelID)
	res, err := http.Get(url)
	check(err)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	check(err)

	messages_structured := make([]interface{}, MaxMsgNum)
	doc.Find("div.tgme_widget_message_wrap").Each(func(i int, s *goquery.Selection) {
		message_structured := make(map[string]string)
		message_structured["text"] = s.Find("div.tgme_widget_message_text").Text()
		message_structured["html"], err = s.Find("div.tgme_widget_message_text").Html()
		check(err)
		message_structured["datetime"], _ = s.Find(".tgme_widget_message_date time").Attr("datetime")
		message_structured["owner"] = s.Find(".tgme_widget_message_owner_name").Text()
		message_structured["views"] = s.Find(".tgme_widget_message_views").Text()
		messages_structured[MaxMsgNum-i-1] = message_structured
	})

	return messages_structured
}
