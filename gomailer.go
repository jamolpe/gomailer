package gomailer

import (
	"gomailer/core"
	"gomailer/pkg/models"
	"html/template"
)

type IMailer interface {
	SendPlainEmail(data models.Email) (bool, *ValidationError, error)
	SendTemplateMail(data models.Email, tmplt *template.Template, templateData interface{}) (bool, *ValidationError, error)
}

type Mailer struct {
	core core.ICore
}

type ValidationError struct {
	Code    string
	Message string
}

func New(config models.Configuration) IMailer {
	return &Mailer{
		core: core.New(config),
	}
}

func (m *Mailer) validateMinimumData(data models.Email) *ValidationError {
	if data.Body == "" {
		return &ValidationError{Code: "MAILER_VALIDATION_ERROR001", Message: "email body con not be void"}
	}
	if data.From == "" {
		return &ValidationError{Code: "MAILER_VALIDATION_ERROR002", Message: "email from can not be void"}
	}
	if !(len(data.To) > 0) {
		return &ValidationError{Code: "MAILER_VALIDATION_ERROR003", Message: "email has no from to send"}
	}
	return nil
}

func (m *Mailer) SendPlainEmail(data models.Email) (bool, *ValidationError, error) {
	if validation := m.validateMinimumData(data); validation != nil {
		return false, validation, nil
	}
	worked, err := m.core.SendMail(data, true)
	if worked {
		return worked, nil, nil
	}
	return false, nil, err
}

func (m *Mailer) SendTemplateMail(data models.Email, tmplt *template.Template, templateData interface{}) (bool, *ValidationError, error) {
	if validation := m.validateMinimumData(data); validation != nil {
		return false, validation, nil
	}
	worked, err := m.core.SendHTMLMail(data, tmplt, templateData)
	if worked {
		return worked, nil, nil
	}
	return false, nil, err
}
