package email

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type MailTrapCredentials struct {
	Url   string
	Token string
}

type MailTrapFrom struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type MailTrapTo struct {
	Email string `json:"email"`
}

type MailTrapPayload struct {
	From     any    `json:"from"`
	To       any    `json:"to"`
	Subject  string `json:"subject"`
	Text     string `json:"text"`
	Category string `json:"category"`
}

func (mtc MailTrapCredentials) SendEmail(metadata MailTrapPayload) (err error) {
	log.Print("Preparing to send the email...")
	client := &http.Client{}
	payload, err := json.Marshal(metadata)
	if err != nil {
		panic("Cannot Marshal the payload!")
	}
	ioPayload := strings.NewReader(string(payload))
	request, err := http.NewRequest("POST", mtc.Url, ioPayload)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", mtc.Token))
	request.Header.Set("Content-Type", "application/json")
	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	log.Print("Email sent successfuly!")
	return
}
