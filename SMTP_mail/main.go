package main

import (
	"fmt"
	"net/smtp"
)

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

// serverName URI to smtp server
func (s *smtpServer) serverName() string {
	return s.host + ":" + s.port
}

func main() {
	// Sender data.
	from := "markussnow@gmail.com"
	password := "wjbvxwrqlbkeijki" // you can enter original password or password generated with App password
	// Receiver email address.
	to := []string{
		"markus19@mail.ru",
		//"second receiver email address",
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message.
	message := []byte("Hi Markus! \nThis is the test message.\nNo need to answer it!")
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending email.
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
