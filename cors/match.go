package cors

import (
	"strings"
)

type wildcard struct {
	prefix string
	suffix string
}

func (w wildcard) match(s string) bool {
	return len(s) >= len(w.prefix)+len(w.suffix) && strings.HasPrefix(s, w.prefix) && strings.HasSuffix(s, w.suffix)
}

type domainMatch struct {
	// Set to true when allowed domains contains a "*"
	allowedDomainsAll bool
	// Normalized list of plain allowed domains
	allowedDomains []string
	// List of allowed domains containing wildcards
	allowedWDomains []wildcard
}

func newDomainMatch(domains []string) *domainMatch {
	m := &domainMatch{}

	for _, domain := range domains {
		domain = strings.ToLower(domain)
		if domain == "*" {
			m.allowedDomainsAll = true
			break
		} else if i := strings.IndexByte(domain, '*'); i >= 0 {
			w := wildcard{domain[0:i], domain[i+1:]}
			m.allowedWDomains = append(m.allowedWDomains, w)
		} else {
			m.allowedDomains = append(m.allowedDomains, domain)
		}
	}

	return m
}

// IsAllowed match ...
func (m *domainMatch) IsAllowed(domain string) bool {
	if m.allowedDomainsAll {
		return true
	}

	for _, d := range m.allowedDomains {
		if d == domain {
			return true
		}
	}

	for _, w := range m.allowedWDomains {
		if w.match(domain) {
			return true
		}
	}

	return false
}
