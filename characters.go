package gameofthrones

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/grokify/gotilla/encoding/csvutil"

	"github.com/grokify/oauth2util-go/scimutil"
)

const (
	PackagePath    = "github.com/grokify/gameofthrones"
	CharactersFile = "characters.csv"
)

type Character struct {
	Actor     scimutil.User `json:"actor,omitempty"`
	Character scimutil.User `json:"character,omitempty"`
}

func ReadCharacters() ([]Character, error) {
	return ReadCharactersPath(GetCharacterPath())
}

func ReadCharactersPath(filepath string) ([]Character, error) {
	chars := []Character{}
	csv, file, err := csvutil.NewReader(filepath, ',', false)
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
			err = errors.New(fmt.Sprintf("Bad Data: %v\n", rec))
			break
		}
		char := Character{
			Actor:     scimutil.User{DisplayName: rec[0]},
			Character: scimutil.User{Name: scimutil.Name{}}}
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

func GetCharacterPath() string {
	return path.Join(os.Getenv("GOPATH"), "src", PackagePath, CharactersFile)
}
