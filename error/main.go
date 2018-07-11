package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gunsluo/go-example/error/a"
	"github.com/pkg/errors"
)

func main() {
	_, err := ReadConfig()
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%+v\n", err)
	}

	fmt.Println("--------")
	if IsErrDenied(err) {
		fmt.Println("return err", err)
	} else {
		fmt.Println("don't return err", err)
	}

	fmt.Println("--------")
	err = a.GetErr()
	if IsErrDenied(err) {
		fmt.Println("return err", err)
	} else {
		fmt.Println("don't return err", err)
	}
}

// IsErrDenied checks if the error is an denied error
func IsErrDenied(err error) bool {
	type errDenied interface {
		ErrDenied()
	}

	_, ok := errors.Cause(err).(errDenied)
	return ok
}

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}
	return buf, nil
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home, ".settings.xml"))
	return config, errors.Wrap(err, "could not read config")
}
