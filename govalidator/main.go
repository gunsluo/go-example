package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gunsluo/govalidator"
	"github.com/gunsluo/govalidator/custom"
)

func init() {
	govalidator.CustomTypeTagMap.Set("max", govalidator.CustomTypeValidator(custom.MaxCustomTypeTagFn))
	govalidator.CustomTypeTagMap.Set("min", govalidator.CustomTypeValidator(custom.MinCustomTypeTagFn))
	govalidator.CustomTypeTagMap.Set("gt", govalidator.CustomTypeValidator(custom.GtCustomTypeTagFn))
	govalidator.CustomTypeTagMap.Set("gte", govalidator.CustomTypeValidator(custom.MinCustomTypeTagFn))
	govalidator.CustomTypeTagMap.Set("lt", govalidator.CustomTypeValidator(custom.LtCustomTypeTagFn))
	govalidator.CustomTypeTagMap.Set("lte", govalidator.CustomTypeValidator(custom.MaxCustomTypeTagFn))
}

// UserLoginVO 仅用于用户登录的展示层对象
type UserLoginVO struct {
	UserName string `json:"userName" valid:"required,int,max=1,lte=1"` //用户名
	Password string `json:"password" valid:"required,int"`             //用户密码
	Count    int    `json:"count" valid:"max=1"`                       //
}

func (user *UserLoginVO) Validate() error {

	if res, err := govalidator.ValidateStruct(user); !res {
		return user.customizeValidationErr(err)
	}

	return nil
}

func (user *UserLoginVO) customizeValidationErr(err error) error {

	if _, ok := err.(*govalidator.UnsupportedTypeError); ok {
		return nil
	}

	var errs []string
	for _, ve := range err.(govalidator.Errors) {

		e, ok := ve.(govalidator.Error)
		if !ok {
			continue
		}
		switch e.Name {
		case "userName":
			if e.Tag == "required" {
				errs = append(errs, "please input user name.")
			}
			if e.Tag == "int" {
				errs = append(errs, "user name is number.")
			}
			if e.Tag == "max" {
				errs = append(errs, fmt.Sprintf("user name[%s] max len is 1.", e.Value))
			}
			if e.Tag == "lte" {
				errs = append(errs, fmt.Sprintf("user name[%s] less than equal 1.", e.Value))
			}
		case "password":
			errs = append(errs, "password is incorrect")
		case "count":
			errs = append(errs, "count max 1")
		}
	}

	if len(errs) == 0 {
		return err
	}

	return errors.New(strings.Join(errs, ";"))
}

func main() {

	fmt.Println(govalidator.StringMatches("luojiYY-123", "^[a-zA-Z0-9-]*$"))

	fmt.Println(govalidator.ValidateVar("", "required"))
	fmt.Println(govalidator.ValidateVar("a", "required,int"))
	fmt.Println(govalidator.ValidateVar("a%", "required,matches(^[a-zA-Z0-9-]*$)"))

	str := "abc"
	fmt.Println(govalidator.ValidateVar(&str, "required,int"))

	fmt.Println(govalidator.ValidateVar(map[string]string{"c": "1", "b": "b"}, "required,int"))

	fmt.Println(govalidator.ValidateVar([]string{"1", "b", "c"}, "required,int"))
	fmt.Println(govalidator.ValidateVar([3]string{"1", "2", "c"}, "required,int"))

	user := &UserLoginVO{
		UserName: "luoji-",
		Password: "YYhd123%-",
		Count:    10,
	}

	if err := user.Validate(); err != nil {
		println("error: " + err.Error())
	}

	fmt.Println(govalidator.ValidateVar([]UserLoginVO{*user}, "required"))
	fmt.Println(govalidator.ValidateVar([]*UserLoginVO{user}, "required"))
}
