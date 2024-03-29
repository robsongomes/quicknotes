package main

import (
	"fmt"
	"log"
	"net/smtp"

	"gopkg.in/gomail.v2"
)

var (
	from       = "quicknotes@quick.com"
	msg        = []byte("Seja bem vindo ao Quicknotes!")
	recipients = []string{"user1@quick.com", "user2@quick.com"}
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", recipients...)
	m.SetAddressHeader("Bcc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Boas vindas")
	m.SetBody("text/html", "<h1 style='color: green'>Seja bem vindo ao Quicknotes.</h1>!")
	m.Attach("README.md")

	d := gomail.NewDialer("localhost", 1025, "", "")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func sendMail() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "", "", "localhost")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"user1@quick.com"}
	msg := []byte("To: user1@quick.com\r\n" +
		"Subject: Bem vindo!\r\n" +
		"Content-Type: text/html\r\n" +
		"\r\n" +
		"<h1 style='color: red'>Seja bem vindo ao Quicknotes.</h1>\r\n")
	err := smtp.SendMail("localhost:1025", auth, "quicknotes@quick.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func sendMailPlainAuth() {
	// hostname is used by PlainAuth to validate the TLS certificate.
	hostname := "localhost"
	auth := smtp.PlainAuth("", "", "", hostname)

	err := smtp.SendMail(hostname+":1025", auth, from, recipients, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func sendMailPackage() {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial("localhost:1025")
	if err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient first
	if err := c.Mail("quicknotes@quick.com"); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt("user1@quick.com"); err != nil {
		log.Fatal(err)
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(wc, "Seja bem vindo ao Quicknotes")
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
