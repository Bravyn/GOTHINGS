packgage main

import (
	"bufio",
	"fmt",
	"os",
	"strings"
)

type Task struct {
	Description: string
	Completed bool
}

var tasks []Task
//function to add task
func addTask(description string){
	tasks = append(tasks)
}