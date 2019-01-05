package characters

import (
	"github.com/slpeople/errors"

	"github.com/go-chi/render"
)

func ErrCharacterFrequency(err error) render.Renderer {
	return &errors.ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Error while calculating character frequency",
		ErrorText:      err.Error(),
	}
}
