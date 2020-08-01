package test

import (
	"gomailer"
	"gomailer/pkg/models"
	"testing"
)

var smtpMail, smtpPassword, smtpHost = "email@gmail.com", "emailpassword", "smtp.gmail.com"

func validateError(expected gomailer.ValidationError, result gomailer.ValidationError) bool {
	if result.Code == "" {
		return false
	}
	if result.Message == "" {
		return false
	}
	if expected.Code != result.Code {
		return false
	}
	if expected.Message != result.Message {
		return false
	}
	return true
}

func Test_ValidationError_plain_noBody(t *testing.T) {
	config := models.Configuration{SmtpMail: smtpMail, SmtpPassword: smtpPassword, SmtpHost: smtpHost}
	mailer := gomailer.New(config)
	_, validationError, _ := mailer.SendPlainEmail(models.Email{To: []string{"tomail@mail.com"}, Subject: "subject@gmail.com", Body: "", From: "from@mail.com"})
	expected := validateError(gomailer.ValidationError{Code: "MAILER_VALIDATION_ERROR001", Message: "email body con not be void"}, *validationError)
	if !expected {
		t.Error("validation not the same")
	}
}

func Test_ValidationError_plain_noFrom(t *testing.T) {
	config := models.Configuration{SmtpMail: smtpMail, SmtpPassword: smtpPassword, SmtpHost: smtpHost}
	mailer := gomailer.New(config)
	_, validationError, _ := mailer.SendPlainEmail(models.Email{To: []string{"tomail@mail.com"}, Subject: "subject@gmail.com", Body: "boody", From: ""})
	expected := validateError(gomailer.ValidationError{Code: "MAILER_VALIDATION_ERROR002", Message: "email from can not be void"}, *validationError)
	if !expected {
		t.Error("validation not the same")
	}
}

func Test_ValidationError_plain_noTo(t *testing.T) {
	config := models.Configuration{SmtpMail: smtpMail, SmtpPassword: smtpPassword, SmtpHost: smtpHost}
	mailer := gomailer.New(config)
	_, validationError, _ := mailer.SendPlainEmail(models.Email{To: []string{}, Subject: "subject@gmail.com", Body: "boody", From: "from@gmail.com"})
	expected := validateError(gomailer.ValidationError{Code: "MAILER_VALIDATION_ERROR003", Message: "email has no from to send"}, *validationError)
	if !expected {
		t.Error("validation not the same")
	}
}
