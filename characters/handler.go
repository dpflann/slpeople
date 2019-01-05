package characters

import (
	"net/http"

	"github.com/go-chi/render"
	errors "github.com/slpeople/errors"
	slapi "github.com/slpeople/salesloftapi"
)

var (
	blackList = map[string]bool{
		".": true,
		"@": true,
	}
)

/*** Level 2: Unique Character Frequencies ***/
func EmailCharacterFrequenciesHandler(w http.ResponseWriter, r *http.Request) {
	people, err := slapi.ListPeople()
	if err != nil {
		render.Render(w, r, ErrCharacterFrequency(err))
		return
	}
	emailAddresses := make([]string, len(*people))
	for i := range *people {
		emailAddresses[i] = (*people)[i].EmailAddress
	}
	charFrequencies := CharacterFrequencyCountOfStrings(emailAddresses, blackList)
	if err := render.Render(w, r, NewSortedCharacterFrequenciesResponse(&charFrequencies)); err != nil {
		render.Render(w, r, errors.ErrRender(err))
		return
	}
}

func NewSortedCharacterFrequenciesResponse(charFrequencies *CharacterFrequencies) *SortedCharacterFrequenciesResponse {
	return &SortedCharacterFrequenciesResponse{SortedCharFreqs: charFrequencies.Sorted()}
}

func (c *SortedCharacterFrequenciesResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
