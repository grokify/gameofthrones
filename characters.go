package gameofthrones

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/grokify/goauth/scim"
	"github.com/grokify/gocharts/v2/data/table"
	"github.com/grokify/mogo/os/osutil"
)

const (
	PackagePath            = "github.com/grokify/gameofthrones"
	CharactersFileCSV      = "characters.csv"
	CharactersFilepathJSON = "examples/build_data/characters_out_inflated.json"

	// csvColumnCount is the expected number of columns in the characters CSV.
	csvColumnCount = 4
)

// Character represents a Game of Thrones character with actor and character information.
type Character struct {
	Actor        scim.User    `json:"actor,omitempty"`
	Character    scim.User    `json:"character,omitempty"`
	Organization Organization `json:"organization,omitempty"`
}

func buildCharacterDisplayName(u scim.User) string {
	var parts []string
	if len(u.Name.GivenName) > 0 {
		parts = append(parts, u.Name.GivenName)
	}
	if len(u.NickName) > 0 {
		parts = append(parts, fmt.Sprintf("\"%v\"", u.NickName))
	}
	if len(u.Name.FamilyName) > 0 {
		parts = append(parts, u.Name.FamilyName)
	}
	return strings.Join(parts, " ")
}

// NewCharacterSimpleOpts contains options for creating a Character.
type NewCharacterSimpleOpts struct {
	ActorName       string
	GivenName       string
	FamilyName      string
	NickName        string
	AddOrganization bool
}

// NewCharacterSimple creates a Character from the provided options.
func NewCharacterSimple(opts NewCharacterSimpleOpts) Character {
	c := Character{
		Actor: scim.User{DisplayName: strings.TrimSpace(opts.ActorName)},
		Character: scim.User{
			Name: scim.Name{
				GivenName:  strings.TrimSpace(opts.GivenName),
				FamilyName: strings.TrimSpace(opts.FamilyName),
			},
			NickName: strings.TrimSpace(opts.NickName),
		}}
	c.Character.DisplayName = buildCharacterDisplayName(c.Character)
	if opts.AddOrganization {
		c.Organization = GetOrganizationForUser(c.Character)
	}
	return c
}

// Inflate populates the Character's Organization based on their family name.
func (char *Character) Inflate() {
	char.Organization = GetOrganizationForUser(char.Character)
}

// ReadCharactersJSON reads characters from a JSON file. If no path is provided,
// it reads from the default CharactersFilepathJSON location.
func ReadCharactersJSON(filepaths ...string) ([]Character, error) {
	switch len(filepaths) {
	case 0:
		return ReadCharactersPathJSON(GetPackagePath(CharactersFilepathJSON))
	case 1:
		return ReadCharactersPathJSON(filepaths[0])
	default:
		return []Character{}, errors.New("too many file paths, only 0 or 1 allowed")
	}
}

// ReadCharactersPathJSON reads characters from the specified JSON file path.
func ReadCharactersPathJSON(filepath string) ([]Character, error) {
	chars := []Character{}
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return chars, err
	}
	err = json.Unmarshal(bytes, &chars)
	return chars, err
}

// GetPackagePath returns the full path to a file within the package directory.
func GetPackagePath(pathPart string) string {
	return osutil.GoPath("src", PackagePath, pathPart)
	// return path.Join(os.Getenv("GOPATH"), "src", PackagePath, pathPart)
}

//go:embed characters.csv
var charactersDataBytes []byte

// Characters returns all Game of Thrones characters from the embedded CSV data.
// This function panics if the embedded data cannot be parsed, which indicates
// a build-time error that should never occur in normal operation.
func Characters() []Character {
	var chars []Character
	tbl, err := table.ParseReadSeeker(&table.ParseOptions{
		TrimSpace: true,
	}, bytes.NewReader(charactersDataBytes))
	if err != nil {
		panic("failed to parse embedded characters CSV: " + err.Error())
	}
	for _, row := range tbl.Rows {
		if len(row) != csvColumnCount {
			continue
		}
		chars = append(chars, NewCharacterSimple(NewCharacterSimpleOpts{
			ActorName:       row[0],
			GivenName:       row[1],
			FamilyName:      row[2],
			NickName:        row[3],
			AddOrganization: true,
		}))
	}
	return chars
}
