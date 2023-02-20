package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/grokify/gameofthrones"
	"github.com/grokify/go-salesforce/sobjects"
	"github.com/grokify/goauth/credentials"
	"github.com/grokify/goauth/salesforce"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/mogo/net/http/httputilmore"
	"github.com/ttacon/libphonenumber"
)

// SObject Reference
// https://developer.salesforce.com/docs/atlas.en-us.object_reference.meta/object_reference/sforce_api_objects_list.htm
// Field Types
// https://developer.salesforce.com/docs/atlas.en-us.api.meta/api/field_types.htm

type Type struct {
	Type        string `json:"type,omitempty"`
	ReferenceID string `json:"referenceId,omitempty"`
}

/*
type AccountTreeResp struct {
	ID          string `json:"id,omitempty"`
	ReferenceID string `json:"referenceId,omitempty"`
}
*/

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

func GetAccounts() (CreateAccountsRequest, error) {
	recs := CreateAccountsRequest{Records: []Account{}}
	orgs, err := gameofthrones.GetDemoOrganizations()
	if err != nil {
		return recs, err
	}

	for _, org := range orgs.OrganizationsMap {
		sf := Account{
			Name: org.Name,
			Site: "Headquarters",
			Attributes: Type{
				Type:        "Account",
				ReferenceID: fmt.Sprintf("ref0%v", org.Domain),
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
	return recs, nil
}

type Contact struct {
	FirstName string `json:",omitempty"`
	LastName  string `json:",omitempty"`
	Name      string `json:",omitempty"`
	Email     string `json:",omitempty"`
	Phone     string `json:",omitempty"`
	AccountID string `json:",omitempty"`
}

func LoadCharacters(sc salesforce.SalesforceClient, chars []gameofthrones.Character, sfActs SfAccounts) error {
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
			if err != nil {
				return err
			}
			contact.Phone = libphonenumber.Format(num, libphonenumber.NATIONAL)
			if err != nil {
				return err
			}
		}

		if len(char.Organization.Name) > 0 {
			orgName := char.Organization.Name
			if actID, ok := sfActs.NameToIDMap[orgName]; ok {
				contact.AccountID = actID
			}
		}

		resp, err := sc.CreateContact(contact) //cm.PostToJSON(apiURL.String(), contact)
		if err != nil {
			return err
		}
		// fmt.Printf("%v\n", resp.StatusCode)
		err = httputilmore.PrintResponse(resp, true)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewSalesforceClientEnv() (salesforce.SalesforceClient, error) {
	_, err := config.LoadDotEnv([]string{os.Getenv("ENV_PATH"), "./.env"}, 1)
	if err != nil {
		return salesforce.SalesforceClient{}, err
	}

	o2Creds := salesforce.OAuth2Credentials{
		CredentialsOAuth2: credentials.CredentialsOAuth2{
			Endpoint:     salesforce.Endpoint,
			ClientID:     os.Getenv("SALESFORCE_CLIENT_ID"),
			ClientSecret: os.Getenv("SALESFORCE_CLIENT_SECRET"),
			Username:     os.Getenv("SALESFORCE_USERNAME"),
			Password: strings.Join([]string{
				os.Getenv("SALESFORCE_PASSWORD"),
				os.Getenv("SALESFORCE_SECURITY_TOKEN"),
			}, ""),
		},
		InstanceName: os.Getenv("SALESFORCE_INSTANCE_NAME"),
	}
	/*
		usrCreds := goauth.UserCredentials{
			Username: os.Getenv("SALESFORCE_USERNAME"),
			Password: strings.Join([]string{
				os.Getenv("SALESFORCE_PASSWORD"),
				os.Getenv("SALESFORCE_SECURITY_TOKEN"),
			}, ""),
		}
	*/

	// fmtutil.PrintJSON(o2Creds)
	// fmtutil.PrintJSON(usrCreds)
	return salesforce.NewSalesforceClientPassword(o2Creds)
}

/*
func GetCharsJSONInflated(debug bool) ([]gameofthrones.Character, error) {
	if debug {
		filepath := "github.com/grokify/gameofthrones/examples/build_data/characters_out_inflated.json"
		filepath = path.Join(os.Getenv("GOPATH"), "src", filepath)
		bytes, err := os.ReadFile(filepath)
		chars := []gameofthrones.Character{}
		if err != nil {
			return chars, err
		}
		err = json.Unmarshal(bytes, &chars)
		return chars, err
	}
	return gameofthrones.ReadCharactersJSON()
}
*/

type SfAccounts struct {
	AccountSet  sobjects.AccountSet
	NameToIDMap map[string]string
}

func GetSfAccounts(sc salesforce.SalesforceClient) SfAccounts {
	acts, err := sc.GetAccountsAll()
	if err != nil {
		panic(err)
	}
	sfActs := SfAccounts{
		AccountSet:  acts,
		NameToIDMap: map[string]string{}}
	for _, rec := range acts.Records {
		sfActs.NameToIDMap[rec.Name] = rec.ID
	}
	return sfActs
}

func CreateCases(sc salesforce.SalesforceClient) error {
	userinfo, err := sc.UserInfo()
	if err != nil {
		return err
	}
	// fmtutil.PrintJSON(userinfo)

	// fmt.Println("CREATING_CASES")
	cases := map[string]sobjects.Case{
		"Jon Snow": {
			Subject:     "Needs rescue north of The Wall",
			Reason:      "Got trapped trying to catch a wight",
			Priority:    "High",
			IsEscalated: true,
			OwnerID:     userinfo.UserID,
		},
	}

	contacts, err := sc.GetContactsAll()
	if err != nil {
		return err
	}
	for contactName, sfCase := range cases {
		contact, err := contacts.GetContactByName(contactName)
		if err == nil {
			sfCase.ContactID = contact.ID
			sfCase.AccountID = contact.AccountID
		}
		resp, err := sc.CreateSobject("Case", sfCase)
		if err != nil {
			return err
		}
		fmt.Println(resp.StatusCode)
		if resp.StatusCode > 399 {
			err := httputilmore.PrintResponse(resp, true)
			if err != nil {
				return err
			}
		}
	}
	return nil
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
		panic("Please enter one valid action\nUsage: create -action create_accounts|delete_accounts|create_contacts|delete_contacts|create_cases")
	}

	sc, err := NewSalesforceClientEnv()
	logutil.FatalErr(err)

	resp, err := sc.GetServicesData()
	logutil.FatalErr(err)

	logutil.FatalErr(httputilmore.PrintResponse(resp, true))

	switch action {
	case "create_accounts":
		acts, err := GetAccounts()
		logutil.FatalErr(err)

		apiURL := sc.URLBuilder.Build("services/data/v41.0/composite/tree/Account/")

		resp, err = sc.ClientMore.PostToJSON(apiURL.String(), acts)
		logutil.FatalErr(err)

		fmt.Printf("%v\n", resp.StatusCode)
		logutil.FatalErr(httputilmore.PrintResponse(resp, true))
	case "delete_accounts":
		logutil.FatalErr(sc.DeleteAccountsAll())
	case "create_contacts":
		chars, err := gameofthrones.GetDemoCharacters()
		logutil.FatalErr(err)
		logutil.FatalErr(LoadCharacters(sc, chars.CharactersSorted(), GetSfAccounts(sc)))
	case "delete_contacts":
		logutil.FatalErr(sc.DeleteContactsAll())
	case "create_cases":
		logutil.FatalErr(CreateCases(sc))
	}
	fmt.Println("DONE")
}
