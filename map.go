package fp

const MapName = "map"

type MapCallback = func(item interface{}, index int) interface{}

type MapPipe interface {
	Callback() MapCallback
}

type mapPipe struct {
	callback MapCallback
}

func (p *mapPipe) Callback() MapCallback {
	return p.callback
}

func NewMapPipe(callback MapCallback) Pipe {
	return NewPipe(MapName, &mapPipe{
		callback: callback,
	})
}

func Map[I interface{}, R interface{}](items []I, callback MapCallback) []R {
	result := make([]R, len(items))
	for i := range items {
		result[i] = callback(items[i], i).(R)
	}
	return result
}
