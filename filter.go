package fp

const FilterName = "filter"

type FilterCallback = func(item interface{}, index int) bool

type FilterPipe interface {
	Callback() FilterCallback
}

type filterPipe struct {
	callback FilterCallback
}

func (p *filterPipe) Callback() FilterCallback {
	return p.callback
}

func NewFilterPipe(callback FilterCallback) Pipe {
	return NewPipe(FilterName, &filterPipe{
		callback: callback,
	})
}

func Filter[I interface{}](items []I, callback FilterCallback) []I {
	var result []I
	for i := range items {
		if callback(items[i], i) {
			result = append(result, items[i])
		}
	}
	return result
}
