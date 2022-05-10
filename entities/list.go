package entities

type List struct {
	ID       int32
	CreateAt string
	Name     string
	Tasks    []Task
}
