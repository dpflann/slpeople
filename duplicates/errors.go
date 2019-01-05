package duplicates

import (
	"github.com/go-chi/render"
	"github.com/slpeople/errors"
)

func ErrDuplicates(err error) render.Renderer {
	return &errors.ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Error while finding possible duplicate email addresses",
		ErrorText:      err.Error(),
	}
}
