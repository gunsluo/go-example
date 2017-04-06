package main

import (
	"encoding/json"

	"github.com/gunsluo/go-example/ffjson/m"
	"github.com/pquerna/ffjson/ffjson"
)

func marshalByFFJson(f *m.Foo) []byte {

	buf, err := ffjson.Marshal(&f)
	if err != nil {
		panic(err)
	}

	return buf
}

func marshalByJson(f *m.Foo2) []byte {

	buf, err := json.Marshal(&f)
	if err != nil {
		panic(err)
	}

	return buf
}

func unmarshalByFFJson(buf []byte, f *m.Foo) {

	err := ffjson.Unmarshal(buf, f)
	if err != nil {
		panic(err)
	}
}

func unmarshalByJson(buf []byte, f *m.Foo2) {

	err := json.Unmarshal(buf, f)
	if err != nil {
		panic(err)
	}
}
