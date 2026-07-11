package augmented

import (
	"fmt"

	"github.com/grokify/mogo/net/urlutil"
)

const (
	// DefaultFallbackDomain is used for characters without an organization domain.
	DefaultFallbackDomain = "westeros.com"
)

// formatE164 formats a raw phone number as E.164.
func formatE164(number uint64) string {
	if number == 0 {
		return ""
	}
	return fmt.Sprintf("+%d", number)
}

// DomainForOrganization returns the appropriate domain for an organization.
// Non-profit organizations use .org, others use .com.
func DomainForOrganization(orgName string) string {
	// Organizations that should use .org
	orgDomains := map[string]bool{
		"Free Folk":         true,
		"Night's Watch":     true,
		"The Lord of Light": true,
		"Order of Maesters": true,
		"The Sparrows":      true,
	}

	domainPart := urlutil.ToSlugLowerString(orgName)
	if orgDomains[orgName] {
		return fmt.Sprintf("%s.org", domainPart)
	}
	return fmt.Sprintf("%s.com", domainPart)
}
