package main

import (
	"fmt"
	"gomailer"
	"gomailer/pkg/models"
	"html/template"
	"log"
)

var smtpMail, smtpPassword, smtpHost = "email@gmail.com", "emailpassword", "smtp.gmail.com"

type templateData struct {
	Name string
	URL  string
}

func main() {
	config := models.Configuration{SmtpMail: smtpMail, SmtpPassword: smtpPassword, SmtpHost: smtpHost}
	mailer := gomailer.New(config)
	mail := models.Email{To: []string{"email@gmail.com", smtpMail}, Subject: "im html template", Body: "im basic email", From: "email@gmail.com"}

	t, err := template.ParseFiles("basicTemplate.html")
	if err != nil {
		log.Fatal(err)
	}
	templateData := templateData{Name: "number", URL: "google.com"}
	sent, validationErr, err := mailer.SendTemplateMail(mail, t, templateData)
	fmt.Println("result: %v,%v,%v", sent, validationErr, err)

}
