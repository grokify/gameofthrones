package gameofthrones

import (
	"github.com/grokify/oauth2util-go/scimutil"
)

var Organizations = []string{
	"Baelish Keep",
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
	"Myr",
	"Night's Watch",
	"Second Sons",
	"Tarth",
	"The Dreadfort",
	"The Lord of Light",
	"The Neck",
	"The Riverlands",
	"The Stormlands",
	"The Vale",
	"Volantis",
	"Westeros",
	"Winterfell"}

var FamilyNameToOrganization = map[string]string{
	"of Myr":     "Myr",
	"of Tarth":   "Tarth",
	"Aemon":      "Night's Watch",
	"Baelish":    "Baelish Keep",
	"Baratheon":  "King's Landing",
	"Bolton":     "The Dreadfort",
	"Cassel":     "Winterfell",
	"Clegane":    "House Clegane",
	"Drogo":      "Dothraki",
	"Giantsbane": "Free Folk",
	"Greyjoy":    "Iron Islands",
	"Lannister":  "Casterly Rock",
	"Luwin":      "Winterfell",
	"Maegyr":     "Volantis",
	"Marsh":      "The Neck",
	"Martell":    "Dorne",
	"Mormont":    "Bear Island",
	"Naharis":    "Second Sons",
	"Reed":       "Greywater Watch",
	"Sand":       "Dorne",
	"Seaworth":   "King's Landing",
	"Selmy":      "King's Landing",
	"Snow":       "Winterfell",
	"Stark":      "Winterfell",
	"Targaryen":  "Dragonstone",
	"Tarly":      "Horn Hill",
	"Tollett":    "Night's Watch",
	"Tyrell":     "Highgarden",
}

// Schema.org Organization
type Organization struct {
	Thing
}

// Schema.org Thing
type Thing struct {
	Name string `json:"name,omitempty"`
}

func GetOrganizationForUser(user scimutil.User) Organization {
	familyName := user.Name.FamilyName
	org := Organization{Thing: Thing{}}
	if orgName, ok := FamilyNameToOrganization[familyName]; ok {
		org.Name = orgName
	}
	return org
}
