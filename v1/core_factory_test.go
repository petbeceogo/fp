package fp_test

import (
	"testing"

	"github.com/petbeceogo/fp/v1"
)

type (
	FactoryGiven struct {
		Value int
	}
)

func TestFactory(t *testing.T) {
	t.Run("works multiple pipe", func(t *testing.T) {
		givens := []FactoryGiven{
			{1}, {2}, {3}, {4}, {5}, {6}, {7},
		}
		expected := 58

		result := fp.Pipe[FactoryGiven, int]([]*fp.Worker{
			fp.Filter(func(item FactoryGiven, index int) bool {
				return item.Value > 2 && item.Value < 7
			}),
			fp.Map(func(item FactoryGiven, index int) int {
				return item.Value + 10
			}),
			fp.Reduce(func(sum int, item int, index int) int {
				return sum + item
			}, func() int {
				return 0
			}),
		}).Execute(givens)

		if expected != result {
			t.Fail()
			return
		}
	})
}
