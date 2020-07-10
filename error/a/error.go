package main

import (
	"errors"
	"fmt"
)

// NewDeniedError returns a access denied error
func NewDeniedError(message string, err error) error {
	return errDenied{
		message: message,
		err:     err,
	}
}

type errDenied struct {
	subject  string
	resource string
	action   string

	message string
	err     error
}

// Error gets error message
func (e errDenied) Error() string {
	msg := e.message
	if msg == "" {
		msg = "access denied"
	}
	if e.err == nil {
		return msg
	}
	return fmt.Sprintf("%s: %s", msg, e.err)
}

func (errDenied) ErrDenied() {}

// IsErrDenied checks if the error is an denied error
func IsErrDenied(err error) bool {
	var target errDenied
	return errors.As(err, &target)
}

func main() {
	err := errors.New("deny")
	err = NewDeniedError("verify", err)
	err = fmt.Errorf("failed to verify %w", err)

	//fmt.Printf("%T\n", err)
	//fmt.Printf("---%v\n", errors.Unwrap(err))
	fmt.Println("-->", IsErrDenied(err))
	//err := perrors.Wrap(oerr, "tow")
	//err = perrors.Cause(err)

	/*
		fmt.Printf("%T\n", err)
		//err = errors.Unwrap(err)
		fmt.Printf("===%T\n", err)

		var target errDenied
		ok := errors.As(err, &target)
		//ok := errors.Is(err, oerr)
		fmt.Println("-->", ok)
	*/
}
