package telegram


import (

	"net/url"
	"net/http"
	"fmt"
	"log"
	"io/ioutil"

)



func SendMessage(message string, chatId string, token string) (r string) {

	var client http.Client

        response, err := client.PostForm(
                fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token),
                url.Values{"chat_id": {chatId}, "text": {message}})

        defer response.Body.Close()

        if err != nil {
                log.Println(err)
                return 
        }

        body, err := ioutil.ReadAll(response.Body)

        if err != nil {
                log.Println(err)
                return 
        }

        return string(body)
}
