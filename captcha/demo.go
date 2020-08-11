// example of HTTP server that uses the captcha package.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mojocn/base64Captcha"
)

//configJsonBody json request body.
type configJsonBody struct {
	Id            string                       `json:"id"`
	CaptchaType   string                       `json:"captchaType"`
	VerifyValue   string                       `json:"verifyValue"`
	DriverAudio   *base64Captcha.DriverAudio   `json:"-"`
	DriverString  *base64Captcha.DriverString  `json:"-"`
	DriverChinese *base64Captcha.DriverChinese `json:"-"`
	DriverMath    *base64Captcha.DriverMath    `json:"-"`
	DriverDigit   *base64Captcha.DriverDigit   `json:"-"`
}

var store = base64Captcha.DefaultMemStore

// base64Captcha create http handler
func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	//parse request parameters
	decoder := json.NewDecoder(r.Body)
	var param configJsonBody
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	var driver base64Captcha.Driver

	//create base64 encoding captcha
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		param.DriverString = base64Captcha.NewDriverString(100, 240, 0, 3, 6, "1234567890qwertyuioplkjhgfdsazxcvbnm", nil, nil)
		driver = param.DriverString.ConvertFonts()
	case "math":
		param.DriverMath = base64Captcha.NewDriverMath(60, 240, 0, 3, nil, nil)
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}

// base64Captcha verify http handler
func captchaVerifyHandle(w http.ResponseWriter, r *http.Request) {

	//parse request json body
	decoder := json.NewDecoder(r.Body)
	var param configJsonBody
	err := decoder.Decode(&param)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	//verify the captcha
	body := map[string]interface{}{"code": 0, "msg": "failed"}
	fmt.Println("--->", param.Id, param.VerifyValue)
	if store.Verify(param.Id, param.VerifyValue, true) {
		body = map[string]interface{}{"code": 1, "msg": "ok"}
	}

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(body)
}

//start a net/http server
func main() {
	//serve Vuejs+ElementUI+Axios Web Application
	http.Handle("/", http.FileServer(http.Dir("./static")))

	//api for create captcha
	http.HandleFunc("/api/getCaptcha", generateCaptchaHandler)

	//api for verify captcha
	http.HandleFunc("/api/verifyCaptcha", captchaVerifyHandle)

	fmt.Println("Server is at :8777")
	if err := http.ListenAndServe(":8777", nil); err != nil {
		log.Fatal(err)
	}
}
