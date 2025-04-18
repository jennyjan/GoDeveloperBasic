package main

import (
	"fmt"
	"sync"

	"project/internal/repository"
	"project/internal/service"
)

func main() {

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go repository.LogChanges(wg)

	for i := 0; i < 10; i++ {
		service.CreateModel()
	}

	fmt.Println(repository.TaskSlice)
	fmt.Println(repository.TagSlice)

	wg.Wait()
}
