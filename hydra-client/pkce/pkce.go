package pkce

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"hash"
	"math/rand"
	"strings"
	"time"
)

const (
	CodeChallengeMethodSHA1   = "S1"
	CodeChallengeMethodSHA256 = "S256"
	CodeChallengeMethodSHA512 = "S512"
	CodeChallengeMethodMD5    = "MD5"
)

var supportedAlgorithms = []string{
	CodeChallengeMethodSHA1,
	CodeChallengeMethodSHA256,
	CodeChallengeMethodSHA512,
	CodeChallengeMethodMD5,
}

type Suite struct {
	CodeVerifier        string
	CodeChallenge       string
	CodeChallengeMethod string
}

type Option func(o *options)

type options struct {
	codeChallengeMethod string
}

func WithCodeChallengeMethod(method string) Option {
	return func(o *options) {
		if method != "" {
			o.codeChallengeMethod = method
		}
	}
}

func Generate(opts ...Option) (*Suite, error) {
	o := options{codeChallengeMethod: CodeChallengeMethodSHA256}
	for _, opt := range opts {
		opt(&o)
	}

	method := strings.ToUpper(o.codeChallengeMethod)
	if !contains(supportedAlgorithms, method) {
		return nil, errors.New("code challenge method is not supported")
	}

	codeVerifier := generateCodeVerifier()
	codeChallenge := generateCodeChallenge(method, codeVerifier)
	suite := &Suite{
		CodeChallengeMethod: method,
		CodeVerifier:        codeVerifier,
		CodeChallenge:       codeChallenge,
	}
	return suite, nil
}

func generateCodeVerifier() string {
	length := 32
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length, length)
	for i := 0; i < length; i++ {
		b[i] = byte(r.Intn(255))
	}

	return Base64URLEncode(b)
}

func generateCodeChallenge(method, codeVerifier string) string {
	var h hash.Hash
	switch method {
	case CodeChallengeMethodSHA1:
		h = sha1.New()
	case CodeChallengeMethodSHA256:
		h = sha256.New()
	case CodeChallengeMethodSHA512:
		h = sha512.New()
	case CodeChallengeMethodMD5:
		h = md5.New()
	}
	h.Write([]byte(codeVerifier))
	return Base64URLEncode(h.Sum(nil))
}

func Base64URLEncode(str []byte) string {
	encoded := base64.StdEncoding.EncodeToString(str)
	encoded = strings.Replace(encoded, "+", "-", -1)
	encoded = strings.Replace(encoded, "/", "_", -1)
	encoded = strings.Replace(encoded, "=", "", -1)
	return encoded
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
