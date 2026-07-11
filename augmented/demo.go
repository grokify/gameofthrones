package augmented

import (
	"fmt"
	"sort"
	"strings"

	got "github.com/grokify/gameofthrones"
	"github.com/grokify/goauth/scim"
	"github.com/grokify/gophonenumbers"
	"github.com/grokify/mogo/net/urlutil"
	"github.com/grokify/mogo/strconv/phonenumber"
)

// Characters holds a map of character display names to their augmented Character data.
type Characters struct {
	CharactersMap map[string]Character
}

// NewCharacters creates a new Characters container.
func NewCharacters() *Characters {
	return &Characters{
		CharactersMap: make(map[string]Character),
	}
}

// LoadCharacters populates the map from a slice of base Characters.
func (c *Characters) LoadCharacters(chars []got.Character) {
	for _, char := range chars {
		c.CharactersMap[char.Character.DisplayName] = Character{Character: char}
	}
}

// Get returns a character by display name.
func (c *Characters) Get(name string) (Character, bool) {
	char, ok := c.CharactersMap[name]
	return char, ok
}

// NamesSorted returns character display names in alphabetical order.
func (c *Characters) NamesSorted() []string {
	names := make([]string, 0, len(c.CharactersMap))
	for name := range c.CharactersMap {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// CharactersSorted returns characters sorted by display name.
func (c *Characters) CharactersSorted() []Character {
	names := c.NamesSorted()
	chars := make([]Character, 0, len(names))
	for _, name := range names {
		if char, ok := c.CharactersMap[name]; ok {
			chars = append(chars, char)
		}
	}
	return chars
}

// Slice returns all characters as a slice.
func (c *Characters) Slice() []Character {
	return c.CharactersSorted()
}

// Count returns the number of characters.
func (c *Characters) Count() int {
	return len(c.CharactersMap)
}

// GetCharacters returns all characters with generated demo data including
// email addresses and phone numbers based on their organization affiliations.
func GetCharacters() (*Characters, error) {
	result := NewCharacters()

	a2g := gophonenumbers.NewAreaCodeToGeo()
	if err := a2g.ReadData(); err != nil {
		return nil, fmt.Errorf("failed to read area code data: %w", err)
	}

	fng := gophonenumbers.NewFakeNumberGenerator(a2g.AreaCodes())
	unique := make(map[uint64]int8) // tracks used phone numbers
	acsOrgs := make(map[uint16]struct{})
	var acsOther []uint16

	aci := phonenumber.NewAreaCodeIncrementor(phoneNumberStart)

	orgs, err := GetOrganizations()
	if err != nil {
		return nil, fmt.Errorf("failed to get organizations: %w", err)
	}

	// Track which area codes are assigned to organizations
	for _, org := range orgs.OrganizationsMap {
		if org.Phone > 0 {
			unique[org.Phone] = 1
		}
		if org.AreaCode > 0 {
			acsOrgs[org.AreaCode] = struct{}{}
		}
	}

	// Get area codes not used by organizations
	for _, ac := range a2g.AreaCodes() {
		if _, ok := acsOrgs[ac]; !ok {
			acsOther = append(acsOther, ac)
		}
	}

	result.LoadCharacters(got.Characters())

	for _, baseChar := range result.CharactersSorted() {
		dispName := strings.TrimSpace(baseChar.Character.Character.DisplayName)
		slug := urlutil.ToSlugLowerString(dispName)

		org := got.GetOrganizationForUser(baseChar.Character.Character)

		augChar := Character{
			Character: baseChar.Character,
		}
		augChar.Character.Organization = org

		if augOrg, ok := orgs.OrganizationsMap[org.Name]; ok {
			// Generate email using organization domain
			if domain := strings.TrimSpace(augOrg.Domain); domain != "" {
				email := fmt.Sprintf("%s@%s", slug, domain)
				augChar.Email = email
				augChar.Character.Character.Emails = append(
					augChar.Character.Character.Emails,
					scim.Item{Value: email})
			}

			// Generate phone using organization area code
			if augOrg.AreaCode > 0 {
				num := aci.GetNext(augOrg.AreaCode)
				e164 := fmt.Sprintf("+%d", num)
				augChar.Phone = e164
				augChar.Character.Character.PhoneNumbers = append(
					augChar.Character.Character.PhoneNumbers,
					scim.Item{Value: e164})
			} else {
				// Use random area code not assigned to an organization
				var genErr error
				num, unique, genErr := fng.RandomLocalNumberUSUniqueAreaCodeSet(unique, acsOther)
				if genErr != nil {
					return nil, fmt.Errorf("failed to generate phone number: %w", genErr)
				}
				e164 := fmt.Sprintf("+%d", num)
				augChar.Phone = e164
				augChar.Character.Character.PhoneNumbers = append(
					augChar.Character.Character.PhoneNumbers,
					scim.Item{Value: e164})
				_ = unique // silence unused warning
			}
		}

		// Fallback email if none was generated
		if augChar.Email == "" {
			email := fmt.Sprintf("%s@%s", slug, DefaultFallbackDomain)
			augChar.Email = email
			augChar.Character.Character.Emails = append(
				augChar.Character.Character.Emails,
				scim.Item{Value: email})
		}

		result.CharactersMap[dispName] = augChar
	}

	return result, nil
}

// GetCharactersByOrganization returns characters grouped by their organization.
func GetCharactersByOrganization() (map[string][]Character, error) {
	chars, err := GetCharacters()
	if err != nil {
		return nil, err
	}

	result := make(map[string][]Character)
	for _, char := range chars.CharactersMap {
		orgName := char.Character.Organization.Name
		if orgName == "" {
			orgName = "Unknown"
		}
		result[orgName] = append(result[orgName], char)
	}

	// Sort characters within each organization
	for orgName := range result {
		sort.Slice(result[orgName], func(i, j int) bool {
			return result[orgName][i].DisplayName() < result[orgName][j].DisplayName()
		})
	}

	return result, nil
}

// GetCharactersByStation returns characters grouped by their station.
func GetCharactersByStation() (map[got.Station][]Character, error) {
	chars, err := GetCharacters()
	if err != nil {
		return nil, err
	}

	result := make(map[got.Station][]Character)
	for _, char := range chars.CharactersMap {
		station, ok := got.CharacterStation[char.DisplayName()]
		if !ok {
			station = got.StationSmallfolk // Default
		}
		result[station] = append(result[station], char)
	}

	// Sort characters within each station
	for station := range result {
		sort.Slice(result[station], func(i, j int) bool {
			return result[station][i].DisplayName() < result[station][j].DisplayName()
		})
	}

	return result, nil
}
