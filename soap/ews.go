package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gunsluo/goews/v2"
)

func main() {
	c, err := goews.NewClient(
		goews.Config{
			Address:  "https://outlook.office365.com/EWS/Exchange.asmx",
			Username: "email@exchangedomain",
			Password: "password",
			Dump:     true,
			NTLM:     false,
			Domain:   "",
			SkipTLS:  false,
		},
	)
	if err != nil {
		log.Fatal("->: ", err.Error())
	}

	filename := "a.txt"
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("read file: ", err.Error())
	}

	htmlBody := `<!DOCTYPE html>
<html lang="en">
<head>
  <title>Simple HTML document</title>
</head>
<body>
  <h1>The email body, as html!</h1>
</body>
</html>`

	err = c.SendEmail(
		goews.SendEmailParams{
			From:     "email@exchangedomain",
			To:       []string{"ji.luo@target-energysolutions.com"},
			Cc:       []string{"junkun.ren@target-energysolutions.com"},
			Bcc:      []string{"Dongsheng.liu@target-energysolutions.com"},
			Subject:  "An email subject",
			Body:     htmlBody,
			BodyType: goews.BodyTypeHtml,
			FileAttachments: []goews.AttachmentParams{
				{
					Name:        filename,
					ContentType: "",
					Size:        int64(len(content)),
					Content:     content,
				},
			},
		})
	if err != nil {
		log.Fatal("err>: ", err.Error())
	}

	fmt.Println("--- success ---")
}
