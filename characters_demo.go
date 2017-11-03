package gameofthrones

import (
	"fmt"
	"sort"
	"strings"

	"github.com/grokify/gotilla/net/urlutil"
	"github.com/grokify/gotilla/strconv/phonenumber"
	"github.com/grokify/oauth2util/scimutil"
)

type DemoCharacters struct {
	CharactersMap map[string]Character
}

func (dc *DemoCharacters) LoadCharacters(chars []Character) {
	for _, char := range chars {
		dc.CharactersMap[char.Character.DisplayName] = char
	}
}

func (dc *DemoCharacters) NamesSorted() []string {
	names := []string{}
	for name, _ := range dc.CharactersMap {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

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

func GetDemoCharacters() (DemoCharacters, error) {
	a2g := phonenumber.NewAreaCodeToGeo()
	a2g.ReadData()
	fng := phonenumber.NewFakeNumberGenerator(a2g.AreaCodes())
	unique := map[uint64]int8{}
	acsOrgs := map[uint16]int8{}
	acsOther := []uint16{}

	aci := phonenumber.NewAreaCodeIncrementor(100)

	demoChars := DemoCharacters{CharactersMap: map[string]Character{}}
	demoOrgs := GetDemoOrganizations()

	for _, demoOrg := range demoOrgs.OrganizationsMap {
		if demoOrg.Phone > 0 {
			unique[demoOrg.Phone] = int8(1)
		}
		if demoOrg.AreaCode > 0 {
			acsOrgs[demoOrg.AreaCode] = int8(1)
		}
	}
	for _, ac := range a2g.AreaCodes() {
		if _, ok := acsOrgs[ac]; !ok {
			acsOther = append(acsOther, ac)
		}
	}

	chars, err := ReadCharactersCSV()
	if err != nil {
		return demoChars, err
	}

	demoChars.LoadCharacters(chars)

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
					scimutil.Item{Value: email})
			}
			if demoOrg.AreaCode > 0 {
				num := aci.GetNext(demoOrg.AreaCode)
				e164 := fmt.Sprintf("+%d", num)
				char.Character.PhoneNumbers = append(
					char.Character.PhoneNumbers,
					scimutil.Item{Value: e164})
			} else {
				num := uint64(0)
				num, unique = fng.RandomLocalNumberUSUniqueAreaCodeSet(unique, acsOther)

				e164 := fmt.Sprintf("+%d", num)
				char.Character.PhoneNumbers = append(
					char.Character.PhoneNumbers,
					scimutil.Item{Value: e164})
			}
		}
		if len(char.Character.Emails) == 0 {
			email := fmt.Sprintf("%v@westeros.com", slug)
			char.Character.Emails = append(
				char.Character.Emails,
				scimutil.Item{Value: email})
		}

		demoChars.CharactersMap[char.Character.DisplayName] = char
	}

	return demoChars, nil
}
