package errs

import (
	 "fmt"
	"time"
)

type WithTimeError struct {
	wterr       error
	errocTime time.Time
}

func (t WithTimeError) Time() time.Time {
	return t.errocTime
}

func (t WithTimeError) Error() string {
	return fmt.Errorf("%w", t.wterr).Error()
}

func (t WithTimeError) Unwrap() error {
    return t.wterr
}

func NewWithTimeError(err error) error {
	return newWithTimeError(err, time.Now)
}

func newWithTimeError(err error, timeFunc func() time.Time) error {
    timeOcured := timeFunc()
	e := WithTimeError{
		wterr: err,
        errocTime: timeOcured,
	}
	return e
}
