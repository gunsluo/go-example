package rbac

import (
	"log"
	"strings"
)

// ArrayRemoveDuplicates removes any duplicated elements in a string array.
func ArrayRemoveDuplicates(s *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *s {
		if !found[x] {
			found[x] = true
			(*s)[j] = (*s)[i]
			j++
		}
	}
	*s = (*s)[:j]
}

// ArrayEquals determines whether two string arrays are identical.
func ArrayEquals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	var found bool
	for i := range a {
		found = false
		for j := range b {
			if a[i] == b[j] {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// LogPrint prints the log.
func LogPrint(v ...interface{}) {
	log.Print(v)
}

// RemoveComments removes the comments starting with # in the text.
func RemoveComments(s string) string {
	pos := strings.Index(s, "#")
	if pos == -1 {
		return s
	}
	return strings.TrimSpace(s[0:pos])
}

// EscapeAssertion escapes the dots in the assertion, because the expression evaluation doesn't support such variable names.
func EscapeAssertion(s string) string {
	s = strings.Replace(s, "r.", "r_", -1)
	s = strings.Replace(s, "p.", "p_", -1)
	return s
}
