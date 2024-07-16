package errplus

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeverity(t *testing.T) {
	testCases := []struct {
		desc          string
		err           error
		expectedError Error
	}{
		{
			desc: "newf",
			err:  Normal.Newf("foo %s", "bar"),
			expectedError: Error{
				Severity: Normal,
				Err:      errors.New("foo bar"),
			},
		},
		{
			desc: "new",
			err:  Normal.New("fubar"),
			expectedError: Error{
				Severity: Normal,
				Err:      errors.New("fubar"),
			},
		},
		{
			desc: "wrapf",
			err:  Caution.Wrapf(errors.New("foo"), "bar %s", "baz"),
			expectedError: Error{
				Severity: Caution,
				Err:      fmt.Errorf("bar baz: %w", errors.New("foo")),
			},
		},
		{
			desc: "wrap",
			err:  Critical.Wrap(errors.New("foo"), "bar"),
			expectedError: Error{
				Severity: Critical,
				Err:      fmt.Errorf("bar: %w", errors.New("foo")),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, tc.expectedError, tc.err)
		})
	}
}
