package errplus

import (
	"errors"
)

type Error struct {
	Err      error
	Severity Severity
	Args     []any
}

func ToError(err error) Error {
	var e Error
	ok := errors.As(err, &e)
	if !ok {
		return Error{
			Severity: Normal,
			Err:      err,
		}
	}

	return e
}

func (e Error) SetSeverity(severity Severity) Error {
	e.Severity = severity

	return e
}

func (e Error) Error() string {
	return e.Err.Error()
}

func (e Error) Unwrap() error {
	return e.Err
}

func (e Error) Add(key string, value any) Error {
	e.Args = append(e.Args, key, value)

	return e
}

func Newf(format string, args ...interface{}) Error {
	return Normal.Newf(format, args...)
}

func New(message string) Error {
	return Normal.New(message)
}

func Wrapf(err error, format string, args ...interface{}) Error {
	return Normal.Wrapf(err, format, args...)
}

func Wrap(err error, message string) Error {
	return Normal.Wrap(err, message)
}
