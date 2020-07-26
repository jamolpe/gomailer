package main

import (
	"gomailer"
	"html/template"
	"log"
)

var smtpMail, smtpPassword, smtpHost = "---", "----", "smtp.gmail.com"

type templateData struct {
	Name string
	URL  string
}

func main() {
	config := gomailer.Configuration{From: "----", SmtpMail: smtpMail, SmtpPassword: smtpPassword, SmtpHost: smtpHost}
	mailer := gomailer.New(config)
	mail := gomailer.Email{To: []string{"----", "---"}, Subject: "im html template"}
	t, err := template.ParseFiles("basicTemplate.html")
	if err != nil {
		log.Fatal(err)
	}
	templateData := templateData{Name: "number", URL: "google.com"}
	mailer.SendTemplateMail(mail, t, templateData)
}
