package service

import (
	"sync"

	"project/internal/models"
	"project/internal/repository"
)

func CreateModel() {
	task := models.NewTask("new task", "health")
	tag := models.NewTag("health")

	taskPointer := &task
	tagPointer := &tag

	taskChannel := make(chan repository.Item, 10)
	tagChannel := make(chan repository.Item, 10)

	var wg sync.WaitGroup

	wg.Add(1)
	go repository.FillSlice(taskChannel, &wg)

	wg.Add(1)
	go repository.FillSlice(tagChannel, &wg)

	taskChannel <- taskPointer
	tagChannel <- tagPointer

	close(taskChannel)
	close(tagChannel)

	wg.Wait()
}
