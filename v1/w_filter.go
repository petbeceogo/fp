package fp

func Filter[I any](predecate func(item I, index int) (leave bool)) *Worker {
	return &Worker{
		task: func(input any) any {
			items := input.([]I)
			results := []I{}
			for i, item := range items {
				if predecate(item, i) {
					results = append(results, item)
				}
			}

			return results
		},
	}
}
