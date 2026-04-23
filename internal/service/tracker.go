package service

import (
	"encoding/json"
	"fmt"
	"os"
	"task-tracker/model"
)

type Tracker struct {
	Tasks []model.Task
}

func NewTracker() (*Tracker, error) {
	// check data.json exists or not
	if _, err := os.Stat("data.json"); os.IsNotExist(err) {
		// create data.json
		file, err := os.Create("data.json")
		if err != nil {
			fmt.Println("Error creating data.json:", err)
			return nil, err
		}
		file.WriteString("[]")
		defer file.Close()
	}

	data, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading data.json:", err)
		return nil, err
	}

	var tasks []model.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		panic(err)
	}

	return &Tracker{
		Tasks: tasks,
	}, nil
}

func (t *Tracker) AddTask(task model.Task) {
	// set ID for the new task with last ID + 1
	if len(t.Tasks) == 0 {
		task.ID = 1
	} else {
		task.ID = t.Tasks[len(t.Tasks)-1].ID + 1
	}
	t.Tasks = append(t.Tasks, task)
}

func (t *Tracker) UpdateTask(id int, updatedTask model.Task) {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks[i] = updatedTask
			return
		}
	}
}

func (t *Tracker) DeleteTask(id int) {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			return
		}
	}
}

func (t *Tracker) GetTask(id int) *model.Task {
	for _, task := range t.Tasks {
		if task.ID == id {
			return &task
		}
	}
	return nil
}

func (t *Tracker) ListTasks() []model.Task {
	return t.Tasks
}

func (t *Tracker) ListTasksByStatus(status string) []model.Task {
	var filteredTasks []model.Task
	for _, task := range t.Tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}

func (t *Tracker) Save() error {
	data, err := json.MarshalIndent(t.Tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling tasks:", err)
		return err
	}

	err = os.WriteFile("data.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing to data.json:", err)
		return err
	}

	return nil
}
