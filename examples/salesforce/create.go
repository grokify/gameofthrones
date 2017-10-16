package main

import (
	"fmt"
	"os"

	"github.com/grokify/gameofthrones"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/net/httputilmore"
	ou "github.com/grokify/oauth2util-go"
	"github.com/grokify/oauth2util-go/services/salesforce"
)

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
	FirstName string
	LastName  string
}

func LoadCharacters(cm httputilmore.ClientMore, urlBuilder salesforce.URLBuilder) {
	//https://developer.salesforce.com/forums/?id=906F0000000ApxUIAS
	chars, err := gameofthrones.ReadCharacters()
	if err != nil {
		panic(err)
	}
	fmtutil.PrintJSON(chars)

	apiURL := urlBuilder.Build("/services/data/v40.0/sobjects/Contact/")

	for _, char := range chars {
		contact := Contact{
			FirstName: char.Character.Name.GivenName,
			LastName:  char.Character.Name.FamilyName}
		resp, err := cm.PostToJSON(apiURL.String(), contact)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", resp.StatusCode)
		httputilmore.PrintResponse(resp, true)
		//break
	}
	fmt.Println(apiURL)
}

func main() {
	err := config.LoadDotEnv()
	if err != nil {
		panic(err)
	}

	client, err := salesforce.NewClientPassword(
		ou.ApplicationCredentials{
			ClientID:     os.Getenv("SALESFORCE_CLIENT_ID"),
			ClientSecret: os.Getenv("SALESFORCE_CLIENT_SECRET")},
		ou.UserCredentials{
			Username: os.Getenv("SALESFORCE_USERNAME"),
			Password: fmt.Sprintf("%v%v",
				os.Getenv("SALESFORCE_PASSWORD"),
				os.Getenv("SALESFORCE_SECURITY_KEY"))})

	if err != nil {
		panic(err)
	}

	urlBuilder := salesforce.NewURLBuilder(os.Getenv("SALESFORCE_INSTANCE_NAME"))

	apiURL := urlBuilder.Build("services/data")

	resp, err := client.Get(apiURL.String())
	if err != nil {
		panic(err)
	}

	httputilmore.PrintResponse(resp, true)

	cm := httputilmore.ClientMore{Client: client}

	if 1 == 1 {
		acts := GetAccounts()

		apiURL = urlBuilder.Build("services/data/v34.0/composite/tree/Account/")
		fmt.Println(apiURL)

		resp, err = cm.PostToJSON(apiURL.String(), acts)

		fmt.Printf("%v\n", resp.StatusCode)
		httputilmore.PrintResponse(resp, true)
	}
	if 1 == 0 {
		LoadCharacters(cm, urlBuilder)
	}
	fmt.Println("DONE")
}
