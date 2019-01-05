package salesloftapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type (
	SalesLoftClient struct {
		apiKey string
		apiUrl string
	}
	SimplifiedPersonView struct {
		ID                    int    `json:"id"`
		CreatedAt             string `json:"created_at"`
		UpdatdedAt            string `json:"updated_at"`
		FirstName             string `json:"first_name"`
		LastName              string `json:"last_name"`
		DisplayName           string `json:"display_name"`
		EmailAddress          string `json:"email_address"`
		SecondaryEmailAddress string `json:"secondary_email_address"`
		PersonalEmailAddress  string `json:"personal_email_address"`
		Title                 string `json:"title"`
	}
	People             []SimplifiedPersonView
	PeopleListResponse struct {
		*People `json:"people"`
	}
	SalesLoftApiPagingMetadata struct {
		PerPage     *int `json:"per_page"`
		CurrentPage *int `json:"current_page"`
		NextPage    *int `json:"next_page"`
		PrevPage    *int `json:prev_page"`
		TotalPages  *int `json:"total_pages,omitempty"`
		TotalCount  *int `json:"total_count,omitempty"`
	}
	SalesLoftApiMetadata struct {
		Paging SalesLoftApiPagingMetadata `json:"paging"`
	}
)

var (
	slClient *SalesLoftClient
)

func InitializeClient(apiKey, apiUrl string) *SalesLoftClient {
	slClient = &SalesLoftClient{
		apiKey: apiKey,
		apiUrl: apiUrl,
	}
	return slClient
}

func ListPeople() (*People, error) {
	people := People{}
	pp := 100
	np := 0
	var perPage *int
	var nextPage *int
	perPage = &pp
	nextPage = &np
	resp := &SalesLoftApiPeopleResponse{}
	var err error
	for err == nil && nextPage != nil {
		resp, err = slClient.getPeople(*perPage, *nextPage)
		if err != nil || resp.Data == nil {
			break
		}
		people = append(people, []SimplifiedPersonView(*resp.Data)...)
		perPage = resp.Metadata.Paging.PerPage
		nextPage = resp.Metadata.Paging.NextPage
	}
	if err != nil {
		return nil, err
	}
	return &people, nil
}

func (slClient *SalesLoftClient) getPeople(perPage, page int) (*SalesLoftApiPeopleResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", slClient.apiUrl, nil)
	req.Header.Add("Authorization", "Bearer "+slClient.apiKey)
	q := req.URL.Query()
	q.Add("per_page", strconv.Itoa(perPage))
	q.Add("page", strconv.Itoa(page))
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	salesLoftPeople := &SalesLoftApiPeopleResponse{}
	if err := json.Unmarshal(body, salesLoftPeople); err != nil {
		return nil, err
	}
	return salesLoftPeople, nil
}
