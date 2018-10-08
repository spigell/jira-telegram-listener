package jira

import (
	"fmt"
	"encoding/json"
	"log"
	"strings"
)


type Message struct {
	Issue 		Issue 	 `json:"issue"`
	User		User 	 `json:"user"`
	Comment         Comment  `json:"comment"`
	WebhookEvent 	string `json:"webhookEvent"`
}

type Issue struct {
	Key		string 	 `json:"key"`
	Self 		string   `json:"self"`
	Fields		Fields
}

type User struct {
	Name 		string	 `json:"displayName"`
}

type Fields struct {
	Summary 	string   `json:"summary"`
	Description     string   `json:"description"`
}

type Comment struct {
	Author          Author   `json:"author"`
	Body            string   `json:"body"`
}

type Author struct {
	Name 		string	 `json:"displayName"`
}



func MakeMessageFromApi (body []byte) (string, error) {

	var message Message 

	if err  := json.Unmarshal(body, &message); err != nil {
		log.Println(err)
	}

	eventType := message.WebhookEvent

	log.Printf("Event %s was registered", eventType)

	slice := strings.Split(message.Issue.Self, "/")
	issueLink := fmt.Sprintf("%s//%s/browse/%s", slice[0], slice[2], message.Issue.Key)
	
	s := ""

	switch eventType {
	        case "jira:issue_created":
	        	s = fmt.Sprintf("New task created \nLink: %s \nSummary: %s \nName: %s \nDesctiption: %s\n", issueLink, message.Issue.Fields.Summary, message.User.Name, message.Issue.Fields.Description)
	        case "comment_created":
	        	s = fmt.Sprintf("New comment created \nLink: %s \nSummary: %s \nName: %s \nComment: %s\n", issueLink, message.Issue.Fields.Summary, message.Comment.Author.Name, message.Comment.Body)
	}
	return s, nil
}
