package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const FilePath = "tasks.json"

// FindNextID находит максимальный ID среди существующих задач и возвращает следующий доступный ID.
func findNextID(tasks []Task) int {
	nextID := 0
	for _, task := range tasks {
		if task.ID > nextID {
			nextID = task.ID
		}
	}

	return nextID + 1
}

// LoadTasks загружает задачи из файла tasks.json, возвращает задачи в виде слайса и следующий доступный ID.
func loadTasks() ([]Task, int, error) {
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

	lastID := findNextID(tasks)

	return tasks, lastID, nil
}

// StartCommand обрабатывает команду, переданную через аргументы командной строки,
// и вызывает соответствующую функцию для выполнения задачи: add, list, done, delete, help.
func startCommand(commands []string, tasks []Task, nextID int) {
	args := commands[2:]
	switch commands[1] {
	case "add":
		if len(args) == 0 {
			fmt.Println("Ошибка: нет названия задачи")
			return
		}		

		fullTitle := strings.Join(args, " ")
		addTask(tasks, nextID, fullTitle)
	case "list":
		if len(tasks) == 0 {
			fmt.Println("Список задач пуст, вы можете их добавить")
			return
		}

		listTasks(tasks)
	case "done":
		if len(args) == 0 {
			fmt.Println("Ошибка: не передан ID задачи")
			return
		}

		intIDArg, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Ошибка: ID задачи должен быть числом")
			return
		}

		markTaskDone(tasks, intIDArg)
	case "delete":
		if len(args) == 0 {
			fmt.Println("Ошибка: не передан ID задачи")
			return
		}

		intIDArg, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Ошибка: ID задачи должен быть числом")
			return
		}
		
		deleteTask(tasks, intIDArg)
	case "help":
		fmt.Println("Список команд: add, list, done, delete, help")
	default:
		fmt.Printf("Неизвестная команда: %s\n", commands[1])
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Ошибка: команда не указана. Используйте 'add', 'list', 'done', 'delete' или 'help'")
		return
	}

	tasks, nextID, err := loadTasks()
	if err != nil {
		fmt.Println("Ошибка загрузки задач:", err)
		return
	}

	startCommand(os.Args, tasks, nextID)
}