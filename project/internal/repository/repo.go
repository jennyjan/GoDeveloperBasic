package repository

import (
	"project/internal/models"
)

var TaskSlice []interface{}
var TagSlice []interface{}

func CreateSlice(item any) {
	switch item.(type) {
	case models.Task:
		TaskSlice = append(TaskSlice, item)
	case models.Tag:
		TagSlice = append(TagSlice, item)
	}
}
