package current_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/prymitive/current"
	"github.com/stretchr/testify/require"
)

type store struct {
	data []any
}

func (s *store) push(v any) {
	s.data = append(s.data, v)
}

func (s *store) reset() {
	s.data = []any{}
}

type testCaseT struct {
	name     string
	iter     current.Iterator
	body     string
	expected []any
	err      string
}

func runTestCase(t *testing.T, tc testCaseT, got *store) {
	t.Run(tc.name, func(t *testing.T) {
		got.reset()
		dec := json.NewDecoder(strings.NewReader(tc.body))
		err := current.Stream(dec, tc.iter)
		if tc.err != "" {
			require.EqualError(t, err, tc.err)
		} else {
			require.NoError(t, err)
			require.ElementsMatch(t, tc.expected, got.data)
		}
	})
}
