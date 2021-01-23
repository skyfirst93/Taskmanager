package adder

import (
	"Taskmanager/pkg/utils"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Adder is responsible for adding a task to the task queue
func Adder(waitGroup sync.WaitGroup) {
	taskCounter := 1
	for {
		newTaskID := strconv.Itoa(taskCounter)
		task := utils.Task{newTaskID, false, utils.UnTouched, time.Now(), "taskdata"}
		utils.TaskQueue = append(utils.TaskQueue, task)
		fmt.Println("", task.Id)
		taskCounter = taskCounter + 1
		time.Sleep(5 * time.Second)
	}
	waitGroup.Done()
}

/*
func generateTaskID() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 8
	var id strings.Builder
	for i := 0; i < length; i++ {
		id.WriteRune(chars[rand.Intn(len(chars))])
	}
	randomId := id.String()
	return randomId
}
*/
