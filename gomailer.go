package gomailer

import (
	"gomailer/core"
	"gomailer/pkg/models"
	"html/template"
)

type IMailer interface {
	SendPlainEmail(data models.Email) (bool, error)
	SendTemplateMail(data models.Email, tmplt *template.Template, templateData interface{}) (bool, error)
}

type Mailer struct {
	core core.ICore
}

func New(config models.Configuration) IMailer {
	return &Mailer{
		core: core.New(config),
	}
}

func (m *Mailer) SendPlainEmail(data models.Email) (bool, error) {
	worked, err := m.core.SendMail(data, true)
	if worked {
		return worked, nil
	}
	return false, err
}

func (m *Mailer) SendTemplateMail(data models.Email, tmplt *template.Template, templateData interface{}) (bool, error) {
	worked, err := m.core.SendHTMLMail(data, tmplt, templateData)
	if worked {
		return worked, nil
	}
	return false, err
}
