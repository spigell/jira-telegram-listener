package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	jira "jira-telegram-listener/jira"
	tg "jira-telegram-listener/telegram"
)

var (
	config = flag.String("config", "/etc/fira-telegram-listener.json", "path to config file")
)

type Config struct {
	TelegramToken string
	ChatId        string
	ListenPort    string
}

func TelegramHandler(token string, chatid string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}

		message, _ := jira.MakeMessageFromApi(body)

		result := tg.SendMessage(message, chatid, token)
		if result == 0 {
			log.Println("Message sent!")
		}

	}
}

func main() {
	flag.Parse()

	file, _ := os.Open(*config)
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)

	if err != nil {
		log.Panic(err)
	}

	listenPort := ":" + configuration.ListenPort

	http.HandleFunc("/", TelegramHandler(configuration.TelegramToken, configuration.ChatId))

	if err := http.ListenAndServe(listenPort, nil); err != nil {
		log.Fatal(err)
	}
}
