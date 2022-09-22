package main

import (
	"flag"
	"net/url"
	"strconv"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	"bufio"
	"bytes"
	"encoding/base32"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
	"time"
)

func display(key *otp.Key, data []byte) {
	fmt.Printf("Issuer:       %s\n", key.Issuer())
	fmt.Printf("Account Name: %s\n", key.AccountName())
	fmt.Printf("Secret:       %s\n", key.Secret())
	fmt.Printf("Orig:       %s\n", key.String())
	fmt.Println("Writing PNG to qr-code.png....")
	ioutil.WriteFile("qr-code.png", data, 0644)
	fmt.Println("")
	fmt.Println("Please add your TOTP to your OTP Application now!")
	fmt.Println("")
}

func promptForPasscode() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Passcode: ")
	text, _ := reader.ReadString('\n')
	return text
}

// Demo function, not used in main
// Generates Passcode using a UTF-8 (not base32) secret and custom parameters
func GeneratePassCode(utf8string string) string {
	secret := base32.StdEncoding.EncodeToString([]byte(utf8string))
	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA512,
	})
	if err != nil {
		panic(err)
	}
	return passcode
}

func getTOTPValidateOpts(k *otp.Key) totp.ValidateOpts {
	return totp.ValidateOpts{
		Period:    uint(k.Period()),
		Skew:      1,
		Digits:    getDigits(k),
		Algorithm: getAlgorithm(k),
	}
}

func getDigits(k *otp.Key) otp.Digits {
	s := k.String()

	u, err := url.Parse(s)
	if err != nil {
		return otp.DigitsSix
	}

	q := u.Query()
	num, err := strconv.ParseUint(q.Get("digits"), 10, 64)
	if err != nil {
		return otp.DigitsSix
	}

	if num == uint64(otp.DigitsEight) {
		return otp.DigitsEight
	}
	return otp.DigitsSix
}

func getAlgorithm(k *otp.Key) otp.Algorithm {
	s := k.String()

	u, err := url.Parse(s)
	if err != nil {
		return otp.AlgorithmSHA1
	}

	q := u.Query()
	algorithm := q.Get("algorithm")
	switch algorithm {
	case "SHA256":
		return otp.AlgorithmSHA256
	case "SHA512":
		return otp.AlgorithmSHA512
	case "MD5":
		return otp.AlgorithmMD5
	}

	return otp.AlgorithmSHA1
}

func main() {
	var orig string
	flag.StringVar(&orig, "orig", "", "a orig value")
	flag.Parse()

	if orig == "" {
		bind()
	} else {
		validate(orig)
	}
}

func bind() {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Target",
		AccountName: "gunsluo@gmail.com",
		Digits:      otp.DigitsSix,
		Period:      30,
		Algorithm:   otp.AlgorithmSHA256,
	})
	if err != nil {
		panic(err)
	}
	// Convert TOTP key into a PNG
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		panic(err)
	}
	png.Encode(&buf, img)

	// display the QR code to the user.
	display(key, buf.Bytes())

	// Now Validate that the user's successfully added the passcode.
	fmt.Println("Validating TOTP...")
	passcode := promptForPasscode()
	// note: Validate only for AlgorithmSHA1
	// valid := totp.Validate(passcode, key.Secret())
	valid, err := totp.ValidateCustom(passcode, key.Secret(), time.Now().UTC(), getTOTPValidateOpts(key))
	if err != nil {
		panic(err)
	}
	if valid {
		println("Valid passcode!")
		os.Exit(0)
	} else {
		println("Invalid passcode!")
		os.Exit(1)
	}
}

func validate(orig string) {
	key, err := otp.NewKeyFromURL(orig)
	if err != nil {
		panic(err)
	}
	fmt.Println("secret:", key.Secret())

	// Now Validate that the user's successfully added the passcode.
	fmt.Println("Validating TOTP...")
	passcode := promptForPasscode()
	valid, err := totp.ValidateCustom(passcode, key.Secret(), time.Now().UTC(), getTOTPValidateOpts(key))
	if err != nil {
		panic(err)
	}

	if valid {
		println("Valid passcode!")
		os.Exit(0)
	} else {
		println("Invalid passcode!")
		os.Exit(1)
	}
}
