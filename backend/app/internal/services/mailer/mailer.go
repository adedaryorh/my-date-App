package mailer

type Mailer interface {
	SendOTP(email string, otp string) error
}

type EmailService struct{}

func NewMailerService() *EmailService {
	return &EmailService{}
}

func (*EmailService) SendOTP(email string, otp string) error {
	//TODO: Implement me
	return nil
}
