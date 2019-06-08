package main

import (
	"fmt"

	"golang.org/x/xerrors"
)

/*
func main() {
	err := io.ErrUnexpectedEOF
	//err = errors.Errorf("new err %v", err)
	err = fmt.Errorf("new err %v", err)
	if xerrors.Is(err, io.ErrUnexpectedEOF) {
		fmt.Println("equal")
	}
}
*/

var (
	ErrBase = xerrors.New("a new error")
)

/*
func main() {
	err := xerrors.Errorf("raiseError: %w", ErrBase)
	fmt.Println(ErrBase == ErrBase)                       // 地址相同
	fmt.Println(err == ErrBase)                           // 基于ErrBase包装之后不同
	fmt.Println(xerrors.Is(err, ErrBase))                 // 验证是否为基于ErrBase
	fmt.Println(xerrors.Opaque(err) == err)               // 解除关系链之后不为相同地址
	fmt.Println(xerrors.Is(xerrors.Opaque(err), ErrBase)) // 解除关系链之后无法确定关系
	fmt.Println(xerrors.Unwrap(err) == ErrBase)           // 获取内层错误为原错误，地址相同
}
*/

func main() {
	err := xerrors.Errorf("raiseError: %w", ErrBase)
	err2 := xerrors.Errorf("wrap#01: %w", err)
	err3 := xerrors.Errorf("wrap#02: %w", err2)
	fmt.Println(xerrors.Is(err, ErrBase))
	fmt.Println(xerrors.Is(err3, ErrBase)) // 能够正确识别关系，打印为true

	fmt.Printf("%+v\n", err3)
}
