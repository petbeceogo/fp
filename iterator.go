package fp

type Evaluator[I interface{}, R interface{}] interface {
	Map(callback MapCallback) Evaluator[I, R]
	Filter(callback FilterCallback) Evaluator[I, R]
	Set(pipe Pipe) Evaluator[I, R]
	Eval() []R
}

type evaluator[I interface{}, R interface{}] struct {
	items []I
	pipes []Pipe
}

func (e *evaluator[I, R]) Map(callback MapCallback) Evaluator[I, R] {
	return e.Set(NewMapPipe(callback))
}

func (e *evaluator[I, R]) Filter(callback FilterCallback) Evaluator[I, R] {
	return e.Set(NewFilterPipe(callback))
}

func (e *evaluator[I, R]) Set(pipe Pipe) Evaluator[I, R] {
	// TODO pipe validate check
	e.pipes = append(e.pipes, pipe)
	return e
}

func (e *evaluator[I, R]) Eval() []R {
	var slice []R

	for i := range e.items {
		var item interface{} = e.items[i]
		noAppend := false

		for j := range e.pipes {
			if noAppend {
				continue
			}
			pipe := e.pipes[j]
			pipeName := pipe.Name()
			workPipe := pipe.Val()

			switch pipeName {
			case FilterName:
				noAppend = !workPipe.(FilterPipe).Callback()(item, i)
			case MapName:
				item = workPipe.(MapPipe).Callback()(item, i)
			}
		}

		if !noAppend {
			slice = append(slice, item.(R))
		}
	}

	return slice
}

func Of[I interface{}, R interface{}](items []I) Evaluator[I, R] {
	return &evaluator[I, R]{items: items}
}
