package main

import (
	"bytes"
	"fmt"
	"text/template"
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

func main() {
	str := "we are {{.CompanyName}} {{.CompanyLocation}} {{.JobSeekerName}} {{.JobTitle}} {{.JobDescription}} {{.InterviewDate}} {{.InterviewLocation}}"
	p := EmailPlaceholder{
		CompanyName:       "demo",
		CompanyLocation:   "chengdu",
		JobSeekerName:     "luoji",
		JobTitle:          "",
		JobDescription:    "none",
		InterviewDate:     time.Now().Format("2006-01-02 15:04:05"),
		InterviewLocation: "chengdu",
	}

	t := template.Must(template.New("test").Parse(str))
	var buf bytes.Buffer
	err := t.Execute(&buf, p)
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", buf.String())
}
