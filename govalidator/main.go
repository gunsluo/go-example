package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gunsluo/govalidator"
)

func init() {
	govalidator.CustomTypeTagMap.Set("max", govalidator.CustomTypeValidator(
		func(i interface{}, ctx interface{}, params ...string) bool {
			if len(params) == 0 {
				return true
			}

			max, err := strconv.Atoi(params[0])
			if err != nil {
				return true
			}

			switch v := i.(type) {
			case int:
				return v <= max
			case string:
				return len(v) <= max
			default:
			}

			return true
		}))
}

// UserLoginVO 仅用于用户登录的展示层对象
type UserLoginVO struct {
	UserName string `json:"userName" valid:"required,int"` //用户名
	Password string `json:"password" valid:"required,int"` //用户密码
	Count    int    `json:"count" valid:"-"`               //
}

func (user *UserLoginVO) Validate() error {

	if res, err := govalidator.ValidateStruct(user); !res {
		//return err
		return user.customizeValidationErr(err)
	}

	return nil
}

func (user *UserLoginVO) customizeValidationErr(err error) error {

	if _, ok := err.(*govalidator.UnsupportedTypeError); ok {
		return nil
	}

	for _, ve := range err.(govalidator.Errors) {

		e, ok := ve.(govalidator.Error)
		if !ok {
			continue
		}
		switch e.Name {
		case "UserName":
			return errors.New("用户名不符合规范")
		case "Password":
			return errors.New("密码不符合规范")
		case "Count":
			return errors.New("11")
		}
	}

	return err
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

	fmt.Println("ok")
}
