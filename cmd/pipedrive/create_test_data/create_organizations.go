package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/grokify/gameofthrones"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/mogo/net/http/httputilmore"
	"github.com/grokify/mogo/net/urlutil"
	flags "github.com/jessevdk/go-flags"
)

var BaseURL = "https://companydomain.pipedrive.com/v1"

type RequestBody struct {
	Name string `json:"name,omitempty"`
}

type PipedriveClient struct {
	APIKey     string
	ClientMore httputilmore.ClientMore
}

func NewPipedriveClient(apiKey string) PipedriveClient {
	client := &http.Client{}
	cm := httputilmore.ClientMore{Client: client}
	return PipedriveClient{
		APIKey:     apiKey,
		ClientMore: cm}
}

func (pc *PipedriveClient) BuildURL(path string) (string, error) {
	urlString := urlutil.JoinAbsolute(BaseURL, path)
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

func (pc *PipedriveClient) GetOrganizationFields() (*http.Response, error) {
	apiURL, err := pc.BuildURL("/organizationFields")
	if err != nil {
		return nil, err
	}
	return pc.ClientMore.Client.Get(apiURL)
}

func (pc *PipedriveClient) GetPersons() (*http.Response, error) {
	apiURL, err := pc.BuildURL("/persons")
	if err != nil {
		return nil, err
	}
	return pc.ClientMore.Client.Get(apiURL)
}

func (pc *PipedriveClient) CreateOrganization(reqBody RequestBody) (*http.Response, error) {
	apiURL, err := pc.BuildURL("/organizations")
	if err != nil {
		return &http.Response{}, err
	}
	return pc.ClientMore.PostToJSON(apiURL, reqBody)
}

func (pc *PipedriveClient) CreatePerson(reqBody RequestBody) (*http.Response, error) {
	apiURL, err := pc.BuildURL("/persons")
	if err != nil {
		return &http.Response{}, err
	}
	return pc.ClientMore.PostToJSON(apiURL, reqBody)
}

func (pc *PipedriveClient) CreateOrIgnorePerson(reqBody RequestBody) (*http.Response, error) {
	apiURL, err := pc.BuildURL("/persons")
	if err != nil {
		return &http.Response{}, err
	}
	return pc.ClientMore.PostToJSON(apiURL, reqBody)
}

type Options struct {
	Command string `short:"c" long:"create" description:"Command [get_persons,get_org_fields,create_people,create_orgs]"`
}

func main() {
	_, err := config.LoadDotEnv([]string{}, 1)
	logutil.FatalErr(err)

	opts := Options{}
	_, err = flags.Parse(&opts)
	logutil.FatalErr(err)

	pc := NewPipedriveClient(os.Getenv("PIPEDRIVE_API_KEY"))

	switch opts.Command {
	case "get_persons":
		resp, err := pc.GetPersons()
		logutil.FatalErr(err)
		logutil.FatalErr(httputilmore.PrintResponse(resp, true))
	case "get_org_fields":
		resp, err := pc.GetOrganizationFields()
		logutil.FatalErr(err)
		logutil.FatalErr(httputilmore.PrintResponse(resp, true))
	case "create_people":
		orgs := gameofthrones.Organizations
		for _, org := range orgs {
			fmt.Println(org)
			resp, err := pc.CreateOrganization(RequestBody{Name: org})
			logutil.FatalErr(err)
			fmt.Println(resp.StatusCode)
			// break
		}
	case "create_orgs":
		chars := gameofthrones.Characters()
		for _, char := range chars {
			fmtutil.MustPrintJSON(char)
			resp, err := pc.CreatePerson(RequestBody{Name: char.Character.DisplayName})
			logutil.FatalErr(err)
			fmt.Println(resp.StatusCode)
		}
	default:
		log.Fatal("command must be one of [get_persons,get_org_fields,create_people,create_orgs]")
	}
}
