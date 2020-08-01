package core

import (
	"bytes"
	"gomailer/pkg/models"
	"html/template"
	"net/smtp"
)

type ICore interface {
	SendMail(data models.Email, plain bool) (bool, error)
	SendHTMLMail(data models.Email, tmplt *template.Template, templateData interface{}) (bool, error)
}

type Core struct {
	auth smtp.Auth
}

func New(config models.Configuration) ICore {
	smtp := smtp.PlainAuth("", config.SmtpMail, config.SmtpPassword, config.SmtpHost)
	return &Core{
		auth: smtp,
	}
}

func (core *Core) SendMail(data models.Email, plain bool) (bool, error) {
	mime := ""
	if !plain {
		mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	}
	subject := "Subject: " + data.Subject + "\n"
	msg := []byte(subject + mime + data.Body)
	addr := "smtp.gmail.com:587"
	if err := smtp.SendMail(addr, core.auth, data.From, data.To, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (core *Core) SendHTMLMail(data models.Email, tmplt *template.Template, templateData interface{}) (bool, error) {
	buf := new(bytes.Buffer)
	if err := tmplt.Execute(buf, templateData); err != nil {
		return false, err
	}
	body := buf.String()
	data.Body = body
	if _, err := core.SendMail(data, false); err != nil {
		return false, err
	}
	return true, nil
}
