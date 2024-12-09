package tasks

import (
	"errors"
	"log"
	"strconv"
	"sync"
	"time"
)

type TaskFunc func(taskID string)

func getFunctionByName(name string) interface{} {
	funcs := map[string]interface{}{
		"task1": TaskFunc(task1),
	}

	if fn, exists := funcs[name]; exists {
		return fn
	}
	return nil
}

type Task struct {
	ID           string   `json:"id"`
	Status       string   `json:"status"` // pending, in_progress, done, error
	Filename     string   `json:"filename"`
	TaskID       int      `json:"task_id"`
	TaskFunction TaskFunc `json:"-"`
}

var tasks = make(map[string]*Task)
var mutex = &sync.Mutex{}

func CreateTask(taskID int) (string, error) {
	kTasks := 0
	for _, task := range tasks {
		if task.Status == "pending" || task.Status == "in_progress" {
			kTasks++
		}
	}
	if kTasks > 5 {
		return "", errors.New("the task limit has been exceeded")
	}

	uniqueID := generateID()

	taskName := "task" + strconv.Itoa(taskID)
	function := getFunctionByName(taskName)
	if function == nil {
		return "", errors.New("function has incorrect name: " + taskName)
	}

	ptr, ok := function.(TaskFunc)
	if !ok {
		return "", errors.New("function has incorrect signature: " + taskName)
	}

	task := &Task{
		ID:           uniqueID,
		TaskID:       taskID,
		Status:       "pending",
		TaskFunction: ptr,
	}
	mutex.Lock()
	tasks[uniqueID] = task
	mutex.Unlock()
	log.Println("Successfully created task: " + uniqueID)
	return uniqueID, nil
}

func GetTask(taskID string) *Task {
	mutex.Lock()
	defer mutex.Unlock()
	return tasks[taskID]
}

func RunTask(uniqueID string) string {
	task := GetTask(uniqueID)
	if task == nil || task.TaskFunction == nil {
		return "task not found"
	}

	log.Println("Start task: " + uniqueID)
	go task.TaskFunction(uniqueID)
	return "task has started"
}

func generateID() string {
	return time.Now().Format("20060102150405")
}
