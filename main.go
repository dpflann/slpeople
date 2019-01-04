package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

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

	PossibleDuplicates         [][]string
	PossibleDuplicatesResponse struct {
		*PossibleDuplicates
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
	apikey    = flag.String("apikey", "", "SalesLoft API Key for communications with SalesLoft API (https://developers.salesloft.com/api.html)")
	port      = flag.String("port", "3000", "The port for the service. The default value is 3000.")
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
	} else {
		log.Printf("Using API key: %s\n", *apikey)
	}
	if *port == "" {
		fmt.Fprintf(os.Stderr, "The port was set to empty string. :(")
		os.Exit(2)
	} else {
		log.Printf("Using port: %s\n", *port)
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
		r.Get("/char_frequencies", EmailCharacterFrequencies)
		r.Get("/duplicates", PossibleDuplicateEmails)
	})
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	FileServer(r, "/static", http.Dir(filesDir))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.ListenAndServe(":"+*port, r)
}

/*** Level 1: List People ***/
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

func (p *PeopleListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
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

func NewSortedCharacterFrequenciesResponse(charFrequencies *CharacterFrequencies) *SortedCharacterFrequenciesResponse {
	return &SortedCharacterFrequenciesResponse{SortedCharFreqs: charFrequencies.Sorted()}
}

func (c *SortedCharacterFrequenciesResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

/*** Level 3: Duplicate Email Addresses ***/
//https://stackoverflow.com/questions/577463/finding-how-similar-two-strings-are

func PossibleDuplicateEmails(w http.ResponseWriter, r *http.Request) {
	people, err := listSalesLoftPeople()
	if err != nil {
		render.Render(w, r, ErrListPeople(err))
		return
	}
	emailAddresses := make([]string, len(*people))
	for i := range *people {
		emailAddresses[i] = (*people)[i].EmailAddress
	}
	duplicateEmailAddresses := FindPossibleDuplicates(emailAddresses)
	if err := render.Render(w, r, NewPossibleDuplicatesResponse(&duplicateEmailAddresses)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func NewPossibleDuplicatesResponse(pdupes *PossibleDuplicates) *PossibleDuplicatesResponse {
	return &PossibleDuplicatesResponse{PossibleDuplicates: pdupes}
}

func (pd *PossibleDuplicatesResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func FindPossibleDuplicates(strs []string) PossibleDuplicates {
	duplicates := PossibleDuplicates{}
	// 1. Compare the lengths
	// 2. Compare the characters
	for i := 0; i < len(strs); i++ {
		dupes := []string{strs[i]}
		for j := i + 1; j < len(strs); j++ {
			if !compareLengths(strs[i], strs[j], 1) {
				continue
			}
			if !compareChars(strs[i], strs[j], 0) {
				continue
			}
			dupes = append(dupes, strs[j])
		}
		if len(dupes) > 1 {
			duplicates = append(duplicates, dupes)
		}
	}
	return duplicates
}

func compareLengths(str1, str2 string, threshold int) bool {
	if len(str1) == len(str2) {
		return true
	}
	if len(str1) > len(str2) {
		return (len(str1) - len(str2) - threshold) == 0
	}
	return (len(str2) - len(str1) - threshold) == 0
}

func compareChars(str1, str2 string, threshold int) bool {
	chars1 := CharacterFrequencyCount(str1, nil)
	onlyChars1 := bytes.NewBuffer(nil)
	chars2 := CharacterFrequencyCount(str2, nil)
	onlyChars2 := bytes.NewBuffer(nil)
	for c1 := range chars1 {
		if _, ok := chars2[c1]; !ok {
			onlyChars1.WriteString(c1)
		}
	}
	for c2 := range chars2 {
		if _, ok := chars1[c2]; !ok {
			onlyChars2.WriteString(c2)
		}
	}
	if onlyChars1.Len() == 0 && onlyChars2.Len() == 0 {
		return true
	}
	return compareLengths(onlyChars1.String(), onlyChars2.String(), 0)
}

/*** Errors ***/
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

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

/*** Web App ***/

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
