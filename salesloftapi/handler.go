package salesloftapi

import (
	"net/http"

	"github.com/slpeople/errors"

	"github.com/go-chi/render"
)

type (
	SalesLoftApiPeopleResponse struct {
		Metadata SalesLoftApiMetadata `json:"metadata"`
		Data     *People              `json:"data"`
	}
)

/*** Level 1: List People ***/
func ListPeopleHandler(w http.ResponseWriter, r *http.Request) {
	people, err := ListPeople()
	if err != nil {
		render.Render(w, r, ErrListPeople(err))
		return
	}
	if err := render.Render(w, r, NewPeopleListResponse(people)); err != nil {
		render.Render(w, r, errors.ErrRender(err))
		return
	}
}

func NewPeopleListResponse(people *People) *PeopleListResponse {
	return &PeopleListResponse{People: people}
}

func (p *PeopleListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
