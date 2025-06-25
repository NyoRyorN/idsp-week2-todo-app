package main

import (
	"encoding/json"
	"os"
	"fmt"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const dataFile = "todo.json"

func loadTasks() ([]Task, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(tasks)
}

func nextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func AddTask(title string) {
	temp_tasks, _ := loadTasks()
	last_id := nextID(temp_tasks)
	new_task := Task{ID:last_id, Title:title, Done:false}
	temp_tasks = append(temp_tasks, new_task)
	_ = saveTasks(temp_tasks)
}

func ListTasks() {
	temp_tasks, _ := loadTasks()
	for i:=0; i<len(temp_tasks); i++ {
		done := " "
		if temp_tasks[i].Done {
			done = "x"
		}
		fmt.Printf("%d: %s [%s] \n", temp_tasks[i].ID, temp_tasks[i].Title, done)
	}
}

func CompleteTask(id int) {
	tasks, err := loadTasks()
	if err == nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	tasks = append(tasks, t1)
	for _, task := range tasks {
		if task.ID == id {
			task.Done = true
			break
		}
	}
	fmt.Println("Task completed:", id)
	saveTasks(tasks)
	panic("unimplemented")
}

func DeleteTask(id int) {
	var updatedTasks []Task
	tasks, err := loadTasks()
	if err == nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	for _, task := range tasks {
		if task.ID != id {
			updatedTasks = append(updatedTasks, task)
		}
	}
	saveTasks(updatedTasks)
	panic("unimplemented")
}
