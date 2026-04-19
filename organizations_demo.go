package gameofthrones

import (
	"fmt"

	"github.com/grokify/gophonenumbers"
	"github.com/grokify/mogo/net/urlutil"
	"github.com/grokify/mogo/sort/sortutil"
	"github.com/grokify/mogo/strconv/phonenumber"
)

const (
	// areaCodeStride is the step size when assigning area codes to organizations.
	// This spreads organizations across different geographic regions.
	areaCodeStride = 9

	// phoneNumberStart is the starting local number suffix for generated phone numbers.
	phoneNumberStart = 100
)

// DemoOrganization represents a Game of Thrones organization with demo contact data.
type DemoOrganization struct {
	Name     string
	AreaCode uint16
	Phone    uint64
	Domain   string
}

// E164 returns the organization's phone number in E.164 format (e.g., "+14155551234").
func (oa *DemoOrganization) E164() string {
	if oa.Phone > 0 {
		return fmt.Sprintf("+%v", oa.Phone)
	}
	return ""
}

// DemoOrganizations holds a map of organization names to their demo data.
type DemoOrganizations struct {
	OrganizationsMap map[string]DemoOrganization
}

// GetDemoOrganizations returns all organizations with generated demo data including
// area codes, phone numbers, and domain names.
func GetDemoOrganizations() (DemoOrganizations, error) {
	demoOrgs := DemoOrganizations{OrganizationsMap: map[string]DemoOrganization{}}
	a2g := gophonenumbers.NewAreaCodeToGeo()
	err := a2g.ReadData()
	if err != nil {
		return demoOrgs, err
	}

	acs := a2g.AreaCodes()
	sortutil.Slice(acs)

	orgs := Organizations

	orgDomains := map[string]int{
		"Free Folk":         1,
		"Night's Watch":     1,
		"The Lord of Light": 1}

	fng := phonenumber.NewFakeNumberGenerator(acs)

	for i, orgName := range orgs {
		j := i * areaCodeStride
		if j >= len(acs) {
			panic("area code index out of bounds: not enough area codes for organizations")
		}
		ac := acs[j]

		demoOrg := DemoOrganization{
			Name:     orgName,
			AreaCode: ac,
			Phone:    fng.LocalNumberUS(ac, phoneNumberStart),
		}
		domainPart := urlutil.ToSlugLowerString(orgName)
		if _, ok := orgDomains[orgName]; ok {
			demoOrg.Domain = fmt.Sprintf("%s.org", domainPart)
		} else {
			demoOrg.Domain = fmt.Sprintf("%s.com", domainPart)
		}
		demoOrgs.OrganizationsMap[demoOrg.Name] = demoOrg
	}
	return demoOrgs, nil
}
