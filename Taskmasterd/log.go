package main

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

const (
	mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

type Request struct {
	From    string
	To      []string
	Subject string
	Body    string
	Path    string
	Data    interface{}
}

func NewRequest(to []string, subject, path string, data interface{}) *Request {
	return &Request{
		To:      to,
		From:    "golangmasters@gmail.com",
		Subject: subject,
		Path:    path,
		Data:    data,
	}
}

func (r *Request) SendMail() {
	var auth smtp.Auth
	auth = smtp.PlainAuth("", r.From, "Petassia01", "smtp.gmail.com")
	subject := "Subject: " + r.Subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.Body)
	addr := "smtp.gmail.com:587"
	err := smtp.SendMail(addr, auth, r.From, r.To, msg)
	if err != nil {
	}
}

func mm(a string) {
	data := struct {
		Path string
	}{
		Path: a,
	}
	r := NewRequest([]string{"golangmasters@gmail.com"}, "Reset Password", "mail", data)
	t, err := template.ParseFiles(r.Path)
	if err != nil {
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, r.Data); err != nil {
	}
	r.Body = buf.String()
	r.SendMail()
}

func registre(a, mess string, bbb ...int) {
	f, err := os.OpenFile("../log/registre.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(a + " " + mess)
	if len(bbb) > 0 {
		go mm(a + " " + mess)
	}
}
