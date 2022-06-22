package main

import (
	"fmt"
	"net/url"

	"github.com/valyala/fasttemplate"
)

func main() {
	test()
}

func test() {
	/*
		template := "http://{{host}}/?foo={{bar}}{{bar}}&q={{query}}&baz={{baz}}"
		t := New(template, "{{", "}}")

		// Substitution map.
		// Since "baz" tag is missing in the map, it will be substituted
		// by an empty string.
		m := map[string]interface{}{
			"host": "google.com",     // string - convenient
			"bar":  []byte("foobar"), // byte slice - the fastest

			// TagFunc - flexible value. TagFunc is called only if the given
			// tag exists in the template.
			"query": TagFunc(func(w io.Writer, tag string) (int, error) {
				return w.Write([]byte(url.QueryEscape(tag + "=world")))
			}),
		}

		s := t.ExecuteString(m)
		fmt.Printf("%s", s)
	*/

	template := "http://{{host}}/?q={{query}}&foo={{bar}}{{bar}}"
	//template := "<html><header></header><body>this is {{bar}}{{bar}}</body></html>"
	t := fasttemplate.New(template, "{{", "}}")
	s := t.ExecuteString(map[string]interface{}{
		"host":  "google.com",
		"query": url.QueryEscape("hello=world"),
		//"query": "hello=world",
		"bar": "foobar",
	})
	fmt.Printf("%s\n", s)
}

func test2() {
	template := "http://${host}/?q=${query}&foo=${bar}${bar}"
	//template := "<html><header></header><body>this is {{bar}}{{bar}}</body></html>"
	t := fasttemplate.New(template, "${", "}")
	s := t.ExecuteString(map[string]interface{}{
		"host":  "google.com",
		"query": url.QueryEscape("hello=world"),
		//"query": "hello=world",
		"bar": "foobar",
	})
	fmt.Printf("%s\n", s)
}
