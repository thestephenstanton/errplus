package errplus

import (
	"errors"
	"fmt"
)

type Severity int

const (
	Normal Severity = iota + 1
	Caution
	Critical
)

func (s Severity) Newf(format string, args ...interface{}) Error {
	return Error{
		Severity: s,
		Err:      fmt.Errorf(format, args...),
	}
}

func (s Severity) New(message string) Error {
	return Error{
		Severity: s,
		Err:      errors.New(message),
	}
}

func (s Severity) Wrapf(err error, format string, args ...interface{}) Error {
	message := fmt.Sprintf(format, args...)

	return Error{
		Severity: s,
		Err:      fmt.Errorf("%s: %w", message, err),
	}
}

func (s Severity) Wrap(err error, message string) Error {
	return Error{
		Severity: s,
		Err:      fmt.Errorf("%s: %w", message, err),
	}
}
