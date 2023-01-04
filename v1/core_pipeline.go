package fp

type (
	Pipeline[I, R any] struct {
		workers []*Worker
	}
)

func (itr *Pipeline[I, R]) Execute(items []I) R {
	var result any = items

	for _, worker := range itr.workers {
		result = worker.Execute(result)
	}

	return result.(R)
}
