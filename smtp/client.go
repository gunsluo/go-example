package main

import (
	"io"
	"log"
	"net"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

func main() {
	test()
}

func test() {
	// Set up authentication information.
	//auth := sasl.NewPlainClient("", "username", "password")
	auth := sasl.NewPlainClient("", "AKIAT2CIH646YJ4YR377", "+Aink7Ap4XqlnsRNM9FcT")

	/*
		auth := sasl.NewOAuthBearerClient(&sasl.OAuthBearerOptions{
			Username: "AKIAT2CIH646YJ4YR377",
			Token:    "+Aink7Ap4XqlnsRNM",
		})
	*/

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"ji.luo@target-energysolutions.com"}
	msg := strings.NewReader("Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, "weilong.yi@target-energysolutions.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

const (
	smtpHost = "email-smtp.us-east-1.amazonaws.com"
	smtpPort = "587"
)

func test2() {
	conn, err := net.Dial("tcp", smtpHost+":"+smtpPort)
	if err != nil {
		log.Fatal(err)
	}

	smtpClient, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		log.Fatal(err)
	}

	defer smtpClient.Close()

	//if err = smtpClient.hello(); err != nil {
	//	return err
	//}
	if ok, _ := smtpClient.Extension("STARTTLS"); ok {
		if err = smtpClient.StartTLS(nil); err != nil {
			log.Fatal(err)
		}
	}

	auth := sasl.NewPlainClient("", "AKIAT2CIH646YJ4YR377", "+Aink7Ap4XqlnsRNM")

	err = smtpClient.Auth(auth)
	if err != nil {
		log.Fatal(err)
	}

	err = smtpClient.Mail("ji.luo@target-energysolutions.com", &smtp.MailOptions{})
	if err != nil {
		log.Fatal(err)
	}

	to := []string{"gunsluo@gmail.com"}
	for _, addr := range to {
		if err = smtpClient.Rcpt(addr); err != nil {
			log.Fatal(err)
		}
	}
	w, err := smtpClient.Data()
	if err != nil {
		log.Fatal(err)
	}

	msg := strings.NewReader("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	_, err = io.Copy(w, msg)
	if err != nil {
		log.Fatal(err)
	}
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = smtpClient.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
