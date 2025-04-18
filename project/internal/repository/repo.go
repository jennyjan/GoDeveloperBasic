package repository

import (
	"project/internal/models"
)

type Item interface {
	GetId() int
}

var TaskSlice []models.Task
var TagSlice []models.Tag

func FillSlice(item Item) {
	switch value := item.(type) {
	case models.Task:
		TaskSlice = append(TaskSlice, value)
	case models.Tag:
		TagSlice = append(TagSlice, value)
	}
}
