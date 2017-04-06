package main

import (
	"testing"

	"github.com/gunsluo/go-example/ffjson/m"
)

func BenchmarkMarshalByFFJson(b *testing.B) {

	b.ReportAllocs()
	f := m.Foo{Bar: "test", M: map[int]string{1: "a", 2: "b"}}
	for i := 0; i < b.N; i++ {
		marshalByFFJson(&f)
	}
}

func BenchmarkMarshalByJson(b *testing.B) {

	b.ReportAllocs()
	f := m.Foo2{Bar: "test", M: map[int]string{1: "a", 2: "b"}}
	for i := 0; i < b.N; i++ {
		marshalByJson(&f)
	}
}

func BenchmarkUnmarshalByFFJson(b *testing.B) {

	b.ReportAllocs()
	buf := []byte(`{"Bar":"test","M":{1:"a",2:"b"}}`)
	f := new(m.Foo)
	for i := 0; i < b.N; i++ {
		unmarshalByFFJson(buf, f)
	}
}

func BenchmarkUnmarshalByJson(b *testing.B) {

	b.ReportAllocs()
	buf := []byte(`{"Bar":"test","M":{"1":"a","2":"b"}}`)
	f := new(m.Foo2)
	for i := 0; i < b.N; i++ {
		unmarshalByJson(buf, f)
	}
}
