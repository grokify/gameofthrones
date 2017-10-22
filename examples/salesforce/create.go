package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/grokify/gameofthrones"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/net/httputilmore"
	"github.com/grokify/oauth2util-go/services/salesforce"
	"github.com/ttacon/libphonenumber"
)

// SObject Reference
// https://developer.salesforce.com/docs/atlas.en-us.object_reference.meta/object_reference/sforce_api_objects_list.htm
// Field Types
// https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/field_types.htm

type Type struct {
	Type        string `json:"type,omitempty"`
	ReferenceId string `json:"referenceId,omitempty"`
}

type AccountTreeResp struct {
	ID          string `json:"id,omitempty"`
	ReferenceId string `json:"referenceId,omitempty"`
}

type Account struct {
	Attributes        Type   `json:"attributes,omitempty"`
	Name              string `json:"name,omitempty"`
	Phone             string `json:"phone,omitempty"`
	Website           string `json:"phone,omitempty"`
	NumberOfEmployees string `json:"phone,omitempty"`
	Industry          string `json:"phone,omitempty"`
}

type CreateAccountsRequest struct {
	Records []Account `json:"records,omitempty"`
}

func GetAccounts() CreateAccountsRequest {
	recs := CreateAccountsRequest{
		Records: []Account{
			{Name: "Casterly Rock"},
			{Name: "Dragonstone"},
			{Name: "Highgarden"},
			{Name: "King's Landing"},
			{Name: "The Vale"},
			{Name: "Winterfell"},
		},
	}
	for i, _ := range recs.Records {
		recs.Records[i].Attributes = Type{Type: "Account",
			ReferenceId: fmt.Sprintf("ref%v", i)}
	}
	return recs
}

type Contact struct {
	FirstName string `json:",omitempty"`
	LastName  string `json:",omitempty"`
	Name      string `json:",omitempty"`
	Email     string `json:",omitempty"`
	Phone     string `json:",omitempty"`
}

func LoadCharacters(sc salesforce.SalesforceClient, chars []gameofthrones.Character) {
	//https://developer.salesforce.com/forums/?id=906F0000000ApxUIAS

	for _, char := range chars {
		e164 := char.Character.PhoneNumbers[0].Value

		num, err := libphonenumber.Parse(e164, "US")
		formattedNum := libphonenumber.Format(num, libphonenumber.NATIONAL)
		if err != nil {
			panic(err)
		}

		contact := Contact{
			FirstName: char.Character.Name.GivenName,
			LastName:  char.Character.Name.FamilyName,
			Email:     char.Character.Emails[0].Value,
			Phone:     formattedNum}
		fmtutil.PrintJSON(char)
		fmtutil.PrintJSON(contact)

		resp, err := sc.CreateContact(contact) //cm.PostToJSON(apiURL.String(), contact)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", resp.StatusCode)
		httputilmore.PrintResponse(resp, true)
	}
}

func NewSalesforceClientEnv() (salesforce.SalesforceClient, error) {
	err := config.LoadDotEnv()
	if err != nil {
		return salesforce.SalesforceClient{}, err
	}
	_, err = salesforce.NewClientPasswordSalesforceEnv()
	if err != nil {
		return salesforce.SalesforceClient{}, err
	}
	return salesforce.NewSalesforceClientEnv()
}

func GetCharsJSONInflated() ([]gameofthrones.Character, error) {
	return gameofthrones.ReadCharactersJSON()

	filepath := "github.com/grokify/gameofthrones/examples/build_data/characters_out_inflated.json"
	filepath = path.Join(os.Getenv("GOPATH"), "src", filepath)
	bytes, err := ioutil.ReadFile(filepath)
	chars := []gameofthrones.Character{}
	if err != nil {
		return chars, err
	}
	err = json.Unmarshal(bytes, &chars)
	return chars, err
}

func main() {
	sc, err := NewSalesforceClientEnv()
	if err != nil {
		panic(err)
	}

	resp, err := sc.GetServicesData()
	if err != nil {
		panic(err)
	}

	httputilmore.PrintResponse(resp, true)

	if 1 == 0 {
		acts := GetAccounts()

		apiURL := sc.URLBuilder.Build("services/data/v34.0/composite/tree/Account/")
		fmt.Println(apiURL)

		resp, err = sc.ClientMore.PostToJSON(apiURL.String(), acts)

		fmt.Printf("%v\n", resp.StatusCode)
		httputilmore.PrintResponse(resp, true)
	}
	if 1 == 0 {
		sc.DeleteContactsAll()
	}
	if 1 == 1 {
		chars, err := GetCharsJSONInflated()
		if err != nil {
			panic(err)
		}

		LoadCharacters(sc, chars)
	}
	fmt.Println("DONE")
}
