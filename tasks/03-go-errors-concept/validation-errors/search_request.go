package requests

import (
	// Доступные пакеты, _ для сохранения импортов.
	"errors"
	"fmt"

	// _ "fmt"
	"regexp"
	_ "strings"
)

const maxPageSize = 100

var (
	errIsNotRegexp     error
	errInvalidPage     error
	errInvalidPageSize error
	validationErrors   error
)


type ValidationErrors []error
// type ValidationErrors struct {
	// errs []error
// }

func (verr ValidationErrors) Error() string {
    var s string
	if len(verr) != 0 {
    	s += "validation errors:\n"
	    for _, e := range verr{
		    s += "\t" + e.Error() + "\n"
    	}
    }
	return s
}

// func (verr ValidationErrors) Unwrap() error {
// 	if len(verr) != 0 {
//         for _, e := range verr {
// 			return e
// 		}
// 	}
// 	return nil
// }

func (verr ValidationErrors) Is(target error) bool {
	return errors.Is(verr, target)
}

type SearchRequest struct {
	Exp      string
	Page     int
	PageSize int
}

func (r SearchRequest) Validate() ValidationErrors {
    var vers ValidationErrors
	haserrs := false
	var validationErrors = errors.New("validation errors:")
	_, err := regexp.Compile(r.Exp)
	if err != nil {
		// var errIsNotRegexp = errors.New("\texp is not regexp: ")
		var errIsNotRegexp = errors.New("exp is not regexp: ")
		errIsNotRegexp = fmt.Errorf("%v%w", errIsNotRegexp, err)
		validationErrors = fmt.Errorf("%v\n%w", validationErrors, errIsNotRegexp)
        vers = append(vers, errIsNotRegexp)
		haserrs = true
	}
	if r.Page <= 0 {
		var errInvalidPage = fmt.Errorf("invalid page: %v", r.Page)
		validationErrors = fmt.Errorf("%v\n%w", validationErrors, errInvalidPage)
		vers = append(vers, errInvalidPage)
		haserrs = true
	}
	if r.PageSize > maxPageSize {
		var errInvalidPageSize = fmt.Errorf("invalid page size: %v > %v", r.PageSize, maxPageSize)
		validationErrors = fmt.Errorf("%v\n%w", validationErrors, errInvalidPageSize)
		vers = append(vers, errInvalidPageSize)
		haserrs = true
	}
	if r.PageSize <= 0 {
		var errInvalidPageSize = fmt.Errorf("invalid page size: %v < 0", r.PageSize)
		validationErrors = fmt.Errorf("%v\n%w", validationErrors, errInvalidPageSize)
		vers = append(vers, errInvalidPageSize)
		haserrs = true
	}
	if haserrs {
	    return vers
		// return validationErrors
	}
	return nil
}
