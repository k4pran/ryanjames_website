package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
)

type Mail struct {
	senderAddr    string
	recipientAddr string
	subject       string
	body          string
}

type SmtpServer struct {
	host string
	port string
}

func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderAddr)
	message += fmt.Sprintf("To: %s\r\n", mail.recipientAddr)

	message += fmt.Sprintf("Subject: %s from %s\r\n", mail.subject, mail.senderAddr)
	message += "\r\n" + mail.body

	return message
}

func SendMail(subject, name, senderAddr, message string, email string, pass string) {
	mail := Mail{senderAddr, email, subject, message}
	messageBody := mail.BuildMessage()

	smtpServer := SmtpServer{host: "smtp.gmail.com", port: "465"}

	log.Println(smtpServer.host)
	//build an auth
	auth := smtp.PlainAuth("", email, pass, smtpServer.host) // todo password

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		log.Panic(err)
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// step 2: add all from and to
	if err = client.Mail(mail.senderAddr); err != nil {
		log.Panic(err)
	}

	if err = client.Rcpt(mail.recipientAddr); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Mail sent successfully")
}