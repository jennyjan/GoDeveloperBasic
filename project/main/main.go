package main

import (
	"fmt"

	"project/internal/repository"
	"project/internal/service"
)

func main() {

	for i := 0; i < 10; i++ {
		service.CreateModel()
	}

	fmt.Println(repository.TaskSlice)
	fmt.Println(repository.TagSlice)
}
