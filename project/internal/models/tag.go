package models

type Tag struct {
	id   int
	name string
}

func NewTag(name string) Tag {
	return Tag{
		name: name,
	}
}

func (tag Tag) GetId() int {
	return tag.id
}
