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
)

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

type NewCharacterSimpleOpts struct {
	ActorName       string
	GivenName       string
	FamilyName      string
	NickName        string
	AddOrganization bool
}

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

func (char *Character) Inflate() {
	char.Organization = GetOrganizationForUser(char.Character)
}

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

func ReadCharactersPathJSON(filepath string) ([]Character, error) {
	chars := []Character{}
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return chars, err
	}
	err = json.Unmarshal(bytes, &chars)
	return chars, err
}

/*
func ReadCharactersCSV(filepaths ...string) ([]Character, error) {
	switch len(filepaths) {
	case 0:
		return ReadCharactersPathCSV(GetCharacterPathCSV())
	case 1:
		return ReadCharactersPathCSV(filepaths[0])
	default:
		return []Character{}, errors.New("too many file paths, only 0 or 1 allowed")
	}
}

func ReadCharacters() ([]Character, error) {
	return ReadCharactersPathCSV(GetCharacterPathCSV())
}

func ReadCharactersPathCSV(filepath string) ([]Character, error) {
	chars := []Character{}
	csv, file, err := csvutil.NewReaderFile(filepath, ',')
	if err != nil {
		return chars, err
	}

	idx := 0
	for {
		rec, errx := csv.Read()
		if errx == io.EOF {
			break
		} else if errx != nil {
			err = errx
			break
		} else if len(rec) < 2 {
			err = fmt.Errorf("bad data [%v]", rec)
			break
		}
		idx++
		if idx == 1 {
			continue
		}
		chars = append(chars, NewCharacterSimple(NewCharacterSimpleOpts{
			ActorName:       rec[0],
			GivenName:       rec[1],
			FamilyName:      rec[2],
			NickName:        rec[3],
			AddOrganization: true,
		}))
	}
	file.Close()
	if err != nil {
		return chars, err
	}
	return chars, nil
}

func GetCharacterPathCSV() string {
	return osutil.GoPath("src", PackagePath, CharactersFileCSV)
	// return path.Join(os.Getenv("GOPATH"), "src", PackagePath, CharactersFileCSV)
}
*/

func GetPackagePath(pathPart string) string {
	return osutil.GoPath("src", PackagePath, pathPart)
	// return path.Join(os.Getenv("GOPATH"), "src", PackagePath, pathPart)
}

//go:embed characters.csv
var charactersDataBytes []byte

func Characters() []Character {
	var chars []Character
	tbl, err := table.ParseReadSeeker(&table.ParseOptions{
		TrimSpace: true,
	}, bytes.NewReader(charactersDataBytes))
	if err != nil {
		panic(err)
	}
	for _, row := range tbl.Rows {
		if len(row) != 4 {
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
