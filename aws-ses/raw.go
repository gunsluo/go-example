package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

var (
	charset = "UTF-8"
	from    = "no-reply@target-energysolutions.com"
	to      = []string{"ji.luo@target-energysolutions.com"}
	cc      = []string{}
	bcc     = []string{}
)

const (
	rawPattern = `From: %s
To: %s
Subject: Test email (contains an attachment)
MIME-Version: 1.0
Content-type: Multipart/Mixed;
	boundary="NextPart"

--NextPart
Content-Type: text/plain; charset=UTF-8
Content-Transfer-Encoding: quoted-printable

This is the message body.

--NextPart
Content-Type: text/html; charset=UTF-8
Content-Transfer-Encoding: quoted-printable

<html>
<head></head>
<body>
<h1>Hello!</h1>
<p>Please see the attached file for a list of customers to contact.</p>
</body>
</html>

--NextPart
Content-Type: application/octet-stream; name="=?UTF-8?B?%s?="
Content-Description: =?UTF-8?B?%s?=
Content-Disposition: attachment; filename="=?UTF-8?B?%s?="
Content-Transfer-Encoding: base64

%s

--NextPart--
`
)

func main() {
	// Create a new session in the us-west-2 region.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsEast1RegionID),
		//Credentials: credentials.NewStaticCredentials("AKID", "SECRET_KEY", "TOKEN"),
	})
	if err != nil {
		panic(err)
	}

	// AWS_ACCESS_KEY_ID AWS_SECRET_ACCESS_KEY
	_, err = sess.Config.Credentials.Get()
	if err != nil {
		panic(err)
	}

	// Create SES session.
	svc := ses.New(sess)

	fn := "Effective Go中文版.pdf"
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	content := base64.StdEncoding.EncodeToString(b)
	fn = base64.StdEncoding.EncodeToString([]byte(fn))

	raw := fmt.Sprintf(rawPattern, from, strings.Join(to, ";"), fn, fn, fn, content)
	var addrs []string
	addrs = append(addrs, to...)
	//addrs = append(addrs, cc...)
	//addrs = append(addrs, bcc...)
	// Assemble the email.
	input := &ses.SendRawEmailInput{
		Destinations: aws.StringSlice(to),
		RawMessage: &ses.RawMessage{
			Data: []byte(raw),
		},
		Source: aws.String(from),
		//FromArn:      aws.String(""),
		//ReturnPathArn: aws.String(""),
		//SourceArn:     aws.String(""),
	}

	// Attempt to send the email.
	result, err := svc.SendRawEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println("===", aerr.Code())
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println("err:", err.Error())
		}

		return
	}

	fmt.Printf("Email Sent to address: %s", to)
	fmt.Println(result)
}
