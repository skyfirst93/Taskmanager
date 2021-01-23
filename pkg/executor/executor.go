package executor

import (
	"Taskmanager/pkg/utils"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var statusList = map[string]string{
	"1": utils.Completed,
	"2": utils.Failed,
	"3": utils.Timeout,
}

// Executor is responsible for excuting the task picked up form the queue and changing the status
// accordingly
func Executor(waitGroup sync.WaitGroup) {
	for {
		// check if the queue is empty if empty pass
		if len(utils.TaskQueue) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}
		//firstTask := utils.TaskQueue[0]
		fmt.Println("               |", utils.TaskQueue[0].Id)
		taskStatus := getStatusRandom()
		utils.TaskQueue[0].Status = statusList[taskStatus]
		if statusList[taskStatus] == utils.Completed {
			utils.TaskQueue[0].IsCompleted = true
		}
		time.Sleep(5 * time.Second)
	}
	waitGroup.Done()
}
func getStatusRandom() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("123")
	var code strings.Builder
	code.WriteRune(chars[rand.Intn(len(chars))])
	randomCode := code.String()
	return randomCode
}
