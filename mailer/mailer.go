package mailer

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func SendMail(from string, password string, body string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", from)

	// Set E-Mail receivers
	m.SetHeader("To", "example@yahoo.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "Public IP Update")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "Your current public IP is "+body+"!")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.mail.yahoo.com", 587, from, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
