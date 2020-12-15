package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	charset  = "UTF-8"
	subject  = "Amazon SES Test (AWS SDK for Go)"
	htmlBody = "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
		"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
		"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"
	textBody = "This email was sent with Amazon SES using the AWS SDK for Go."
)

func main() {
	// Create a new session in the us-west-2 region.
	sess, err := session.NewSession(&aws.Config{
		//Region: aws.String(endpoints.UsWest2RegionID),
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

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			//CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String("luoji <ji.luo@target-energysolutions.com>"),
				aws.String("yanyu <yanyu.li@target-energysolutions.com>"),
				//aws.String("ji.luo@target-energysolutions.com"),
			},
		},
		Message: &ses.Message{
			Subject: &ses.Content{
				Charset: aws.String(charset),
				Data:    aws.String(subject),
			},
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(charset),
					Data:    aws.String(htmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(charset),
					Data:    aws.String(textBody),
				},
			},
		},
		//Source: aws.String("andrei@simionescu.eu"),
		//Source: aws.String("info@fluxble.com"),
		Source: aws.String("Tax Authority of Oman <info@fluxble.com>"),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

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

	fmt.Println("Email Sent to address: gunsluo@gmail.com")
	fmt.Println(result)
}
