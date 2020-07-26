package gomailer

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
	From         string
	SmtpMail     string
	SmtpPassword string
	SmtpHost     string
}

type Email struct {
	To      []string
	Subject string
	Body    string
}

type IMailer interface {
	SendPlainEmail(data Email) (bool, error)
	SendTemplateMail(data Email, tmplt *template.Template, templateData interface{}) (bool, error)
}

func New(config Configuration) IMailer {
	auth := smtp.PlainAuth("", config.SmtpMail, config.SmtpPassword, config.SmtpHost)
	return &Mailer{
		auth: auth,
		from: config.From,
	}
}

func (mailer *Mailer) SendPlainEmail(data Email) (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + data.Subject + "!\n"
	msg := []byte(subject + mime + "\n" + data.Body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, mailer.auth, mailer.from, data.To, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (mailer *Mailer) SendTemplateMail(data Email, tmplt *template.Template, templateData interface{}) (bool, error) {
	buf := new(bytes.Buffer)
	if err := tmplt.Execute(buf, templateData); err != nil {
		return false, err
	}
	body := buf.String()
	data.Body = body
	if _, err := mailer.SendPlainEmail(data); err != nil {
		return false, err
	}
	return true, nil
}
