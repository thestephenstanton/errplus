package errplus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		desc     string
		args1    []any
		args2    []any
		expected []any
	}{
		{
			desc:     "standard",
			args1:    []any{"foo", "bar"},
			args2:    []any{"fizz", "buzz"},
			expected: []any{"foo", "bar", "fizz", "buzz"},
		},
		{
			desc:     "duplicate",
			args1:    []any{"foo", "bar"},
			args2:    []any{"foo", "buzz"},
			expected: []any{"foo", "bar", "foo-duplicate-1", "buzz"},
		},
		{
			desc:     "empty",
			args1:    []any{},
			args2:    nil,
			expected: []any{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, tc.expected, Merge(tc.args1, tc.args2...))
		})
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Merge([]any{"foo", "bar", "fizz", "buzz", "fubar", "foobar"}, "a", "b", "c", "d", "e", "f")
	}
}
