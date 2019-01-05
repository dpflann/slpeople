package duplicates

import (
	"net/http"

	"github.com/go-chi/render"
	errors "github.com/slpeople/errors"
	slapi "github.com/slpeople/salesloftapi"
)

/*** Level 3: Duplicate Email Addresses ***/
// References:
//  >> https://stackoverflow.com/questions/577463/finding-how-similar-two-strings-are
//  >> https://github.com/agnivade/levenshtein/blob/master/levenshtein.go
//  >> https://gist.github.com/andrei-m/982927#gistcomment-1931258
type (
	PossibleDuplicatesResponse struct {
		*PossibleDuplicates `json:"possibleDuplicates"`
	}
)

func PossibleDuplicateEmailsHandler(w http.ResponseWriter, r *http.Request) {
	people, err := slapi.ListPeople()
	if err != nil {
		render.Render(w, r, ErrDuplicates(err))
		return
	}
	emailAddresses := make([]string, len(*people))
	for i := range *people {
		emailAddresses[i] = (*people)[i].EmailAddress
	}
	// TODO: Make this configurable at startup.
	settings := thresholdSettings{
		distanceThreshold: 1,
		lengthThreshold:   1,
	}
	duplicateEmailAddresses := FindPossibleDuplicates(emailAddresses, settings)
	if err := render.Render(w, r, NewPossibleDuplicatesResponse(&duplicateEmailAddresses)); err != nil {
		render.Render(w, r, errors.ErrRender(err))
		return
	}
}

func NewPossibleDuplicatesResponse(pdupes *PossibleDuplicates) *PossibleDuplicatesResponse {
	return &PossibleDuplicatesResponse{PossibleDuplicates: pdupes}
}

func (pd *PossibleDuplicatesResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
