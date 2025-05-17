package email

type EmailInterface interface {
	SendEmail(any) error
}
