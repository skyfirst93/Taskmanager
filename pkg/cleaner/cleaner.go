package cleaner

import (
	"Taskmanager/pkg/utils"
	"fmt"
	"sync"
	"time"
)

// Cleaner is responsible for removing a task from the task queue
// or if the it is taking to much time move it to the end of the queue
func Cleaner(waitGroup sync.WaitGroup) {
	for {
		// check if the queue is empty if empty pass
		if len(utils.TaskQueue) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Printf("               |                  |  %s            |", utils.TaskQueue[0].Id)
		fmt.Println("")
		if utils.TaskQueue[0].IsCompleted == true {
			fmt.Println("Task Completed. Task Id", utils.TaskQueue[0].Id)
			utils.TaskQueue = utils.TaskQueue[1:]
		} else if utils.TaskQueue[0].Status == utils.Failed {
			fmt.Println("")
			fmt.Println("Task failed to be moved to the end for retrying :", utils.TaskQueue[0].Id)
			utils.TaskQueue = append(utils.TaskQueue, utils.TaskQueue[0])
			utils.TaskQueue = utils.TaskQueue[1:]

		} else if utils.TaskQueue[0].Status == utils.Timeout {
			fmt.Println("")
			fmt.Println("Task have timed out.Task will removed. Task ID :", utils.TaskQueue[0].Id)
			utils.TaskQueue = utils.TaskQueue[1:]

		}
		time.Sleep(5 * time.Second)
	}
	waitGroup.Done()
}
