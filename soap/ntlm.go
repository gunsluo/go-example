package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"

	httpntlm "github.com/vadimi/go-http-ntlm/v2"
)

func main() {

	// configure http client
	client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			Domain:   "outlook.office365.com",
			User:     "email@exchangedomain",
			Password: "password",
			// Configure RoundTripper if necessary, otherwise DefaultTransport is used
			RoundTripper: &http.Transport{
				// provide tls config
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				// other properties RoundTripper, see http.DefaultTransport
			},
		},
	}

	req, err := http.NewRequest("POST", "https://outlook.office365.com/EWS/Exchange.asmx", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.StatusCode, body)
}
