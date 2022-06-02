package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/achiku/xml"
)

func main() {
	{
		// action
		url := "http://127.0.0.1:10000/dispatch/soapaction"
		action := "processA"

		req, err := http.NewRequest("POST", url, nil)
		req.Header.Set("soapAction", action)
		if err != nil {
			panic(err)
		}

		c := &http.Client{}
		resp, err := c.Do(req)
		if err != nil {
			panic(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%+v\n", string(body))
	}

	{
		url := "http://127.0.0.1:10000/dispatch/soapbody"

		// body
		x := ProcessARequest{RequestID: "request-a-id"}
		envelope := SOAPEnvelope{
			Body: SOAPBody{
				Content: x,
			},
		}
		buf, err := xml.MarshalIndent(envelope, "", "")
		if err != nil {
			panic(err)
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
		if err != nil {
			panic(err)
		}

		c := &http.Client{}
		resp, err := c.Do(req)
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Response:\n%s\n", body)
	}
}
