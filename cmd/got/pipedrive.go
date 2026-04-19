package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/grokify/gameofthrones"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/net/http/httputilmore"
	"github.com/grokify/mogo/net/urlutil"
	"github.com/spf13/cobra"
)

var pipedriveBaseURL string

var pipedriveCmd = &cobra.Command{
	Use:   "pipedrive",
	Short: "Pipedrive CRM integration",
	Long:  `Create and manage Game of Thrones demo data in Pipedrive.`,
}

var pdCreateOrgsCmd = &cobra.Command{
	Use:   "create-orgs",
	Short: "Create Pipedrive organizations from GoT houses",
	RunE: func(cmd *cobra.Command, args []string) error {
		pc, err := newPipedriveClient()
		if err != nil {
			return err
		}

		for _, org := range gameofthrones.Organizations {
			fmt.Printf("Creating organization: %s\n", org)
			resp, err := pc.CreateOrganization(pdRequestBody{Name: org})
			if err != nil {
				return err
			}
			fmt.Printf("Status: %d\n", resp.StatusCode)
		}
		return nil
	},
}

var pdCreatePeopleCmd = &cobra.Command{
	Use:   "create-people",
	Short: "Create Pipedrive persons from GoT characters",
	RunE: func(cmd *cobra.Command, args []string) error {
		pc, err := newPipedriveClient()
		if err != nil {
			return err
		}

		chars := gameofthrones.Characters()
		for _, char := range chars {
			fmtutil.MustPrintJSON(char)
			resp, err := pc.CreatePerson(pdRequestBody{Name: char.Character.DisplayName})
			if err != nil {
				return err
			}
			fmt.Printf("Status: %d\n", resp.StatusCode)
		}
		return nil
	},
}

var pdGetPersonsCmd = &cobra.Command{
	Use:   "get-persons",
	Short: "List all Pipedrive persons",
	RunE: func(cmd *cobra.Command, args []string) error {
		pc, err := newPipedriveClient()
		if err != nil {
			return err
		}

		resp, err := pc.GetPersons()
		if err != nil {
			return err
		}
		return httputilmore.PrintResponse(resp, true)
	},
}

var pdGetOrgFieldsCmd = &cobra.Command{
	Use:   "get-org-fields",
	Short: "List Pipedrive organization fields",
	RunE: func(cmd *cobra.Command, args []string) error {
		pc, err := newPipedriveClient()
		if err != nil {
			return err
		}

		resp, err := pc.GetOrganizationFields()
		if err != nil {
			return err
		}
		return httputilmore.PrintResponse(resp, true)
	},
}

func init() {
	pipedriveCmd.PersistentFlags().StringVar(&pipedriveBaseURL, "base-url", "", "Pipedrive API base URL (e.g., https://yourcompany.pipedrive.com/v1)")

	pipedriveCmd.AddCommand(pdCreateOrgsCmd)
	pipedriveCmd.AddCommand(pdCreatePeopleCmd)
	pipedriveCmd.AddCommand(pdGetPersonsCmd)
	pipedriveCmd.AddCommand(pdGetOrgFieldsCmd)
}

// Pipedrive types and helpers

type pdRequestBody struct {
	Name string `json:"name,omitempty"`
}

type pipedriveClient struct {
	APIKey     string
	BaseURL    string
	ClientMore httputilmore.ClientMore
}

func newPipedriveClient() (*pipedriveClient, error) {
	_, err := config.LoadDotEnv([]string{os.Getenv("ENV_PATH"), "./.env"}, 1)
	if err != nil {
		return nil, err
	}

	baseURL := pipedriveBaseURL
	if baseURL == "" {
		baseURL = os.Getenv("PIPEDRIVE_BASE_URL")
	}
	if baseURL == "" {
		return nil, fmt.Errorf("Pipedrive base URL required: use --base-url flag or PIPEDRIVE_BASE_URL env var")
	}

	apiKey := os.Getenv("PIPEDRIVE_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("PIPEDRIVE_API_KEY environment variable required")
	}

	client := &http.Client{}
	return &pipedriveClient{
		APIKey:     apiKey,
		BaseURL:    baseURL,
		ClientMore: httputilmore.ClientMore{Client: client},
	}, nil
}

func (pc *pipedriveClient) buildURL(path string) (string, error) {
	urlString := urlutil.JoinAbsolute(pc.BaseURL, path)
	u, err := url.ParseRequestURI(urlString)
	if err != nil {
		return "", err
	}
	v := url.Values{}
	v.Add("api_token", pc.APIKey)
	uz, err := urlutil.URLStringAddQuery(u.String(), v, false)
	if err != nil {
		return "", err
	}
	return uz.String(), nil
}

func (pc *pipedriveClient) GetOrganizationFields() (*http.Response, error) {
	apiURL, err := pc.buildURL("/organizationFields")
	if err != nil {
		return nil, err
	}
	return pc.ClientMore.Client.Get(apiURL) //nolint:gosec
}

func (pc *pipedriveClient) GetPersons() (*http.Response, error) {
	apiURL, err := pc.buildURL("/persons")
	if err != nil {
		return nil, err
	}
	return pc.ClientMore.Client.Get(apiURL) //nolint:gosec
}

func (pc *pipedriveClient) CreateOrganization(reqBody pdRequestBody) (*http.Response, error) {
	apiURL, err := pc.buildURL("/organizations")
	if err != nil {
		return nil, err
	}
	return pc.ClientMore.PostToJSON(apiURL, reqBody)
}

func (pc *pipedriveClient) CreatePerson(reqBody pdRequestBody) (*http.Response, error) {
	apiURL, err := pc.buildURL("/persons")
	if err != nil {
		return nil, err
	}
	return pc.ClientMore.PostToJSON(apiURL, reqBody)
}
