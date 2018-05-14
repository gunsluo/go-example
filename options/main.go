package main

import "fmt"

type Options struct {
	DefaultSubject string
	Source         string
	Action         string
}

type Option func(*Options)

func NewOptions(opt ...Option) *Options {
	options := new(Options)
	for _, o := range opt {
		o(options)
	}
	return options
}

func setDefaultSubject(subject string) Option {
	return func(o *Options) {
		o.DefaultSubject = subject
	}
}

func setAction(action string) Option {
	return func(o *Options) {
		o.Action = action
	}
}

func main() {
	opt1 := setDefaultSubject("abc")
	opt2 := setAction("test")

	op := NewOptions(opt1, opt2)

	fmt.Println(op.DefaultSubject, op.Source, op.Action)
}
