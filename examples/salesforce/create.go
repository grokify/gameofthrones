package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/grokify/gameofthrones"
	"github.com/grokify/go-salesforce/sobjects"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/net/httputilmore"
	om "github.com/grokify/oauth2more"
	"github.com/grokify/oauth2more/salesforce"
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
	Website           string `json:"website,omitempty"`
	NumberOfEmployees string `json:"numberOfEmployees,omitempty"`
	Industry          string `json:"industry,omitempty"`
	Site              string `json:"site,omitempty"`
}

type CreateAccountsRequest struct {
	Records []Account `json:"records,omitempty"`
}

func GetAccounts() CreateAccountsRequest {
	orgs := gameofthrones.GetDemoOrganizations()

	recs := CreateAccountsRequest{Records: []Account{}}

	for _, org := range orgs.OrganizationsMap {
		sf := Account{
			Name: org.Name,
			Site: "Headquarters",
			Attributes: Type{
				Type:        "Account",
				ReferenceId: fmt.Sprintf("ref0%v", org.Domain),
			},
		}

		if len(strings.TrimSpace(org.Domain)) > 0 {
			sf.Website = fmt.Sprintf("https://%s", org.Domain)
		}
		e164 := org.E164()
		if len(e164) > 0 {
			num, err := libphonenumber.Parse(e164, "US")
			if err == nil {
				sf.Phone = libphonenumber.Format(num, libphonenumber.NATIONAL)
			}
		}

		recs.Records = append(recs.Records, sf)
	}
	return recs
}

type Contact struct {
	FirstName string `json:",omitempty"`
	LastName  string `json:",omitempty"`
	Name      string `json:",omitempty"`
	Email     string `json:",omitempty"`
	Phone     string `json:",omitempty"`
	AccountId string `json:",omitempty"`
}

func LoadCharacters(sc salesforce.SalesforceClient, chars []gameofthrones.Character, sfActs SfAccounts) {
	//https://developer.salesforce.com/forums/?id=906F0000000ApxUIAS

	for _, char := range chars {
		contact := Contact{
			FirstName: char.Character.Name.GivenName,
			LastName:  char.Character.Name.FamilyName,
			Email:     char.Character.Emails[0].Value,
		}

		if len(char.Character.PhoneNumbers) > 0 {
			e164 := char.Character.PhoneNumbers[0].Value

			num, err := libphonenumber.Parse(e164, "US")
			contact.Phone = libphonenumber.Format(num, libphonenumber.NATIONAL)
			if err != nil {
				panic(err)
			}
		}

		if len(char.Organization.Name) > 0 {
			orgName := char.Organization.Name
			if actId, ok := sfActs.NameToIdMap[orgName]; ok {
				contact.AccountId = actId
			}
		}

		resp, err := sc.CreateContact(contact) //cm.PostToJSON(apiURL.String(), contact)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", resp.StatusCode)
		httputilmore.PrintResponse(resp, true)
	}
}

func NewSalesforceClientEnv() (salesforce.SalesforceClient, error) {
	err := config.LoadDotEnvSkipEmpty(os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		return salesforce.SalesforceClient{}, err
	}

	appCreds := salesforce.ApplicationCredentials{
		ApplicationCredentials: om.ApplicationCredentials{
			ClientID:     os.Getenv("SALESFORCE_CLIENT_ID"),
			ClientSecret: os.Getenv("SALESFORCE_CLIENT_SECRET"),
			Endpoint:     salesforce.Endpoint,
		},
		InstanceName: os.Getenv("SALESFORCE_INSTANCE_NAME"),
	}

	usrCreds := om.UserCredentials{
		Username: os.Getenv("SALESFORCE_USERNAME"),
		Password: strings.Join([]string{
			os.Getenv("SALESFORCE_PASSWORD"),
			os.Getenv("SALESFORCE_SECURITY_TOKEN"),
		}, ""),
	}

	fmtutil.PrintJSON(appCreds)
	fmtutil.PrintJSON(usrCreds)
	return salesforce.NewSalesforceClientPassword(appCreds, usrCreds)
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

type SfAccounts struct {
	AccountSet  sobjects.AccountSet
	NameToIdMap map[string]string
}

func GetSfAccounts(sc salesforce.SalesforceClient) SfAccounts {
	acts, err := sc.GetAccountsAll()
	if err != nil {
		panic(err)
	}
	sfActs := SfAccounts{
		AccountSet:  acts,
		NameToIdMap: map[string]string{}}
	for _, rec := range acts.Records {
		sfActs.NameToIdMap[rec.Name] = rec.Id
	}
	return sfActs
}

func CreateCases(sc salesforce.SalesforceClient) {
	userinfo, err := sc.UserInfo()
	if err != nil {
		panic(err)
	}
	fmtutil.PrintJSON(userinfo)

	fmt.Println("CREATING_CASES")
	cases := map[string]sobjects.Case{
		"Jon Snow": sobjects.Case{
			Subject:     "Needs rescue north of The Wall",
			Reason:      "Got trapped trying to catch a wight",
			Priority:    "High",
			IsEscalated: true,
			OwnerId:     userinfo.UserId,
		},
	}

	contacts, err := sc.GetContactsAll()
	if err != nil {
		panic(err)
	}
	for contactName, sfCase := range cases {
		contact, err := contacts.GetContactByName(contactName)
		if err == nil {
			sfCase.ContactId = contact.Id
			sfCase.AccountId = contact.AccountId
		}
		resp, err := sc.CreateSobject("Case", sfCase)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.StatusCode)
		if resp.StatusCode > 399 {
			httputilmore.PrintResponse(resp, true)
		}
	}
}

var commands = map[string]int{
	"create_accounts": 1,
	"delete_accounts": 1,
	"create_contacts": 1,
	"delete_contacts": 1,
	"create_cases":    1,
}

func main() {
	var action string
	flag.StringVar(&action, "action", "(create_accounts|delete_accounts|create_contacts|delete_contacts|create_cases)", "Action")
	flag.Parse()

	action = strings.ToLower(strings.TrimSpace(action))

	if _, ok := commands[action]; !ok {
		panic(fmt.Sprintf("Please enter one valid action\nUsage: create -action create_accounts|delete_accounts|create_contacts|delete_contacts|create_cases"))
	}

	sc, err := NewSalesforceClientEnv()
	if err != nil {
		panic(err)
	}

	resp, err := sc.GetServicesData()
	if err != nil {
		panic(err)
	}

	httputilmore.PrintResponse(resp, true)

	switch action {
	case "create_accounts":
		acts := GetAccounts()

		apiURL := sc.URLBuilder.Build("services/data/v41.0/composite/tree/Account/")

		resp, err = sc.ClientMore.PostToJSON(apiURL.String(), acts)

		fmt.Printf("%v\n", resp.StatusCode)
		httputilmore.PrintResponse(resp, true)
	case "delete_accounts":
		err := sc.DeleteAccountsAll()
		if err != nil {
			panic(err)
		}
	case "create_contacts":
		chars, err := gameofthrones.GetDemoCharacters()
		if err != nil {
			panic(err)
		}
		LoadCharacters(sc, chars.CharactersSorted(), GetSfAccounts(sc))
	case "delete_contacts":
		sc.DeleteContactsAll()
	case "create_cases":
		CreateCases(sc)
	}
	fmt.Println("DONE")
}
