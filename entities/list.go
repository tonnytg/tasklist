package entities

type List struct {
	ID       int32
	CreateAt string
	Name     string
	Tasks    []Task
}

func NewList() *List {
	l := List{}
	return &l
}
