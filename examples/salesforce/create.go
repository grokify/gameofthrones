package main

import (
	"fmt"

	"github.com/grokify/gameofthrones"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/net/httputilmore"
	"github.com/grokify/oauth2util-go/services/salesforce"
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

func LoadCharacters(sc salesforce.SalesforceClient, urlBuilder salesforce.URLBuilder) {
	//https://developer.salesforce.com/forums/?id=906F0000000ApxUIAS

	chars, err := gameofthrones.ReadCharactersJSON()
	if err != nil {
		panic(err)
	}

	for _, char := range chars {
		contact := Contact{
			FirstName: char.Character.Name.GivenName,
			LastName:  char.Character.Name.FamilyName,
			//	Name:      char.Character.DisplayName,
			Email: char.Character.Emails[0].Value,
			Phone: char.Character.PhoneNumbers[0].Value}
		fmtutil.PrintJSON(char)
		fmtutil.PrintJSON(contact)
		//panic("B")
		resp, err := sc.CreateContact(contact) //cm.PostToJSON(apiURL.String(), contact)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", resp.StatusCode)
		httputilmore.PrintResponse(resp, true)
		//panic("C")
		//break
	}
}

func main() {
	err := config.LoadDotEnv()
	if err != nil {
		panic(err)
	}
	_, err = salesforce.NewClientPasswordSalesforceEnv()
	if err != nil {
		panic(err)
	}

	sc, err := salesforce.NewSalesforceClientEnv()
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
	if 1 == 1 {
		LoadCharacters(sc, sc.URLBuilder)
	}
	fmt.Println("DONE")
}
