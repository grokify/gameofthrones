package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/grokify/gameofthrones"
	"github.com/grokify/goauth/scim"
	"github.com/grokify/gophonenumbers"
	"github.com/spf13/cobra"
)

var (
	rcOutput   string
	rcFromFile string
	rcPrettify bool
)

var ringcentralCmd = &cobra.Command{
	Use:   "ringcentral",
	Short: "RingCentral data export",
	Long:  `Export Game of Thrones data in RingCentral format.`,
}

var rcBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build RingCentral contact data",
	Long:  `Generate RingCentral-formatted contact data from GoT characters.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var chars []gameofthrones.Character
		var err error

		if rcFromFile != "" {
			chars, err = loadCharactersFromFile(rcFromFile)
			if err != nil {
				return err
			}
		} else {
			chars = gameofthrones.Characters()
		}

		chars, err = addPhoneNumbers(chars)
		if err != nil {
			return err
		}
		chars = addEmails(chars)

		// Inflate organizations
		for i, char := range chars {
			char.Inflate()
			chars[i] = char
		}

		rcContacts := toRingCentralContacts(chars)

		// Add numeric IDs
		for i := range rcContacts {
			rcContacts[i].ID = ""
			rcContacts[i].IDNum = i + 1
		}

		if rcOutput != "" {
			return writeJSONFile(rcOutput, rcContacts, rcPrettify)
		}

		return printJSON(rcContacts, rcPrettify)
	},
}

func init() {
	rcBuildCmd.Flags().StringVarP(&rcOutput, "output", "o", "", "Output file path (prints to stdout if not specified)")
	rcBuildCmd.Flags().StringVar(&rcFromFile, "from-file", "", "Load characters from JSON file instead of embedded data")
	rcBuildCmd.Flags().BoolVar(&rcPrettify, "pretty", true, "Prettify JSON output")

	ringcentralCmd.AddCommand(rcBuildCmd)
}

// RingCentral types

type rcContact struct {
	ID           string          `json:"id,omitempty"`
	IDNum        int             `json:"_id,omitempty"`
	Name         string          `json:"name,omitempty"`
	Type         string          `json:"type,omitempty"`
	PhoneNumbers []rcPhoneNumber `json:"phoneNumbers,omitempty"`
}

type rcPhoneNumber struct {
	PhoneNumber string `json:"phoneNumber,omitempty"`
	PhoneType   string `json:"phoneType,omitempty"`
}

func toRingCentralContacts(chars []gameofthrones.Character) []rcContact {
	contacts := make([]rcContact, 0, len(chars))
	for _, char := range chars {
		if len(char.Character.PhoneNumbers) == 0 {
			continue
		}
		contacts = append(contacts, rcContact{
			ID:   char.Character.PhoneNumbers[0].Value,
			Name: char.Character.DisplayName,
			Type: "Game of Thrones",
			PhoneNumbers: []rcPhoneNumber{
				{
					PhoneNumber: char.Character.PhoneNumbers[0].Value,
					PhoneType:   "directPhone",
				},
			},
		})
	}
	return contacts
}

func addPhoneNumbers(chars []gameofthrones.Character) ([]gameofthrones.Character, error) {
	a2g := gophonenumbers.NewAreaCodeToGeo()
	if err := a2g.ReadData(); err != nil {
		return chars, err
	}
	fng := gophonenumbers.NewFakeNumberGenerator(a2g.AreaCodes())

	set := map[uint64]int8{}
	for i, char := range chars {
		num, newSet, err := fng.RandomLocalNumberUSUnique(set)
		if err != nil {
			return chars, err
		}
		set = newSet
		e164 := fmt.Sprintf("+%d", num)
		char.Character.PhoneNumbers = append(
			char.Character.PhoneNumbers,
			scim.Item{Value: e164})
		chars[i] = char
	}
	return chars, nil
}

func addEmails(chars []gameofthrones.Character) []gameofthrones.Character {
	for i, char := range chars {
		charName := char.Character.DisplayName
		charSlug := toSlug(charName)
		email := fmt.Sprintf("%s@example.com", charSlug)
		char.Character.Emails = append(
			char.Character.Emails,
			scim.Item{Value: email})
		chars[i] = char
	}
	return chars
}

func toSlug(s string) string {
	// Simple slug conversion
	var result []byte
	for _, c := range []byte(s) {
		if c >= 'a' && c <= 'z' {
			result = append(result, c)
		} else if c >= 'A' && c <= 'Z' {
			result = append(result, c+32) // lowercase
		} else if c >= '0' && c <= '9' {
			result = append(result, c)
		} else if c == ' ' || c == '-' || c == '_' {
			if len(result) > 0 && result[len(result)-1] != '-' {
				result = append(result, '-')
			}
		}
	}
	return string(result)
}

func loadCharactersFromFile(filepath string) ([]gameofthrones.Character, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var chars []gameofthrones.Character
	if err := json.Unmarshal(data, &chars); err != nil {
		return nil, err
	}
	return chars, nil
}

func writeJSONFile(filepath string, v any, prettify bool) error {
	var data []byte
	var err error
	if prettify {
		data, err = json.MarshalIndent(v, "", "  ")
	} else {
		data, err = json.Marshal(v)
	}
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, data, 0600)
}

func printJSON(v any, prettify bool) error {
	var data []byte
	var err error
	if prettify {
		data, err = json.MarshalIndent(v, "", "  ")
	} else {
		data, err = json.Marshal(v)
	}
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
