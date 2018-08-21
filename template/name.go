package main

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

type EmailPlaceholder struct {
	CompanyName       string
	CompanyLocation   string
	JobSeekerName     string
	JobTitle          string
	JobDescription    string
	InterviewDate     string
	InterviewLocation string
}

var emailTmpl *template.Template
var emailTmpl2 *template.Template

func init() {
	emailTmpl = template.New("")
	emailTmpl2 = template.New("")
}

func main() {

	//emailTmpl = template.New("email")
	//emailTmpl2 = template.New("email2")

	str := "we are {{.CompanyName}} {{.CompanyLocation}} {{.JobSeekerName}} {{.JobTitle}} {{.JobDescription}} {{.InterviewDate}} {{.InterviewLocation}}"
	str2 := "Dear [JS name], <br/><br/>We would like to inform you that your application with [company name] for the [job title] position has been received. It would be appreciated  if you wait until the company responds to your request.<br/><br/><br/>[ Company Customized message space ]<br/><br/><br/>Thank you for your cooperation,<br/><br/>[Company name]"
	p := &EmailPlaceholder{
		CompanyName:       "demo",
		CompanyLocation:   "chengdu",
		JobSeekerName:     "luoji",
		JobTitle:          "",
		JobDescription:    "none",
		InterviewDate:     time.Now().Format("2006-01-02 15:04:05"),
		InterviewLocation: "chengdu",
	}

	tmpl, err := emailTmpl.Parse(str)
	if err != nil {
		panic(err)
	}

	tmpl2, err := emailTmpl2.Parse(str2)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, p)
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", buf.String())

	buf = bytes.Buffer{}
	err = tmpl2.Execute(&buf, p)
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", buf.String())

	buf = bytes.Buffer{}
	err = tmpl.Execute(&buf, p)
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", buf.String())

	buf = bytes.Buffer{}
	err = tmpl2.Execute(&buf, p)
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", buf.String())

	/*
		tmpl, err := emailTmpl.Parse(str)
		if err != nil {
			panic(err)
		}

		var buf bytes.Buffer
		err = tmpl.ExecuteTemplate(&buf, "email", p)
		if err != nil {
			panic(err)
		}
		fmt.Println("result:", buf.String())

		tmpl, err = emailTmpl.Parse(str2)
		if err != nil {
			panic(err)
		}
		fmt.Println("result:", buf.String())
	*/
}
