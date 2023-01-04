package fp_test

import (
	"testing"

	"github.com/petbeceogo/fp/v1"
)

type (
	MapGiven struct {
		Value int
	}

	MapResult struct {
		Value int
	}
)

func TestMap(t *testing.T) {
	t.Run("map array items A to B", func(t *testing.T) {
		givens := []MapGiven{
			{1}, {2}, {3}, {4}, {5}, {6},
		}

		results := fp.Map(func(given MapGiven, index int) MapResult {
			return MapResult{
				Value: given.Value + 10,
			}
		}).Execute(givens).([]MapResult)

		if results[0].Value != 11 ||
			results[1].Value != 12 ||
			results[2].Value != 13 ||
			results[3].Value != 14 ||
			results[4].Value != 15 ||
			results[5].Value != 16 {
			t.Fail()
			return
		}
	})
}
