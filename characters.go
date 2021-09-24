package gameofthrones

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/grokify/oauth2more/scim"
	"github.com/grokify/simplego/encoding/csvutil"
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
		return []Character{}, errors.New("too many file paths, only 0 or 1 allowed.")
	}
}

func ReadCharactersPathJSON(filepath string) ([]Character, error) {
	chars := []Character{}
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return chars, err
	}
	err = json.Unmarshal(bytes, &chars)
	return chars, err
}

func ReadCharactersCSV(filepaths ...string) ([]Character, error) {
	switch len(filepaths) {
	case 0:
		return ReadCharactersPathCSV(GetCharacterPathCSV())
	case 1:
		return ReadCharactersPathCSV(filepaths[0])
	default:
		return []Character{}, errors.New("too many file paths, only 0 or 1 allowed.")
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

	for {
		rec, errx := csv.Read()
		if errx == io.EOF {
			break
		} else if errx != nil {
			err = errx
			break
		} else if len(rec) < 2 {
			err = errors.New(fmt.Sprintf("bad data: %v\n", rec))
			break
		}
		char := Character{
			Actor:     scim.User{DisplayName: rec[0]},
			Character: scim.User{Name: scim.Name{}}}
		if len(rec) >= 2 {
			char.Character.Name.GivenName = rec[1]
		}
		if len(rec) >= 3 {
			char.Character.Name.FamilyName = rec[2]
		}
		if len(rec) >= 3 {
			char.Character.NickName = rec[3]
		}

		parts := []string{}
		if len(char.Character.Name.GivenName) > 0 {
			parts = append(parts, char.Character.Name.GivenName)
		}
		if len(char.Character.NickName) > 0 {
			parts = append(parts, fmt.Sprintf("\"%v\"", char.Character.NickName))
		}
		if len(char.Character.Name.FamilyName) > 0 {
			parts = append(parts, char.Character.Name.FamilyName)
		}

		char.Character.DisplayName = strings.Join(parts, " ")

		chars = append(chars, char)
	}
	file.Close()
	if err != nil {
		return chars, err
	}
	return chars, nil
}

func GetCharacterPathCSV() string {
	return path.Join(os.Getenv("GOPATH"), "src", PackagePath, CharactersFileCSV)
}

func GetPackagePath(pathPart string) string {
	return path.Join(os.Getenv("GOPATH"), "src", PackagePath, pathPart)
}
