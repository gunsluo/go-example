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

func init() {
	emailTmpl = template.New("email")
}

func main() {
	str := "we are {{.CompanyName}} {{.CompanyLocation}} {{.JobSeekerName}} {{.JobTitle}} {{.JobDescription}} {{.InterviewDate}} {{.InterviewLocation}}"
	str2 := "you are {{.CompanyName}} {{.CompanyLocation}} {{.JobSeekerName}} {{.JobTitle}} {{.JobDescription}} {{.InterviewDate}} {{.InterviewLocation}}"
	p := EmailPlaceholder{
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
	tmpl, err = emailTmpl.Parse(str2)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, p)
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", buf.String())
}
