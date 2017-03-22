package main

import (
	"fmt"
	"log"
	"net/http"
)

type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ch <- req
	fmt.Fprintf(w, "notification sent")
}

func (ch Chan) Accept(req *http.Request) {
	fmt.Printf("accept-->%#v\n", req)
}

func (ch Chan) Close() {
	close(ch)
}

func NewChan() Chan {
	ch := make(Chan)
	//http.Handle(pattern, ch)

	go func() {
		for {
			req, ok := <-ch
			if !ok {
				return
			}

			ch.Accept(req)
		}
	}()

	return ch
}

func main() {
	ch := NewChan()
	http.Handle("/chan", ch)

	/*
		go func() {
			time.Sleep(10 * time.Second)
			ch.Close()
		}()
	*/

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
