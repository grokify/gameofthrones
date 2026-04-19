package gameofthrones

import (
	"github.com/grokify/goauth/scim"
)

// Organizations is a list of all Game of Thrones organizations (houses, groups, locations).
var Organizations = []string{
	"Baelish Keep",
	"Bear Island",
	"Blackhaven",
	"Braavos",
	"Casterly Rock",
	"Dorne",
	"Dothraki",
	"Dragonstone",
	"Free Folk",
	"Greywater Watch",
	"Highgarden",
	"Horn Hill",
	"House Clegane",
	"Iron Islands",
	"King's Landing",
	"Lord of Light",
	"Myr",
	"Night's Watch",
	"Order of Maesters",
	"Second Sons",
	"Tarth",
	"The Crossroads",
	"The Dreadfort",
	"The Neck",
	"The Riverlands",
	"The Sparrows",
	"The Stormlands",
	"The Vale",
	"Volantis",
	"Westeros",
	"Winterfell"}

// FamilyNameToOrganization maps character family names to their affiliated organization.
var FamilyNameToOrganization = map[string]string{
	"of Myr":     "Myr",
	"of Tarth":   "Tarth",
	"Aemon":      "Night's Watch",
	"Baelish":    "Baelish Keep",
	"Baratheon":  "King's Landing",
	"Bolton":     "The Dreadfort",
	"Cassel":     "Winterfell",
	"Clegane":    "House Clegane",
	"Dondarrion": "Blackhaven",
	"Drogo":      "Dothraki",
	"Florent":    "Dragonstone",
	"Giantsbane": "Free Folk",
	"Greyjoy":    "Iron Islands",
	"H'ghar":     "Braavos",
	"Lannister":  "Casterly Rock",
	"Luwin":      "Winterfell",
	"Maegyr":     "Volantis",
	"Marsh":      "The Neck",
	"Martell":    "Dorne",
	"Mormont":    "Bear Island",
	"Naharis":    "Second Sons",
	"Payne":      "King's Landing",
	"Pie":        "The Crossroads",
	"Pycelle":    "Order of Maesters",
	"Reed":       "Greywater Watch",
	"Sam":        "Horn Hill",
	"Sand":       "Dorne",
	"Seaworth":   "King's Landing",
	"Selmy":      "King's Landing",
	"Slynt":      "King's Landing",
	"Snow":       "Winterfell",
	"Sparrow":    "The Sparrows",
	"Stark":      "Winterfell",
	"Targaryen":  "Dragonstone",
	"Tarly":      "Horn Hill",
	"Thorne":     "Night's Watch",
	"Tollett":    "Night's Watch",
	"Trant":      "King's Landing",
	"Tyrell":     "Highgarden",
	"Waif":       "Braavos",
	"Worm":       "Dragonstone",
	"Yarwyck":    "Night's Watch",
}

// Organization represents a Game of Thrones organization using Schema.org structure.
type Organization struct {
	Thing
}

// Thing represents a Schema.org Thing with a name.
type Thing struct {
	Name string `json:"name,omitempty"`
}

// GetOrganizationForUser returns the organization for a user based on their family name.
func GetOrganizationForUser(user scim.User) Organization {
	familyName := user.Name.FamilyName
	org := Organization{Thing: Thing{}}
	if orgName, ok := FamilyNameToOrganization[familyName]; ok {
		org.Name = orgName
	}
	return org
}
