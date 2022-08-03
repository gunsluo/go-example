package pkce

import (
	"testing"
)

func TestPKEC(t *testing.T) {
	cases := []struct {
		method string
		pass   bool
	}{
		{
			method: CodeChallengeMethodSHA1,
			pass:   true,
		},
		{
			method: CodeChallengeMethodSHA256,
			pass:   true,
		},
		{
			method: CodeChallengeMethodSHA512,
			pass:   true,
		},
		{
			method: CodeChallengeMethodMD5,
			pass:   true,
		},
		{
			method: "",
			pass:   true,
		},
		{
			method: "mock",
			pass:   false,
		},
	}

	for _, c := range cases {
		suite, err := Generate(WithCodeChallengeMethod(c.method))
		isPassed := err == nil
		if isPassed != c.pass {
			t.Fatalf("failed to generate suite, got %v, expect %v.", isPassed, c.pass)
		}

		if suite == nil {
			return
		}

		if c.method != "" && suite.CodeChallengeMethod != c.method {
			t.Fatalf("incorrect method, %s.", suite.CodeChallengeMethod)
		}
	}
}
