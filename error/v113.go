package main

import (
	"errors"
	"fmt"
)

type Error struct {
	Path   string
	Reason string
}

func (e *Error) Error() string {
	return e.Path + ": not found" + " Reason: " + e.Reason
}

var ErrNotFound = errors.New("not found")

func main() {
	{
		e := ErrNotFound
		w := fmt.Errorf("Path %w", e)
		fmt.Println("error:", w)
		fmt.Println("unwrap error:", errors.Unwrap(w))

		// is
		fmt.Println("is:", errors.Is(w, e))
		fmt.Println("is:", errors.Is(e, w))
	}

	{
		e := &Error{Path: "abc", Reason: "bcd"}
		w := fmt.Errorf("Path %w", e)
		fmt.Println("error:", w)
		fmt.Println("unwrap error:", errors.Unwrap(w))

		// is
		fmt.Println("is:", errors.Is(w, e))
		fmt.Println("is:", errors.Is(e, w))

		// new error

		var target *Error
		if errors.As(w, &target) {
			fmt.Println(target.Path)
		}
	}
}
