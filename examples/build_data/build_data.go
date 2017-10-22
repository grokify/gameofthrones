package main

import (
	"encoding/json"
	"errors"
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

func addPhoneNumbers(chars []gameofthrones.Character) []gameofthrones.Character {
	// Add fictitious phone numbers to GOT characters
	a2g := phonenumber.NewAreaCodeToGeo()
	a2g.ReadData()
	fng := phonenumber.NewFakeNumberGenerator(a2g.AreaCodes())

	set := map[int]int{}
	num := 0
	for i, char := range chars {
		num, set = fng.RandomLocalNumberUSUnique(set)
		e164 := fmt.Sprintf("+%d", num)
		/*
			num, err := libphonenumber.Parse(e164, "US")
			formattedNum := libphonenumber.Format(num, libphonenumber.NATIONAL)
			if err != nil {
				panic(err)
			}
		*/
		char.Character.PhoneNumbers = append(
			char.Character.PhoneNumbers,
			scimutil.Item{Value: e164})

		chars[i] = char
	}
	return chars
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
	err := errors.New("")

	if 1 == 1 {
		chars, err = gameofthrones.ReadCharactersCSV()
		if err != nil {
			panic(err)
		}
	}
	if 1 == 0 {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &chars)
		if err != nil {
			panic(err)
		}
	}
	chars = addPhoneNumbers(chars)
	chars = addEmail(chars)
	fmtutil.PrintJSON(chars)
	for i, char := range chars {
		char.Inflate()
		chars[i] = char
	}
	outfile := "characters_out_inflated.json"
	ioutilmore.WriteJSON(outfile, chars, 644, true)
	//fmt.Println("DONE")
}
