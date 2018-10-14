package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"fmt"

	jira "jira-telegram-listener/jira"
	tg "jira-telegram-listener/telegram"
	
	"gopkg.in/yaml.v2"

)

var (
	config = flag.String("config", "/etc/jira-telegram-listener.yml", "path to config file")
	version = flag.Bool("version", false, "show current version")
	BuildVersion = "none"
)

type Config struct {
	TelegramToken 		string
	TelegramChatId		string
	ListenPort    		string
}

func TelegramHandler(token string, chatid string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}

		message, _ := jira.MakeMessageFromApi(body)

		result := tg.SendMessage(message, chatid, token)
		log.Printf("[DEBUG] Responce from telegram: ", result)

	}
}

func main() {

	flag.Parse()

	if *version != false {
		fmt.Println(BuildVersion)
		os.Exit(0)
	}

	file, _ := os.Open(*config)
        configuration := Config{}
        target, _ := ioutil.ReadAll(file)

        err := yaml.Unmarshal(target, &configuration)
        if err != nil {
                log.Printf("[ERROR] Error while parsing configuration: ", err)
        }


	listenPort := ":" + configuration.ListenPort

	http.HandleFunc("/", TelegramHandler(configuration.TelegramToken, configuration.TelegramChatId))

	if err := http.ListenAndServe(listenPort, nil); err != nil {
		log.Fatal(err)
	}
}
