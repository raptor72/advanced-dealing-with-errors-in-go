package requests

import (
	// Доступные пакеты, _ для сохранения импортов.
	"errors"
	"fmt"
	"regexp"
	_ "strings"
)

const maxPageSize = 100

var (
	errIsNotRegexp      = errors.New("exp is not regexp")
	errInvalidPage      = errors.New("invalid page")
	errInvalidPageSize  = errors.New("invalid page size")
)

type ValidationErrors []error


func (verr ValidationErrors) Error() string {
	if len(verr) > 0 {
    	var validationErrors = errors.New("validation errors:\n") 
		for _, e := range verr {
			validationErrors = fmt.Errorf("%v\t%w\n", validationErrors, e)
		}
        return validationErrors.Error()
	}
	return ""
	// var s string
	// if len(verr) != 0 {
    // 	s += "validation errors:\n"
	//     for _, e := range verr{
	// 	    s += "\t" + e.Error() + "\n"
    // 	}
    // }
	// return s

}


func (verr *ValidationErrors) Is(target error) bool {
    for _, e := range *verr {
		if errors.Is(e, target) {
			return true
		}
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
        vers = append(vers, fmt.Errorf("%w: %v", errIsNotRegexp, err))
	}
	if r.Page <= 0 {
		vers = append(vers, fmt.Errorf("%w: %v", errInvalidPage, r.Page))
	}
	if r.PageSize > maxPageSize {
		vers = append(vers, fmt.Errorf("%w: %v > %v", errInvalidPageSize, r.PageSize, maxPageSize) )
	}
	if r.PageSize <= 0 {
		vers = append(vers, fmt.Errorf("%w: %v <= 0", errInvalidPageSize, r.PageSize) )
	}
	if len(vers)>0 {
	    return &vers
	}
	return nil
}
