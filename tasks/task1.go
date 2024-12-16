package tasks

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func task1(taskID string) {
	task := GetTask(taskID)
	if task == nil {
		return
	}

	log.Println("In progress... " + taskID)
	task.Status = "in_progress"
	filename := "export_" + taskID + ".json"
	task.Filename = filename

	// Эмуляция длительного процесса
	time.Sleep(5 * time.Second)

	// Запись данных в файл
	file, err := os.Create(filename)
	if err != nil {
		task.Status = "error"
		log.Println("Error of creating file in task: " + taskID)
		return
	}
	defer file.Close()
	log.Println("Successfully created file: " + filename)

	data := map[string]string{"message": "Data exported successfully"}
	if err := json.NewEncoder(file).Encode(data); err != nil {
		log.Println("Error of writing in file " + filename)
		return
	}
	log.Println("Successfully write in file: " + filename)

	task.Status = "done"
}
