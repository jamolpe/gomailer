package main

import (
	"gomailer"
)

var smtpMail, smtpPassword, smtpHost = "---", "----", "smtp.gmail.com"

func main() {
	config := gomailer.Configuration{From: "----", SmtpMail: smtpMail, SmtpPassword: smtpPassword, SmtpHost: smtpHost}
	mailer := gomailer.New(config)
	mail := gomailer.Email{To: []string{"----", "----"}, Subject: "sup bro", Body: "im basic email"}
	mailer.SendPlainEmail(mail)
}
