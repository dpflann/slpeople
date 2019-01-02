package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

const (
	SalesLoftApiURL = "https://api.salesloft.com/v2/people.json"
)

type (
	SimplifiedPersonView struct {
		ID                    int    `json:"id"`
		CreatedAt             string `json:"created_at"`
		UpdatdedAt            string `json:"updated_at"`
		FirstName             string `json:"first_name"`
		LastName              string `json:"last_name"`
		DisplayName           string `json:"display_name"`
		EmailAddress          string `json:"email_address"`
		SecondaryEmailAddress string `json"secondary_email_address"`
		PersonalEmailAddress  string `json:"personal_email_address"`
	}
	People             []SimplifiedPersonView
	PeopleListResponse struct {
		*People
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
	SalesLoftApiPeopleResponse struct {
		Metadata SalesLoftApiMetadata `json:"metadata"`
		Data     *People              `json:"data"`
	}
	CharacterFrequencies               map[string]int
	SortedCharacterFrequenciesResponse struct {
		*SortedCharFreqs `json:"frequencies"`
	}
	KeyVal struct {
		Key   string `json:"key"`
		Value int    `json:"value"`
	}
	SortedCharFreqs []KeyVal
	ErrResponse     struct {
		Err            error `json:"-"` // low-level runtime error
		HTTPStatusCode int   `json:"-"` // http response status code

		StatusText string `json:"status"`          // user-level status message
		AppCode    int64  `json:"code,omitempty"`  // application-specific error code
		ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
	}
)

var (
	apikey    = flag.String("apikey", "", "SalesLoft API Key for communications with SalesLoft API (https://developers.salesloft.com/api.html)")
	blackList = map[string]bool{
		".": true,
		"@": true,
	}
)

func main() {
	flag.Parse()
	if *apikey == "" {
		fmt.Fprintf(os.Stderr, "An API Key is required. Please obtain a SalesLoft API key from your account or contact SalesLoft Support (support@salesloft.com) for assistance.")
		os.Exit(1)
	}
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// RESTy routes for "articles" resource
	r.Route("/people", func(r chi.Router) {
		r.Get("/", ListPeople)
		r.Get("/frequencies", EmailCharacterFrequencies)
		//r.Get("/duplicates", DuplicateEmails)
	})
	http.ListenAndServe(":3000", r)
}

/*** Level 1: List People ***/
func (p *PeopleListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ListPeople(w http.ResponseWriter, r *http.Request) {
	people, err := listSalesLoftPeople()
	if err != nil {
		render.Render(w, r, ErrListPeople(err))
		return
	}
	if err := render.Render(w, r, NewPeopleListResponse(people)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func listSalesLoftPeople() (*People, error) {
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
		resp, err = getPeople(*perPage, *nextPage)
		people = append(people, []SimplifiedPersonView(*resp.Data)...)
		perPage = resp.Metadata.Paging.PerPage
		nextPage = resp.Metadata.Paging.NextPage
	}
	if err != nil {
		return nil, err
	}
	return &people, nil
}

func getPeople(perPage, page int) (*SalesLoftApiPeopleResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", SalesLoftApiURL, nil)
	req.Header.Add("Authorization", "Bearer "+*apikey)
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

func NewPeopleListResponse(people *People) *PeopleListResponse {
	return &PeopleListResponse{People: people}
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

func ErrListPeople(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Error listing people from SalesLoft API.",
		ErrorText:      err.Error(),
	}
}

/*** Level 2: Unique Character Frequencies ***/
func CharacterFrequencyCount(str string, blackList map[string]bool) CharacterFrequencies {
	frequencies := CharacterFrequencies{}
	for _, c := range str {
		cStr := string(c)
		if _, ok := blackList[cStr]; ok {
			continue
		}
		if _, ok := frequencies[cStr]; ok {
			frequencies[cStr] += 1
		} else {
			frequencies[cStr] = 1
		}
	}
	return frequencies
}

func CharacterFrequencyCountOfStrings(strs []string, blackList map[string]bool) CharacterFrequencies {
	// Naive handling
	frequencies := CharacterFrequencies{}
	for _, s := range strs {
		res := CharacterFrequencyCount(s, blackList)
		for c, v := range res {
			if _, ok := frequencies[c]; ok {
				frequencies[c] += v
			} else {
				frequencies[c] = v
			}
		}
	}
	return frequencies
}

func NewSortedCharacterFrequenciesResponse(charFrequencies *CharacterFrequencies) *SortedCharacterFrequenciesResponse {
	return &SortedCharacterFrequenciesResponse{SortedCharFreqs: charFrequencies.Sorted()}
}

func (c *SortedCharacterFrequenciesResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (c *CharacterFrequencies) Sorted() *SortedCharFreqs {
	var cfs SortedCharFreqs
	for char, count := range *c {
		cfs = append(cfs, KeyVal{char, count})
	}

	sort.Slice(cfs, func(i, j int) bool {
		return cfs[i].Value > cfs[j].Value
	})
	return &cfs
}

func EmailCharacterFrequencies(w http.ResponseWriter, r *http.Request) {
	people, err := listSalesLoftPeople()
	if err != nil {
		render.Render(w, r, ErrListPeople(err))
		return
	}
	emailAddresses := make([]string, len(*people))
	for i := range *people {
		emailAddresses[i] = (*people)[i].EmailAddress
	}
	charFrequencies := CharacterFrequencyCountOfStrings(emailAddresses, blackList)
	if err := render.Render(w, r, NewSortedCharacterFrequenciesResponse(&charFrequencies)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}
