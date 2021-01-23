package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Task structure contains the data managers that make up the task
type Task struct {
	// Id is the Unique ID assigned to each task
	Id string
	// IsCompleted states if a task is completed or not
	IsCompleted bool
	// Status gives the actually status of a task at a particular point of time
	// possible value are untouched, completed, failed, timeout
	Status string
	// CreationTime is the time at which a task was added into the system
	CreationTime time.Time
	// TaskData contains the data about the task
	TaskData string
}

// TaskQueue is queue that maintains the of the task added into the system
var TaskQueue []Task

const (
	// UnTouched is the state of a task for which execution is not started
	UnTouched = "untouched"
	// Completed is the state of a task for which execution was successfully
	Completed = "completed"
	// Failed is the state of a task for which execution failed
	Failed = "failed"
	// Timeout is the state of a task for which execution timed out
	Timeout = "timeout"
)

var (
	// DefaultExecutionTime is the time in which the execution of a task
	// should get completed
	DefaultExecutionTime = 5
)

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}
