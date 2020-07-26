package mailer

import (
	"bytes"
	"html/template"
	"net/smtp"
)

type Mailer struct {
	auth smtp.Auth
	from string
}

type Configuration struct {
	from         string
	smtpMail     string
	smtpPassword string
	smtpHost     string
}

type Email struct {
	to      []string
	subject string
	body    string
}

type IMailer interface {
	SendPlainEmail(data Email) (bool, error)
	SendTemplateMail(data Email, tmplt template.Template, templateData interface{}) (bool, error)
}

func New(config Configuration) IMailer {
	auth := smtp.PlainAuth("", config.smtpMail, config.smtpPassword, config.smtpHost)
	return &Mailer{
		auth: auth,
		from: config.from,
	}
}

func (mailer *Mailer) SendPlainEmail(data Email) (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + data.subject + "!\n"
	msg := []byte(subject + mime + "\n" + data.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, mailer.auth, "dhanush@geektrust.in", data.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (mailer *Mailer) SendTemplateMail(data Email, tmplt template.Template, templateData interface{}) (bool, error) {
	buf := new(bytes.Buffer)
	if err := tmplt.Execute(buf, templateData); err != nil {
		return false, err
	}
	body := buf.String()
	data.body = body
	if _, err := mailer.SendPlainEmail(data); err != nil {
		return false, err
	}
	return true, nil
}
