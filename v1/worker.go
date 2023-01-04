package fp

type (
	Task func(input any) any

	Worker struct {
		task Task
	}
)

func (w *Worker) Execute(input any) any {
	return w.task(input)
}
