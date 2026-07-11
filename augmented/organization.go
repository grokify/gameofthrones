package augmented

import (
	"fmt"

	got "github.com/grokify/gameofthrones"
	"github.com/grokify/gophonenumbers"
	"github.com/grokify/mogo/net/urlutil"
	"github.com/grokify/mogo/sort/sortutil"
	"github.com/grokify/mogo/strconv/phonenumber"
)

const (
	// areaCodeStride is the step size when assigning area codes to organizations.
	areaCodeStride = 9

	// phoneNumberStart is the starting local number suffix for generated phone numbers.
	phoneNumberStart = 100
)

// Organization embeds canonical GoT organization data and adds modern contact fields.
type Organization struct {
	got.Organization
	AreaCode uint16 `json:"area_code,omitempty"`
	Phone    uint64 `json:"phone,omitempty"`
	Domain   string `json:"domain,omitempty"`
}

// NewOrganization creates an augmented Organization from a base Organization.
func NewOrganization(base got.Organization, areaCode uint16, phone uint64, domain string) Organization {
	return Organization{
		Organization: base,
		AreaCode:     areaCode,
		Phone:        phone,
		Domain:       domain,
	}
}

// E164 returns the organization's phone number in E.164 format.
func (o *Organization) E164() string {
	return formatE164(o.Phone)
}

// Organizations holds a map of organization names to their augmented data.
type Organizations struct {
	OrganizationsMap map[string]Organization
}

// NewOrganizations creates a new Organizations container.
func NewOrganizations() *Organizations {
	return &Organizations{
		OrganizationsMap: make(map[string]Organization),
	}
}

// Get returns an organization by name.
func (orgs *Organizations) Get(name string) (Organization, bool) {
	org, ok := orgs.OrganizationsMap[name]
	return org, ok
}

// Names returns all organization names.
func (orgs *Organizations) Names() []string {
	names := make([]string, 0, len(orgs.OrganizationsMap))
	for name := range orgs.OrganizationsMap {
		names = append(names, name)
	}
	return names
}

// GetOrganizations returns all organizations with generated demo data including
// area codes, phone numbers, and domain names.
func GetOrganizations() (*Organizations, error) {
	result := NewOrganizations()

	a2g := gophonenumbers.NewAreaCodeToGeo()
	if err := a2g.ReadData(); err != nil {
		return nil, fmt.Errorf("failed to read area code data: %w", err)
	}

	acs := a2g.AreaCodes()
	sortutil.Slice(acs)

	orgs := got.Organizations

	// Organizations that should use .org
	orgDomains := map[string]bool{
		"Free Folk":         true,
		"Night's Watch":     true,
		"The Lord of Light": true,
		"Order of Maesters": true,
		"The Sparrows":      true,
	}

	fng := phonenumber.NewFakeNumberGenerator(acs)

	for i, orgName := range orgs {
		j := i * areaCodeStride
		if j >= len(acs) {
			return nil, fmt.Errorf("area code index out of bounds: need %d but only have %d area codes", j, len(acs))
		}
		ac := acs[j]

		domainPart := urlutil.ToSlugLowerString(orgName)
		var domain string
		if orgDomains[orgName] {
			domain = fmt.Sprintf("%s.org", domainPart)
		} else {
			domain = fmt.Sprintf("%s.com", domainPart)
		}

		augOrg := Organization{
			Organization: got.Organization{
				Thing: got.Thing{Name: orgName},
			},
			AreaCode: ac,
			Phone:    fng.LocalNumberUS(ac, phoneNumberStart),
			Domain:   domain,
		}
		result.OrganizationsMap[orgName] = augOrg
	}

	return result, nil
}
