package requests

import (
	// Доступные пакеты, _ для сохранения импортов.
	"errors"
	"fmt"

	// _ "fmt"
	"regexp"
	"strings"
)

const maxPageSize = 100

// Реализуй нас.
var (
	// errIsNotRegexp     error
	errIsNotRegexp     ErrIsNotRegexp
	errInvalidPage     ErrInvalidPage
	errInvalidPageSize ErrInvalidPageSize
)

type ErrIsNotRegexp struct {
	Msg    string
	Origin error
}

func (e ErrIsNotRegexp) Error() string {
	if e.Origin != nil {
//        return e.Msg + e.Origin.Error()
		return fmt.Errorf("%v%w", e.Msg, e.Origin).Error()
	}
	return e.Msg
}

func (e ErrIsNotRegexp) Unwrap() error {
	if e.Origin != nil {
        return e.Origin
	}
	return nil
}

func (e *ErrIsNotRegexp) Is(target error) bool {
    return strings.Contains(target.Error(), e.Msg)
}

type ErrInvalidPage struct {
	Msg    string
	Page   int
	Origin error
}

func (e ErrInvalidPage) Error() string {
	return e.Msg + fmt.Sprintf("%d", e.Page)
}
func (e ErrInvalidPage) Unwrap() error {
	return e.Origin
}

type ErrInvalidPageSize struct {
	Msg      string
    Comp     string
    Val      string
	PageSize int
	Origin   error
}

func (e ErrInvalidPageSize) Error() string {
	return e.Msg + fmt.Sprintf("%v %v %v", e.PageSize, e.Comp, e.Val)
	// return e.Msg
}
func (e ErrInvalidPageSize) Unwrap() error {
	return e.Origin
}

type ValidationErrors []error
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

func (verr ValidationErrors) Unwrap() error {
	if len(verr) != 0 {
        for _, e := range verr {
			return e
		}
	}
	return nil
}

func (verr ValidationErrors) Is(target error) bool {
	for _, e := range verr {
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
	// Реализуй меня.
	resErrors := ValidationErrors{}
	_, err := regexp.Compile(r.Exp)
	if err != nil {
		regErr := ErrIsNotRegexp{Msg: "exp is not regexp: ", Origin: err}
		// regErr := ErrIsNotRegexp{Msg: "error parsing regexp: ", Origin: err}
		resErrors = append(resErrors, regErr)
	}
	if r.Page <= 0 {
		pageErr := ErrInvalidPage{Msg: "invalid page: ", Page: r.Page}
		resErrors = append(resErrors, pageErr)
	}
	if r.PageSize <= 0 {
		pageSizeErr := ErrInvalidPageSize{Msg: "invalid page size: ", Comp: "<", Val: "0", PageSize: r.PageSize}
		resErrors = append(resErrors, pageSizeErr)
	}
	if r.PageSize > maxPageSize {
		pageSizeErr := ErrInvalidPageSize{Msg: "invalid page size: ", Comp: ">", Val: fmt.Sprintf("%d", maxPageSize), PageSize: r.PageSize}
		resErrors = append(resErrors, pageSizeErr)
	}
	if len(resErrors) != 0 {
		return resErrors
	}
	return nil
}
