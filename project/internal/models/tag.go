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

func GetTag() Tag {
	return Tag{
		id:   1,
		name: "Tag",
	}
}
