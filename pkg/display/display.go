package display

import (
	"Taskmanager/pkg/utils"
	"fmt"
	"sync"
	"time"
)

// Display is responsible for printing all the elment in the queue
// after every 5 seconds
func Display(waitGroup sync.WaitGroup) {
	for {
		// check if the queue is empty if empty pass
		if len(utils.TaskQueue) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Print("               |                  |               |")
		for _, value := range utils.TaskQueue {
			fmt.Print(" ", value.Id)
		}
		fmt.Println("")
		time.Sleep(10 * time.Second)
	}
	waitGroup.Done()

}
