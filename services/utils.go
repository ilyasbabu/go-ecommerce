package services

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/wneessen/go-mail"
	"gorm.io/gorm"
)

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

func IsImageFile(reader io.Reader) bool {
	buffer := make([]byte, 512)
	_, err := reader.Read(buffer)
	if err != nil {
		return false
	}
	contentType := http.DetectContentType(buffer)
	return strings.HasPrefix(contentType, "image/")
}

var Db *gorm.DB

func SetDB(db *gorm.DB) {
	Db = db
}
