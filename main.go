package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
)

func main() {

	token := os.Getenv("FRIDAY_BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	c := cron.New(cron.WithLocation(time.UTC))

	//18:00 MSK time, 15:00 UTC
	_, err = c.AddFunc("* 15 * * 5", func() {

		postBody, _ := json.Marshal(map[string]string{
			"chat_id":           "-1001995179603",                                                          //friends and family
			"message_thread_id": "4",                                                                       //popizdet'
			"video":             "BAACAgEAAxkDAAMKZuUAAWtF2n4zDMqIXmRBvepgQkeiAALyBAACChMoRx9v_UfIozK6NgQ", //friday.mp4
			"file_type":         "video",
		})
		responseBody := bytes.NewBuffer(postBody)

		resp, err := http.Post("https://api.telegram.org/bot"+token+"/sendVideo", "application/json", responseBody)
		//Handle Error
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()
		//Read the response body
		_, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

	})
	if err != nil {
		log.Fatal(err)
	}

	c.Start()
	select {}

}
