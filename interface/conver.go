package main

import "fmt"

type inner interface {
	PrintInner()
}

type Outer interface {
	inner
	PrintOuter()
}

type innerAndOuterImpl struct{}

func (io *innerAndOuterImpl) PrintInner() {
	fmt.Println("inner")
}

func (io *innerAndOuterImpl) PrintOuter() {
	fmt.Println("outer")
}

func main() {
	var outer innerAndOuterImpl
	outer.PrintOuter()

	conver(&outer)
}

func conver(outer Outer) {
	var inr inner
	inr = outer
	inr.PrintInner()
}
