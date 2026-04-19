package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/grokify/gameofthrones"
	"github.com/grokify/go-salesforce/sobjects"
	"github.com/grokify/goauth"
	"github.com/grokify/goauth/salesforce"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/net/http/httputilmore"
	"github.com/spf13/cobra"
	"github.com/ttacon/libphonenumber"
)

var salesforceCmd = &cobra.Command{
	Use:   "salesforce",
	Short: "Salesforce CRM integration",
	Long:  `Create and manage Game of Thrones demo data in Salesforce.`,
}

var sfCreateAccountsCmd = &cobra.Command{
	Use:   "create-accounts",
	Short: "Create Salesforce accounts from GoT organizations",
	RunE: func(cmd *cobra.Command, args []string) error {
		sc, err := newSalesforceClient()
		if err != nil {
			return err
		}

		acts, err := getSalesforceAccounts()
		if err != nil {
			return err
		}

		apiURL := sc.URLBuilder.Build("services/data/v41.0/composite/tree/Account/")
		resp, err := sc.ClientMore.PostToJSON(apiURL.String(), acts)
		if err != nil {
			return err
		}

		fmt.Printf("Status: %d\n", resp.StatusCode)
		return httputilmore.PrintResponse(resp, true)
	},
}

var sfDeleteAccountsCmd = &cobra.Command{
	Use:   "delete-accounts",
	Short: "Delete all Salesforce accounts",
	RunE: func(cmd *cobra.Command, args []string) error {
		sc, err := newSalesforceClient()
		if err != nil {
			return err
		}
		return sc.DeleteAccountsAll()
	},
}

var sfCreateContactsCmd = &cobra.Command{
	Use:   "create-contacts",
	Short: "Create Salesforce contacts from GoT characters",
	RunE: func(cmd *cobra.Command, args []string) error {
		sc, err := newSalesforceClient()
		if err != nil {
			return err
		}

		chars, err := gameofthrones.GetDemoCharacters()
		if err != nil {
			return err
		}

		sfActs := getSfAccounts(sc)
		return loadSalesforceCharacters(sc, chars.CharactersSorted(), sfActs)
	},
}

var sfDeleteContactsCmd = &cobra.Command{
	Use:   "delete-contacts",
	Short: "Delete all Salesforce contacts",
	RunE: func(cmd *cobra.Command, args []string) error {
		sc, err := newSalesforceClient()
		if err != nil {
			return err
		}
		return sc.DeleteContactsAll()
	},
}

var sfCreateCasesCmd = &cobra.Command{
	Use:   "create-cases",
	Short: "Create sample Salesforce cases",
	RunE: func(cmd *cobra.Command, args []string) error {
		sc, err := newSalesforceClient()
		if err != nil {
			return err
		}
		return createSalesforceCases(sc)
	},
}

func init() {
	salesforceCmd.AddCommand(sfCreateAccountsCmd)
	salesforceCmd.AddCommand(sfDeleteAccountsCmd)
	salesforceCmd.AddCommand(sfCreateContactsCmd)
	salesforceCmd.AddCommand(sfDeleteContactsCmd)
	salesforceCmd.AddCommand(sfCreateCasesCmd)
}

// Salesforce types and helpers

type sfType struct {
	Type        string `json:"type,omitempty"`
	ReferenceID string `json:"referenceId,omitempty"`
}

type sfAccount struct {
	Attributes        sfType `json:"attributes,omitempty"`
	Name              string `json:"name,omitempty"`
	Phone             string `json:"phone,omitempty"`
	Website           string `json:"website,omitempty"`
	NumberOfEmployees string `json:"numberOfEmployees,omitempty"`
	Industry          string `json:"industry,omitempty"`
	Site              string `json:"site,omitempty"`
}

type sfCreateAccountsRequest struct {
	Records []sfAccount `json:"records,omitempty"`
}

type sfContact struct {
	FirstName string `json:",omitempty"`
	LastName  string `json:",omitempty"`
	Name      string `json:",omitempty"`
	Email     string `json:",omitempty"`
	Phone     string `json:",omitempty"`
	AccountID string `json:",omitempty"`
}

type sfAccounts struct {
	AccountSet  sobjects.AccountSet
	NameToIDMap map[string]string
}

func newSalesforceClient() (salesforce.SalesforceClient, error) {
	_, err := config.LoadDotEnv([]string{os.Getenv("ENV_PATH"), "./.env"}, 1)
	if err != nil {
		return salesforce.SalesforceClient{}, err
	}

	creds := salesforce.OAuth2Credentials{
		CredentialsOAuth2: goauth.CredentialsOAuth2{
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

	return salesforce.NewSalesforceClientPassword(creds)
}

func getSalesforceAccounts() (sfCreateAccountsRequest, error) {
	recs := sfCreateAccountsRequest{Records: []sfAccount{}}
	orgs, err := gameofthrones.GetDemoOrganizations()
	if err != nil {
		return recs, err
	}

	for _, org := range orgs.OrganizationsMap {
		sf := sfAccount{
			Name: org.Name,
			Site: "Headquarters",
			Attributes: sfType{
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

func getSfAccounts(sc salesforce.SalesforceClient) sfAccounts {
	acts, err := sc.GetAccountsAll()
	if err != nil {
		panic(err)
	}
	sfActs := sfAccounts{
		AccountSet:  acts,
		NameToIDMap: map[string]string{},
	}
	for _, rec := range acts.Records {
		sfActs.NameToIDMap[rec.Name] = rec.ID
	}
	return sfActs
}

func loadSalesforceCharacters(sc salesforce.SalesforceClient, chars []gameofthrones.Character, sfActs sfAccounts) error {
	for _, char := range chars {
		contact := sfContact{
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
		}

		if len(char.Organization.Name) > 0 {
			if actID, ok := sfActs.NameToIDMap[char.Organization.Name]; ok {
				contact.AccountID = actID
			}
		}

		resp, err := sc.CreateContact(contact)
		if err != nil {
			return err
		}
		if err := httputilmore.PrintResponse(resp, true); err != nil {
			return err
		}
	}
	return nil
}

func createSalesforceCases(sc salesforce.SalesforceClient) error {
	userinfo, err := sc.UserInfo()
	if err != nil {
		return err
	}

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
		fmt.Printf("Status: %d\n", resp.StatusCode)
		if resp.StatusCode > 399 {
			if err := httputilmore.PrintResponse(resp, true); err != nil {
				return err
			}
		}
	}
	return nil
}
