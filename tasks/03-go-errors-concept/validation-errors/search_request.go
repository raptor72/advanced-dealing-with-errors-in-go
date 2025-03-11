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
	// validationErrors   error
)


type ValidationErrors []error
// type ValidationErrors struct {
	// errs []error
// }

func (verr ValidationErrors) Error() string {
	if len(verr) > 0 {
    	var validationErrors = errors.New("validation errors:\n")
        
		for _, e := range verr {
			validationErrors = fmt.Errorf("%v\t%w\n", validationErrors, e)
		}
        return validationErrors.Error()
	}

	// var s string
	// if len(verr) != 0 {
    // 	s += "validation errors:\n"
	//     for _, e := range verr{
	// 	    s += "\t" + e.Error() + "\n"
    // 	}
    // }
	// return s
	return ""
}

// func (verr ValidationErrors) Unwrap() error {
// 	if len(verr) != 0 {
//         for _, e := range verr {
// 			return e
// 		}
// 	}
// 	return nil
// }

func (verr *ValidationErrors) Is(target error) bool {
    for _, e := range *verr {
		return errors.Is(e, target)
	}
    return false
}

type SearchRequest struct {
	Exp      string
	Page     int
	PageSize int
}

func (r SearchRequest) Validate() error {
    var vers ValidationErrors
	_, err := regexp.Compile(r.Exp)
	if err != nil {
		var errIsNotRegexp = errors.New("exp is not regexp: ")
		errIsNotRegexp = fmt.Errorf("%v%w", errIsNotRegexp, err)
        vers = append(vers, errIsNotRegexp)
	}
	if r.Page <= 0 {
		var errInvalidPage = fmt.Errorf("invalid page: %v", r.Page)
		vers = append(vers, errInvalidPage)
	}
	if r.PageSize > maxPageSize {
		var errInvalidPageSize = fmt.Errorf("invalid page size: %v > %v", r.PageSize, maxPageSize)
		vers = append(vers, errInvalidPageSize)
	}
	if r.PageSize <= 0 {
		var errInvalidPageSize = fmt.Errorf("invalid page size: %v < 0", r.PageSize)
		vers = append(vers, errInvalidPageSize)
	}
	if len(vers)>0 {
	    return vers
	}
	return &vers
}
