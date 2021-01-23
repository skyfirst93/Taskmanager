package main

import (
	"Taskmanager/pkg/adder"
	"Taskmanager/pkg/cleaner"
	"Taskmanager/pkg/display"
	"Taskmanager/pkg/executor"
	"Taskmanager/pkg/utils"
	"fmt"
	"sync"
)

func main() {
	//task1 := utils.Task{"Id": 1, "IsCompleted": false, "Status": utils.UnTouched, "CreationTime": time.Now(), "TaskData": "task1"}
	//task1 := utils.Task{"1", false, utils.UnTouched, time.Now(), "task1"}
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("ADDED_TASK_ID  |EXECUTION_TASK_ID |CLEANER_TASK_ID|ITEAMS_IN_QUEUE")
	fmt.Println("------------------------------------------------------------------")
	var waitGroup sync.WaitGroup
	waitGroup.Add(4)
	go adder.Adder(waitGroup)
	go executor.Executor(waitGroup)
	go cleaner.Cleaner(waitGroup)
	go display.Display(waitGroup)
	// Setup Ctrl+C handler
	utils.SetupCloseHandler()
	waitGroup.Wait()
	// Goroutine to check the task to be commpleted
	//Goroutine to clean the completed task
}
