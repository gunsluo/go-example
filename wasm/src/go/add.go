package main

import (
	"syscall/js"
)

func add(this js.Value, inputs []js.Value) interface{} {
	return inputs[0].Int() + inputs[1].Int()
}

func main() {
	js.Global().Set("add", js.FuncOf(add))
	<-make(chan bool)
}
