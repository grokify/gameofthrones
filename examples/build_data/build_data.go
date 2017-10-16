package main

import (
	"fmt"

	"github.com/grokify/gameofthrones"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/strconv/phonenumber"
	"github.com/grokify/oauth2util-go/scimutil"
)

type Person struct {
	AdditionalName string `json:"additionalName,omitempty"`
	GivenName      string `json:"givenName,omitempty"`
	FamilyName     string `json:"familyName,omitempty"`
	DisplayName    string `json:"displayName,omitempty"`
}

func main() {
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
		char.Character.PhoneNumbers = append(char.Character.PhoneNumbers,
			scimutil.Item{Value: fmt.Sprintf("+%d", num)})
		chars[i] = char
	}
	fmtutil.PrintJSON(chars)
}
