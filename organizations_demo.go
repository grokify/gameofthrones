package gameofthrones

import (
	"fmt"

	"github.com/grokify/gotilla/net/urlutil"
	"github.com/grokify/gotilla/sort/sortutil"
	"github.com/grokify/gotilla/strconv/phonenumber"
)

type DemoOrganization struct {
	Name     string
	AreaCode uint16
	Phone    uint64
	Domain   string
}

func (oa *DemoOrganization) E164() string {
	if oa.Phone > 0 {
		return fmt.Sprintf("+%v", oa.Phone)
	}
	return ""
}

type DemoOrganizations struct {
	OrganizationsMap map[string]DemoOrganization
}

func GetDemoOrganizations() DemoOrganizations {
	a2g := phonenumber.NewAreaCodeToGeo()
	a2g.ReadData()

	acs := a2g.AreaCodes()
	sortutil.Uint16s(acs)

	orgs := Organizations

	orgDomains := map[string]int{
		"Free Folk":         1,
		"Night's Watch":     1,
		"The Lord of Light": 1}

	demoOrgs := DemoOrganizations{OrganizationsMap: map[string]DemoOrganization{}}

	fng := phonenumber.NewFakeNumberGenerator(acs)

	for i, orgName := range orgs {
		j := i * 9
		if j >= len(acs) {
			panic("A")
		}
		ac := acs[j]
		fmt.Printf("%v %v %v\n", i, orgName, ac)

		demoOrg := DemoOrganization{
			Name:     orgName,
			AreaCode: ac,
			Phone:    fng.LocalNumberUS(ac, uint16(100)),
		}
		domainPart := urlutil.ToSlugLowerString(orgName)
		if _, ok := orgDomains[orgName]; ok {
			demoOrg.Domain = fmt.Sprintf("%s.org", domainPart)
		} else {
			demoOrg.Domain = fmt.Sprintf("%s.com", domainPart)
		}
		demoOrgs.OrganizationsMap[demoOrg.Name] = demoOrg
	}
	return demoOrgs
}
