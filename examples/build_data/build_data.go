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
	"github.com/grokify/oauth2more/scim"
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

	set := map[uint64]int8{}
	num := uint64(0)
	for i, char := range chars {
		num, set = fng.RandomLocalNumberUSUnique(set)
		e164 := fmt.Sprintf("+%d", num)

		char.Character.PhoneNumbers = append(
			char.Character.PhoneNumbers,
			scim.Item{Value: e164})

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
			scim.Item{Value: email})
		chars[i] = char
	}
	return chars
}

type RcEvContact struct {
	Id           string            `json:"id,omitempty"`
	Name         string            `json:"name,omitempty"`
	Type         string            `json:"type,omitempty"`
	PhoneNumbers []RcEvPhoneNumber `json:"phoneNumbers,omitempty"`
}

type RcEvPhoneNumber struct {
	PhoneNumber string `json:"phoneNumber,omitempty"`
	PhoneType   string `json:"phoneType,omitempty"`
}

func ToRingCentral(chars []gameofthrones.Character) []RcEvContact {
	rcChars := []RcEvContact{}
	for _, char := range chars {
		rcChars = append(rcChars, RcEvContact{
			Id:   char.Character.PhoneNumbers[0].Value,
			Name: char.Character.DisplayName,
			Type: "Game of Thrones",
			PhoneNumbers: []RcEvPhoneNumber{
				{
					PhoneNumber: char.Character.PhoneNumbers[0].Value,
					PhoneType:   "directPhone",
				},
			},
		})
	}
	return rcChars
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
	ioutilmore.WriteJSON(outfile, chars, 0644, true)

	outfile2 := "characters_out_rcev.json"
	ioutilmore.WriteJSON(outfile2, ToRingCentral(chars), 0644, true)
	outfile3 := "characters_out_rcev2.json"
	ioutilmore.WriteJSON(outfile3, ToRingCentral(chars), 0644, false)

	//fmt.Println("DONE")
}
