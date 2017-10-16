package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/grokify/gameofthrones"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/io/ioutilmore"
	"github.com/grokify/gotilla/net/urlutil"
	"github.com/grokify/gotilla/strconv/phonenumber"
	"github.com/grokify/oauth2util-go/scimutil"
)

type Person struct {
	AdditionalName string `json:"additionalName,omitempty"`
	GivenName      string `json:"givenName,omitempty"`
	FamilyName     string `json:"familyName,omitempty"`
	DisplayName    string `json:"displayName,omitempty"`
}

func buildBaseData() {
	chars, err := gameofthrones.ReadCharacters()
	if err != nil {
		panic(err)
	}

	// Add fictitious phone numbers to GOT characters
	a2g := phonenumber.NewAreaCodeToGeo()
	a2g.ReadData()
	fng := phonenumber.NewFakeNumberGenerator(a2g.AreaCodes())

	set := map[int]int{}
	num := 0
	for i, char := range chars {
		num, set = fng.RandomLocalNumberUSUnique(set)
		char.Character.PhoneNumbers = append(
			char.Character.PhoneNumbers,
			scimutil.Item{Value: fmt.Sprintf("+%d", num)})

		chars[i] = char
	}
	fmtutil.PrintJSON(chars)
}

func addEmail(chars []gameofthrones.Character) []gameofthrones.Character {
	for i, char := range chars {
		charName := char.Character.DisplayName
		charSlug := urlutil.ToSlugLowerString(charName)
		email := fmt.Sprintf("%v@example.com", charSlug)
		char.Character.Emails = append(
			char.Character.Emails,
			scimutil.Item{Value: email})
		chars[i] = char
	}
	return chars
}

func main() {
	file := "characters_out.json"
	chars := []gameofthrones.Character{}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &chars)
	if err != nil {
		panic(err)
	}
	chars = addEmail(chars)
	fmtutil.PrintJSON(chars)
	outfile := "characters_out_email.json"
	ioutilmore.WriteJSON(outfile, chars, 644, true)
	//fmt.Println("DONE")
}
