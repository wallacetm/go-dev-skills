package main

import (
	"flag"

	"github.com/wallacetm/go-dev-skills/task"
)

func main() {
	taskPtr := flag.Int("task", 0, "task number to execute (1, 2 or 3)")
	flag.Parse()
	if taskPtr == nil {
		panic("task number is required")
	}
	taskNum := *taskPtr
	t := task.NewTask()
	switch taskNum {
	case 1:
		t.Task1()
	case 2:
		t.Task2()
	case 3:
		panic("task 3 is not implemented yet")
		// t.Task3()
	default:
		panic("task number must be 1, 2 or 3")
	}
}
