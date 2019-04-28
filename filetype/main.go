package main

import (
	"fmt"
	"io/ioutil"

	"github.com/h2non/filetype"
)

func main() {
	buf, err := ioutil.ReadFile("sample.jpeg")
	if err != nil {
		panic(err)
	}

	kind, err := filetype.Match(buf)
	if err != nil {
		panic(err)
	}
	if kind == filetype.Unknown {
		fmt.Println("Unknown file type")
		return
	}

	fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)
}
