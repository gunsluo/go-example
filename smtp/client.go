package main

import (
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

var (
	smtpAddress        = ""
	smtpUsername       = ""
	smtpPassword       = ""
	smtpHost           = ""
	smtpPort           = ""
	smtpAuthType       = ""
	smtpAuthNtlmDomain = ""

	smtpEnableSSL  = false
	smtpServerName = ""

	//smtpInsecure = false
)

func main() {
	flag.StringVar(&smtpAddress, "address", "email-smtp.us-east-1.amazonaws.com:587", "host for smtp server")
	flag.StringVar(&smtpUsername, "username", "AKIAT2CIH646YJ4YR377", "username for smtp server")
	flag.StringVar(&smtpPassword, "password", "BGWfEYmkca5dLTZQEUVJUm2+Aink7Ap4XqlnsRNM9FcT", "password for smtp server")
	flag.StringVar(&smtpAuthType, "auth", "plain", "auth type for stmp, default: plain")
	flag.StringVar(&smtpAuthNtlmDomain, "ntlm-domain", "", "ntlm domain")

	flag.BoolVar(&smtpEnableSSL, "enable-ssl", false, "enable ssl for smtp connection")
	flag.StringVar(&smtpServerName, "servername", "", "certificate server name for smtp server")

	//flag.BoolVar(&smtpInsecure, "insecure", false, "insecure for smtp server")

	flag.Parse()

	host, port, err := net.SplitHostPort(smtpAddress)
	if err != nil {
		log.Fatal(err)
	}
	smtpHost = host
	smtpPort = port

	fmt.Println("-->", smtpAddress, smtpUsername, smtpPassword, smtpHost, smtpPort, smtpEnableSSL, smtpServerName)

	// sendMailByNtlm()
	sendMail()
	// test2()
}

func sendMail() {
	var smtpClient *smtp.Client
	if smtpEnableSSL {
		tlsConfig := &tls.Config{
			ServerName: smtpServerName,
			//InsecureSkipVerify: true,
		}

		client, err := smtp.DialTLS(smtpAddress, tlsConfig)
		if err != nil {
			log.Fatal(err)
		}
		smtpClient = client
	} else {
		client, err := smtp.Dial(smtpAddress)
		if err != nil {
			log.Fatal(err)
		}
		smtpClient = client
	}
	/*
		conn, err := net.Dial("tcp", smtpAddress)
		if err != nil {
			log.Fatal(err)
		}

		smtpClient, err := smtp.NewClient(conn, smtpHost)
		if err != nil {
			log.Fatal(err)
		}
	*/

	defer smtpClient.Close()

	if err := smtpClient.Hello(smtpHost); err != nil {
		log.Fatal(err)
	}
	var err error
	if ok, _ := smtpClient.Extension("STARTTLS"); ok {
		fmt.Println("-->", ok)
		tlsConfig := &tls.Config{InsecureSkipVerify: true}
		if err = smtpClient.StartTLS(tlsConfig); err != nil {
			log.Fatal(err)
		}
	}

	//sasl.NewExternalClient(id)

	var auth sasl.Client
	if smtpAuthType == "ntlm" {
		if smtpEnableSSL {
			err := errors.New("TLS not supported with NTLM")
			log.Fatal(err)
		}
		auth = NewNtlmClient(smtpAuthNtlmDomain, smtpUsername, smtpPassword, smtpUsername)
	} else if smtpAuthType == "login" {
		auth = sasl.NewLoginClient(smtpUsername, smtpPassword)
	} else {
		auth = sasl.NewPlainClient("", smtpUsername, smtpPassword)
	}

	err = smtpClient.Auth(auth)
	if err != nil {
		log.Fatalf("auth: %s", err)
	}

	err = smtpClient.Mail("weilong.yi@target-energysolutions.com", &smtp.MailOptions{})
	if err != nil {
		log.Fatalf("Mail: %s", err)
	}

	to := []string{"ji.luo@target-energysolutions.com"}
	for _, addr := range to {
		if err = smtpClient.Rcpt(addr); err != nil {
			log.Fatalf("Rcpt: %s", err)
		}
	}
	w, err := smtpClient.Data()
	if err != nil {
		log.Fatalf("Data: %s", err)
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

/*
func test() {
	// Set up authentication information.
	//auth := sasl.NewPlainClient("", "username", "password")
	auth := sasl.NewPlainClient("", "AKIAT2CIH646YJ4YR377", "BGWfEYmkca5dLTZQEUVJUm2+Aink7Ap4XqlnsRNM9FcT")

		auth := sasl.NewOAuthBearerClient(&sasl.OAuthBearerOptions{
			Username: "AKIAT2CIH646YJ4YR377",
			Token:    "+Aink7Ap4XqlnsRNM",
		})

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

func test2() {
	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Fatal(err)
	}

	d := gomail.NewDialer(smtpHost, port, smtpUsername, smtpPassword)
	d.SSL = true
	d.TLSConfig = &tls.Config{
		ServerName:         smtpServerName,
		InsecureSkipVerify: true,
	}

	//msg := &gomail.Message{}
	msg := gomail.NewMessage()
	msg.SetHeader("From", "weilong.yi@target-energysolutions.com")
	msg.SetHeader("To", "ji.luo@target-energysolutions.com")
	//msg.SetAddressHeader("Cc", "dan@example.com", "Dan")
	msg.SetHeader("Subject", "Hello!")
	msg.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")

	err = d.DialAndSend(msg)
	if err != nil {
		log.Fatal(err)
	}

	//gomail.Send()
}

func sendMailByNtlm() {
	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Fatal(err)
	}

	email := mail.NewEMail(`{"port":25}`)
	email.From = `weilong.yi@target-energysolutions.com`
	email.Host = smtpHost
	email.Port = port // [587 NTLM AUTH] [465，994]
	email.Username = smtpUsername
	//email.Secure = `` // SSL，TSL
	email.Password = smtpPassword

	email.Auth = mail.NTLMAuth(email.Host, email.Username, email.Password, mail.NTLMVersion2)

	email.To = []string{`weilong.yi@target-energysolutions.com`}
	email.Subject = `send mail success`
	email.Text = "Hello：\r\n this is a test email"
	//email.AttachFile(reportFile)
	if err := email.Send(); err != nil {
		log.Fatal(err)
	}
}
*/
