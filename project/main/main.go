package main

import (
	"fmt"

	"project/internal/models"
)

func main() {
	fmt.Println(models.NewTask("new task", "all"))
	fmt.Println(models.GetTask())
	fmt.Println(models.NewTag("new tag"))
	fmt.Println(models.GetTag())
}
