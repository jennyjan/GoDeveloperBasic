package models

type Task struct {
	id   int
	date string
	name string
	tag  string
}

func NewTask(name string, tag string) Task {
	return Task{
		name: name,
		tag:  tag,
	}
}

func GetTask() Task {
	return Task{
		id:   1,
		name: "New",
	}
}
