package cors

import (
	"testing"
)

func TestDomain(t *testing.T) {
	regions := []string{"*"}
	matcher := newDomainMatch(regions)

	for _, c := range []struct {
		domain string
		pass   bool
	}{
		{domain: "a.test.com", pass: true},
		{domain: "b.test.com", pass: true},
	} {

		expected := c.pass
		actual := matcher.IsAllowed(c.domain)
		if expected != actual {
			t.Fatalf("Not equal: \n"+
				"domain: %v\n"+
				"expected: %v\n"+
				"actual  : %v", c.domain, expected, actual)
		}
	}
}

func TestDomain2(t *testing.T) {
	regions := []string{"dev.test.com"}
	matcher := newDomainMatch(regions)

	for _, c := range []struct {
		domain string
		pass   bool
	}{
		{domain: "a.test.com", pass: false},
		{domain: "b.test.com", pass: false},
		{domain: "dev.test.com", pass: true},
	} {

		expected := c.pass
		actual := matcher.IsAllowed(c.domain)
		if expected != actual {
			t.Fatalf("Not equal: \n"+
				"domain: %v\n"+
				"expected: %v\n"+
				"actual  : %v", c.domain, expected, actual)
		}
	}
}

func TestDomain3(t *testing.T) {
	regions := []string{"*.test.com"}
	matcher := newDomainMatch(regions)

	for _, c := range []struct {
		domain string
		pass   bool
	}{
		{domain: "a.test.com", pass: true},
		{domain: "b.test.com", pass: true},
		{domain: "dev.test.com", pass: true},
		{domain: "dev.test.cn", pass: false},
	} {

		expected := c.pass
		actual := matcher.IsAllowed(c.domain)
		if expected != actual {
			t.Fatalf("Not equal: \n"+
				"domain: %v\n"+
				"expected: %v\n"+
				"actual  : %v", c.domain, expected, actual)
		}
	}
}

func TestDomain4(t *testing.T) {
	regions := []string{"www.*.com"}
	matcher := newDomainMatch(regions)

	for _, c := range []struct {
		domain string
		pass   bool
	}{
		{domain: "a.test.com", pass: false},
		{domain: "b.test.com", pass: false},
		{domain: "www.test.com", pass: true},
		{domain: "www.test2.cn", pass: false},
	} {

		expected := c.pass
		actual := matcher.IsAllowed(c.domain)
		if expected != actual {
			t.Fatalf("Not equal: \n"+
				"domain: %v\n"+
				"expected: %v\n"+
				"actual  : %v", c.domain, expected, actual)
		}
	}
}
