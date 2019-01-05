package salesloftapi

import (
	"github.com/go-chi/render"
	"github.com/slpeople/errors"
)

func ErrListPeople(err error) render.Renderer {
	return &errors.ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Error listing people from SalesLoft API.",
		ErrorText:      err.Error(),
	}
}
