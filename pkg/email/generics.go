package email

type EmailInterface interface {
	SendEmail(interface{}) error
}
