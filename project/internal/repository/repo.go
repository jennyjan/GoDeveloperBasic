package repository

import (
	"log"
	"sync"
	"time"

	"project/internal/models"
)

type Item interface {
	GetId() int
}

var TaskSlice []*models.Task
var TagSlice []*models.Tag

var TaskMutex sync.Mutex
var TagMutex sync.Mutex

var lastTaskCount int = 0
var lastTagCount int = 0

func FillSlice(ch chan Item, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range ch {
		switch value := item.(type) {
		case *models.Task:
			TaskMutex.Lock()
			TaskSlice = append(TaskSlice, value)
			TaskMutex.Unlock()
		case *models.Tag:
			TagMutex.Lock()
			TagSlice = append(TagSlice, value)
			TagMutex.Unlock()
		}
	}
}

func LogChanges(wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			TaskMutex.Lock()
			currentTaskCount := len(TaskSlice)
			if currentTaskCount > lastTaskCount {
				addedTasks := TaskSlice[lastTaskCount:currentTaskCount]
				for _, task := range addedTasks {
					log.Printf("New Task Added: %+v", *task)
				}
				lastTaskCount = currentTaskCount
			}
			TaskMutex.Unlock()

			TagMutex.Lock()
			currentTagCount := len(TagSlice)
			if currentTagCount > lastTagCount {
				addedTags := TagSlice[lastTagCount:currentTagCount]
				for _, tag := range addedTags {
					log.Printf("New Tag Added: %+v", *tag)
				}
				lastTagCount = currentTagCount
			}
			TagMutex.Unlock()
		}
	}
}
