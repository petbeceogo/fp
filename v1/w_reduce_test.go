package fp_test

import (
	"testing"

	"github.com/petbeceogo/fp/v1"
)

type (
	ReduceGiven struct {
		Value int
	}
)

func TestReduce(t *testing.T) {
	t.Run("get result accumulating array element process", func(t *testing.T) {
		givens := []ReduceGiven{
			{1}, {2}, {3}, {4}, {5},
		}
		expected := 15

		result := fp.Reduce(func(ac int, given ReduceGiven, index int) int {
			return ac + given.Value
		}, func() int {
			return 0
		}).Execute(givens).(int)

		if expected != result {
			t.Fail()
			return
		}
	})
}
