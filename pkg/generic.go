package pkg

import "github.com/goforj/godump"

type Model[T any] struct {
	Data T
}

func (m *Model[T]) GetData() T {
	return m.Data
}
func (m *Model[T]) SetData(data T) {
	m.Data = data
}

func RunGeneric() {
	model := Model[string]{Data: "Hello, Generics!"}
	data := model.GetData()
	godump.Dump(model)
	model.SetData("Updated Data")
	godump.Dump(model)
	data = model.GetData()
	godump.Dump(data)
	// godump.Dump(model.GetData())
	// godump.Dump(model.Data)
	// model.SetData("New Data")
	// godump.Dump(model.GetData())
}
