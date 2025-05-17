package email

type EmailInterface interface {
	SendEmail(interface{}) error
}

type From struct {
	From interface{} `json:"from"`
}

type To struct {
	To interface{} `json:"to"`
}
