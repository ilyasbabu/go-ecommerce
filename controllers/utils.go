package controllers

import (
	"log"
	"os"

	"github.com/wneessen/go-mail"
	"gorm.io/gorm"
)

type ResponseStruct struct {
	Status  string `json:"status"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func Response() ResponseStruct {
	return ResponseStruct{
		Status:  "ERROR",
		Data:    nil,
		Message: "Something Went Wrong",
	}
}

var Db *gorm.DB

func SetDB(db *gorm.DB) {
	Db = db
}

func SendMail(to string, subject string, body string) {
	m := mail.NewMsg()
	if err := m.From(os.Getenv("EMAIL")); err != nil {
		log.Fatalf("failed to set From address: %s", err)
	}
	if err := m.To(to); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}
	m.Subject(subject)
	m.SetBodyString(mail.TypeTextPlain, body)
	c, err := mail.NewClient("smtp.gmail.com", mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(os.Getenv("EMAIL")), mail.WithPassword(os.Getenv("EMAIL_PASSWORD")))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}
	if err := c.DialAndSend(m); err != nil {
		log.Fatalf("failed to send mail: %s", err)
	}
}
