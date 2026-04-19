package gameofthrones

import (
	"fmt"
	"sort"
	"strings"

	"github.com/grokify/goauth/scim"
	"github.com/grokify/gophonenumbers"
	"github.com/grokify/mogo/net/urlutil"
	"github.com/grokify/mogo/strconv/phonenumber"
)

const (
	// defaultFallbackDomain is used for characters without an organization domain.
	defaultFallbackDomain = "westeros.com"
)

// DemoCharacters holds a map of character display names to their full Character data.
type DemoCharacters struct {
	CharactersMap map[string]Character
}

// LoadCharacters populates the map from a slice of Characters.
func (dc *DemoCharacters) LoadCharacters(chars []Character) {
	for _, char := range chars {
		dc.CharactersMap[char.Character.DisplayName] = char
	}
}

// NamesSorted returns character display names in alphabetical order.
func (dc *DemoCharacters) NamesSorted() []string {
	names := []string{}
	for name := range dc.CharactersMap {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// CharactersSorted returns characters sorted by display name.
func (dc *DemoCharacters) CharactersSorted() []Character {
	names := dc.NamesSorted()
	chars := []Character{}
	for _, name := range names {
		if char, ok := dc.CharactersMap[name]; ok {
			chars = append(chars, char)
		}
	}
	return chars
}

// GetDemoCharacters returns all characters with generated demo data including
// email addresses and phone numbers based on their organization affiliations.
func GetDemoCharacters() (DemoCharacters, error) {
	demoChars := DemoCharacters{CharactersMap: map[string]Character{}}
	a2g := gophonenumbers.NewAreaCodeToGeo()
	err := a2g.ReadData()
	if err != nil {
		return demoChars, err
	}
	fng := gophonenumbers.NewFakeNumberGenerator(a2g.AreaCodes())
	unique := map[uint64]int8{} // uses int8 for compatibility with gophonenumbers API
	acsOrgs := map[uint16]struct{}{}
	acsOther := []uint16{}

	aci := phonenumber.NewAreaCodeIncrementor(phoneNumberStart)

	demoOrgs, err := GetDemoOrganizations()
	if err != nil {
		return demoChars, err
	}

	for _, demoOrg := range demoOrgs.OrganizationsMap {
		if demoOrg.Phone > 0 {
			unique[demoOrg.Phone] = 1
		}
		if demoOrg.AreaCode > 0 {
			acsOrgs[demoOrg.AreaCode] = struct{}{}
		}
	}
	for _, ac := range a2g.AreaCodes() {
		if _, ok := acsOrgs[ac]; !ok {
			acsOther = append(acsOther, ac)
		}
	}

	demoChars.LoadCharacters(Characters())

	for _, char := range demoChars.CharactersSorted() {
		dispName := strings.TrimSpace(char.Character.DisplayName)
		slug := urlutil.ToSlugLowerString(dispName)

		org := GetOrganizationForUser(char.Character)
		char.Organization = org

		if demoOrg, ok := demoOrgs.OrganizationsMap[org.Name]; ok {
			if len(strings.TrimSpace(demoOrg.Domain)) > 0 {
				email := fmt.Sprintf("%v@%v", slug, strings.TrimSpace(demoOrg.Domain))
				char.Character.Emails = append(
					char.Character.Emails,
					scim.Item{Value: email})
			}
			if demoOrg.AreaCode > 0 {
				num := aci.GetNext(demoOrg.AreaCode)
				e164 := fmt.Sprintf("+%d", num)
				char.Character.PhoneNumbers = append(
					char.Character.PhoneNumbers,
					scim.Item{Value: e164})
			} else {
				num := uint64(0)
				num, unique, err = fng.RandomLocalNumberUSUniqueAreaCodeSet(unique, acsOther)
				if err != nil {
					return demoChars, err
				}

				e164 := fmt.Sprintf("+%d", num)
				char.Character.PhoneNumbers = append(
					char.Character.PhoneNumbers,
					scim.Item{Value: e164})
			}
		}
		if len(char.Character.Emails) == 0 {
			email := fmt.Sprintf("%v@%s", slug, defaultFallbackDomain)
			char.Character.Emails = append(
				char.Character.Emails,
				scim.Item{Value: email})
		}

		demoChars.CharactersMap[char.Character.DisplayName] = char
	}

	return demoChars, nil
}
