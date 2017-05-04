package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type MyData struct {
	aByte   byte
	aShort  int16
	anInt32 int32
	aSlice  []byte
}

type MyData2 struct {
	aByte       byte
	anotherByte byte
	aShort      int16
	anInt32     int32
	aSlice      []byte
}

type MyData3 struct {
	aByte   byte
	anInt32 int32
	aShort  int16
	aSlice  []byte
}

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func main() {

	align(MyData{})
	align(MyData2{})
	align(MyData3{})

	pointer()
}

func align(d interface{}) {

	// First ask Go to give us some information about the MyData type
	typ := reflect.TypeOf(d)
	fmt.Printf("Struct is %d bytes long\n", typ.Size())
	// We can run through the fields in the structure in order
	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}
}

func pointer() {
	data := MyData{
		aByte:   0x1,
		aShort:  0x0203,
		anInt32: 0x04050607,
		aSlice: []byte{
			0x08, 0x09, 0x0a,
		},
	}

	dataBytes := (*[32]byte)(unsafe.Pointer(&data))
	fmt.Printf("Bytes are %#v\n", dataBytes)

	dataslice := *(*reflect.SliceHeader)(unsafe.Pointer(&data.aSlice))

	fmt.Printf("Slice data is %#v\n",
		(*[3]byte)(unsafe.Pointer(dataslice.Data)))
}
