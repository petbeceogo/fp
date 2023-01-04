package fp_test

import (
	"testing"

	"github.com/petbeceogo/fp/v1"
)

type (
	FilterGiven struct {
		Value int
	}
)

func TestFilter(t *testing.T) {
	t.Run("filter elements in array", func(t *testing.T) {
		givens := []FilterGiven{
			{1}, {2}, {3}, {4}, {5}, {6},
		}

		results := fp.Filter(func(given FilterGiven, index int) bool {
			return given.Value > 2 && given.Value < 6
		}).Execute(givens).([]FilterGiven)

		if len(results) != 3 {
			t.Fail()
			return
		}

		if results[0].Value != 3 || results[1].Value != 4 || results[2].Value != 5 {
			t.Fail()
			return
		}
	})
}
