package utils

import (
	"os"
	"strings"

	"github.com/danilopsss/hsdeps/pkg/email"
)

/*
Function that depends on environment variables MAILTRAP_URL and MAILTRAP_TOKEN.
*/
func GetMailTrapCredentials() email.MailTrapCredentials {
	url := os.Getenv("MAILTRAP_URL")
	token := os.Getenv("MAILTRAP_TOKEN")

	if url == "" || token == "" {
		panic("Envvars not present.")
	}
	return email.MailTrapCredentials{
		Url:   strings.TrimSpace(url),
		Token: strings.TrimSpace(token),
	}
}
