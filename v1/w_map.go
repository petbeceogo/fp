package fp

func Map[I, R any](mapper func(item I, index int) (result R)) *Worker {
	return &Worker{
		task: func(input any) any {
			items := input.([]I)
			results := []R{}
			for i, item := range items {
				results = append(results, mapper(item, i))
			}

			return results
		},
	}
}
