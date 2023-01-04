package fp

func Reduce[I, R any](
	reducer func(accumulator R, item I, index int) R,
	initializer func() R,
) *Worker {
	return &Worker{
		task: func(input any) any {
			var result R = initializer()
			items := input.([]I)

			for i, item := range items {
				result = reducer(result, item, i)
			}

			return result
		},
	}
}
