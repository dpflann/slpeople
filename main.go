package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

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
	ErrResponse struct {
		Err            error `json:"-"` // low-level runtime error
		HTTPStatusCode int   `json:"-"` // http response status code

		StatusText string `json:"status"`          // user-level status message
		AppCode    int64  `json:"code,omitempty"`  // application-specific error code
		ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
	}
)

var (
	apikey = flag.String("apikey", "", "SalesLoft API Key for communications with SalesLoft API (https://developers.salesloft.com/api.html)")
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
		//r.Get("/frequencies", EmailCharacterFrequencies)
		//r.Get("/duplicates", DuplicateEmails)
	})
	http.ListenAndServe(":3000", r)
}

/*** Level 1: List People ***/
func (p *PeopleListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ListPeople(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", SalesLoftApiURL, nil)
	req.Header.Add("Authorization", "Bearer "+*apikey)
	payload := url.Values{}
	payload.Add("per_page", "100")
	resp, err := client.Do(req)
	if err != nil {
		render.Render(w, r, ErrListPeople(err))
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		render.Render(w, r, ErrListPeople(err))
		return
	}
	salesLoftPeople := &SalesLoftApiPeopleResponse{}
	if err := json.Unmarshal(body, salesLoftPeople); err != nil {
		render.Render(w, r, ErrListPeople(err))
		return
	}
	if err := render.Render(w, r, NewPeopleListResponse(salesLoftPeople.Data)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
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
func CharacterFrequencyCount(str string) map[string]int {
	frequencies := map[string]int{}
	for _, c := range str {
		cStr := string(c)
		if _, ok := frequencies[cStr]; ok {
			frequencies[cStr] += 1
		} else {
			frequencies[cStr] = 1
		}
	}
	// TODO add sorting by value
	return frequencies
}

func CharacterFrequencyCountOfStrings(strs []string) map[string]int {
	// Naive handling
	result := map[string]int{}
	for _, s := range strs {
		res := CharacterFrequencyCount(s)
		for c, v := range res {
			if _, ok := result[c]; ok {
				result[c] += v
			} else {
				result[c] = v
			}
		}
	}
	return result
}
