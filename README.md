# Taskmanager

### Problem Statement
Implement a Task queue. Have a Task struct and create Task objects and push them to a
queue.
1. Create a go-routine which keeps working on the tast to complete them.
2. Have another go-routine which cleaning the queue. It keeps checking the tasks
queue and inspect if tasks are completed or not. If the task is completed then
remove it from the queue, if its not completed push back at the end of the queue.
If the task is not completed after a certain amount of time then it should be
marked as a timeout, removed from the queue and the event should be logged.
Implement such a system with proper error handling.
Example of Task Struct:
```
type Task struct {
Id string
IsCompleted bool
Status string // untouched, completed, failed, timeout
CreationTime time.Time // when was the task created
TaskData string // field containing data about the task
}
```
You are free to add, remove or alter the fields in the Task struct, if you so desire.
### Example Scenario
Imagine a platform where user interactions cause Email notifications to the interacting
user or other users. For example, on a social networking site, liking or commenting on
someone’s photo, or logging in or out causes an email to be dispatched to the interacting
user or the user who posted the photo.
Let’s assume that our task management system would be used in such a scenario and the
queue would contain tasks containing emails to be dispatched.
In such a scenario, there would be 3 separate workers on the queue that do the following
work:
1. Adder - Adds a task to the queue. Liking or commenting on a picture would cause
this worker to add an email to the queue.
2. Executor - This worker picks up the task and executes them. The system would
attempt to send the email to the intended recipient. If the email fails to be sent,
the executor marks the task as a failure.
3. Cleaner - Cleans up the queue depending on task success, fail or timeout status.
We don’t want emails that have already been sent successfully to be clogging up
the tasks queue. If there are emails that have failed recently, then we want to retry,
so we put those emails at the end of the queue. And if the system has been trying
to send the email since last 5 minutes but it has been failing, we need to prevent
clogging up the queue so we remove it from the queue as and log the task timeout
event.

### OUTPUT
![image](https://user-images.githubusercontent.com/24405247/105570869-c6078000-5d71-11eb-9fb2-acb8b52d6443.png)
