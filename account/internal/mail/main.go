package mail

import (
	"net/smtp"

	"github.com/Bendomey/RideHail/account/pkg/utils"
)

type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

var password, from, host, port string

func init() {
	password = utils.MustGet("MAIL_PASSWORD")
	from = utils.MustGet("MAIL_FROM")
	host = utils.MustGet("MAIL_HOST")
	port = utils.MustGet("MAIL_PORT")
}

//SendMail helps send mail
func SendMail(toAdmin string, messageAdmin string) error {
	// Receiver email address.
	to := []string{
		toAdmin,
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: host, port: port}
	// Message.
	message := []byte(messageAdmin)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		return err
	}
	return nil
}
