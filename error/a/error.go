package a

import (
	"fmt"

	"github.com/pkg/errors"
)

type errDenied struct {
	message string
	err     error
}

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

func GetErr() error {
	e := errors.New("test error")
	err := errDenied{message: "", err: e}
	return errors.Wrap(err, "get err")
}
