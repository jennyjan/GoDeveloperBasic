package service

import (
	"project/internal/models"
	"project/internal/repository"
)

func CreateModel() {
	task := models.NewTask("new task", "health")
	tag := models.NewTag("health")
	repository.CreateSlice(task)
	repository.CreateSlice(tag)
}
