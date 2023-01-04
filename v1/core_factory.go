package fp

func Pipe[I, R any](workers []*Worker) *Pipeline[I, R] {
	return &Pipeline[I, R]{
		workers: workers,
	}
}
