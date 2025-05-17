package utils

import (
	"log"
	"os"
	"strings"

	"github.com/danilopsss/hsdeps/pkg/email"
)

func GetMailTrapCredentials() email.MailTrapCredentials {
	url := os.Getenv("MAILTRAP_URL")
	token := os.Getenv("MAILTRAP_TOKEN")

	if url == "" || token == "" {
		log.Fatal("Envvars not present.")
	}
	return email.MailTrapCredentials{
		Url:   strings.TrimSpace(url),
		Token: strings.TrimSpace(token),
	}
}
