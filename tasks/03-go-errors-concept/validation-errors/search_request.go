package requests

import (
	// Доступные пакеты, _ для сохранения импортов.
	_ "errors"
	_ "fmt"
	"regexp"
	_ "strings"
)

const maxPageSize = 100

// Реализуй нас.
var (
	// errIsNotRegexp     error
	errIsNotRegexp     ErrIsNotRegexp
	errInvalidPage     error
	errInvalidPageSize error
)

type ErrIsNotRegexp struct {
	Msg    string
	Origin error
}
func (e ErrIsNotRegexp) Error() string {
	return e.Msg + e.Origin.Error()
}
func (e ErrIsNotRegexp) Unwrap() error {
    return e.Origin
}


// Реализуй мои методы.
type ValidationErrors []error

type SearchRequest struct {
	Exp      string
	Page     int
	PageSize int
}

func (r SearchRequest) Validate() error {
	// Реализуй меня.
	_, err := regexp.Compile(r.Exp)
    if err != nil {
		return ErrIsNotRegexp{Msg: "exp is not regexp: ", Origin: err}
	}

	return nil
}
