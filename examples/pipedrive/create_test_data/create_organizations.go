package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/grokify/gameofthrones"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/net/httputilmore"
	"github.com/grokify/gotilla/net/urlutil"
)

var BaseURL = "https://companydomain.pipedrive.com/v1"

type RequestBody struct {
	Name string `json:"name,omitempty"`
}

type PipedriveClient struct {
	ApiKey     string
	ClientMore httputilmore.ClientMore
}

func NewPipedriveClient(apiKey string) PipedriveClient {
	client := &http.Client{}
	cm := httputilmore.ClientMore{Client: client}
	return PipedriveClient{
		ApiKey:     apiKey,
		ClientMore: cm}
}

func (pc *PipedriveClient) BuildURL(path string) (string, error) {
	urlString := urlutil.JoinAbsolute(BaseURL, path)
	u, err := url.ParseRequestURI(urlString)
	if err != nil {
		return "", err
	}
	v := url.Values{}
	v.Add("api_token", pc.ApiKey)
	uz := urlutil.BuildURL(u.String(), v)
	return uz, nil
}

func (pc *PipedriveClient) GetOrganizationFields() (*http.Response, error) {
	apiUrl, err := pc.BuildURL("/organizationFields")
	if err != nil {
		return &http.Response{}, err
	}
	return pc.ClientMore.Client.Get(apiUrl)
}

func (pc *PipedriveClient) GetPersons() (*http.Response, error) {
	apiUrl, err := pc.BuildURL("/persons")
	if err != nil {
		return &http.Response{}, err
	}
	return pc.ClientMore.Client.Get(apiUrl)
}

func (pc *PipedriveClient) CreateOrganization(reqBody RequestBody) (*http.Response, error) {
	apiUrl, err := pc.BuildURL("/organizations")
	if err != nil {
		return &http.Response{}, err
	}
	return pc.ClientMore.PostToJSON(apiUrl, reqBody)
}

func (pc *PipedriveClient) CreatePerson(reqBody RequestBody) (*http.Response, error) {
	apiUrl, err := pc.BuildURL("/persons")
	if err != nil {
		return &http.Response{}, err
	}
	return pc.ClientMore.PostToJSON(apiUrl, reqBody)
}

func (pc *PipedriveClient) CreateOrIgnorePerson(reqBody RequestBody) (*http.Response, error) {
	apiUrl, err := pc.BuildURL("/persons")
	if err != nil {
		return &http.Response{}, err
	}
	return pc.ClientMore.PostToJSON(apiUrl, reqBody)
}

func main() {
	config.LoadDotEnv()

	pc := NewPipedriveClient(os.Getenv("PIPEDRIVE_API_KEY"))

	if 1 == 0 {
		resp, err := pc.GetPersons()
		if err != nil {
			panic(err)
		}
		httputilmore.PrintResponse(resp, true)
	}

	if 1 == 0 {
		resp, err := pc.GetOrganizationFields()
		if err != nil {
			panic(err)
		}
		httputilmore.PrintResponse(resp, true)
	}

	if 1 == 0 {
		orgs := gameofthrones.Organizations
		for _, org := range orgs {
			fmt.Println(org)
			resp, err := pc.CreateOrganization(RequestBody{Name: org})
			if err != nil {
				panic(err)
			}
			fmt.Println(resp.StatusCode)
			break
		}
	}

	if 1 == 1 {
		chars, err := gameofthrones.ReadCharacters()
		if err != nil {
			panic(err)
		}
		for _, char := range chars {
			fmtutil.PrintJSON(char)
			resp, err := pc.CreatePerson(RequestBody{Name: char.Character.DisplayName})
			if err != nil {
				panic(err)
			}
			fmt.Println(resp.StatusCode)
		}
	}
}
