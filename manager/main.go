package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const FilePath = "tasks.json"

// FindNextID находит максимальный ID среди существующих задач и возвращает следующий доступный ID.
func FindNextID(tasks []Task) int {
	nextID := 0
	for _, task := range tasks {
		if task.ID > nextID {
			nextID = task.ID
		}
	}
	return nextID + 1
}

// LoadTasks загружает задачи из файла tasks.json, возвращает задачи в виде слайса и следующий доступный ID.
func LoadTasks() ([]Task, int, error) {
	file, err := os.OpenFile(FilePath, os.O_CREATE, 0666)
	if err != nil {
		return nil, 0, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		if err == io.EOF {
			return tasks, 0, nil
		}
		return nil, 0, fmt.Errorf("ошибка декодирования JSON: %w", err)
	}

	lastID := FindNextID(tasks)
	return tasks, lastID, nil
}

// StartCommand обрабатывает команду, переданную через аргументы командной строки,
// и вызывает соответствующую функцию для выполнения задачи: add, list, done, delete, help.
func StartCommand(command []string, tasks []Task, nextID int) {
	switch command[1] {
	case "add":
		AddTask(command[2:], tasks, nextID)
	case "list":
		ListTasks(tasks)
	case "done":
		MarkTaskDone(command[2:], tasks)
	case "delete":
		DeleteTask(command[2:], tasks)
	case "help":
		fmt.Println("Список команд: add, list, done, delete, help")
	default:
		fmt.Printf("Неизвестная команда: %s\n", command[1])
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Ошибка: команда не указана. Используйте 'add', 'list', 'done', 'delete' или 'help'.")
		return
	}

	tasks, nextID, err := LoadTasks()
	if err != nil {
		fmt.Println("Ошибка загрузки задач:", err)
		return
	}
	StartCommand(os.Args, tasks, nextID)
}