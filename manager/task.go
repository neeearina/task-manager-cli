package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct{
	ID int 
	Title string
	Done bool
}

// RewriteTasks полностью перезаписывает задачи в файл tasks.json после их обновления.
func RewriteTasks(tasks []Task) {
	file, err := os.Create(FilePath)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		os.Exit(1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Ошибка записи задач в файл:", err)
		os.Exit(1)
	}
}

// AddTask добавляет новую задачу в слайс задач и перезаписывает файл tasks.json.
func AddTask(args []string, tasks []Task, nextID int) {
	fmt.Println("Добавляем задачу...")

	if len(args) == 0 {
		fmt.Println("Ошибка: нет названия задачи.")
		return
	}

	fullTitle := strings.Join(args, " ")
	newTask := Task{ID: nextID, Title: fullTitle, Done: false}
	tasks = append(tasks, newTask)
	RewriteTasks(tasks)

	fmt.Printf("Добавлена новая задача '%s' с ID %d.\n", fullTitle, nextID)
}

// ListTasks выводит список всех существующих задач.
func ListTasks(tasks []Task) {
	fmt.Println("Выводим список задач...")

	if len(tasks) == 0 {
		fmt.Println("Список задач пуст, вы можете их добавить.")
		return
	}

    for _, task := range tasks {
		if task.Done {
			fmt.Printf("[%d] %s - [x] Выполнена.\n", task.ID, task.Title)
		} else {
			fmt.Printf("[%d] %s - [ ] Не выполнена.\n", task.ID, task.Title)
		}
	}
}

// MarkTaskDone отмечает задачу выполненной по её ID и перезаписывает файл tasks.json.
func MarkTaskDone(args []string, tasks []Task) {
	fmt.Println("Отмечаем задачу как выполненную...")

	if len(args) == 0 {
		fmt.Println("Ошибка: не передан ID задачи.")
		return
	}

	intIDArg, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Ошибка: ID задачи должен быть числом.")
		return
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == intIDArg {
			tasks[i].Done = true
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Ошибка: задача с ID %d не найдена.", intIDArg)
		return
	}

	RewriteTasks(tasks)

	fmt.Printf("Задача с ID %d выполнена.\n", intIDArg)
}

// DeleteTask удаляет задачу по её ID из списка задач и перезаписывает файл tasks.json.
func DeleteTask(args []string, tasks []Task) {
	fmt.Println("Удаляем задачу...")

	if len(args) == 0 {
		fmt.Println("Ошибка: не передан ID задачи.")
		return
	}

	intIDArg, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Ошибка: ID задачи должен быть числом.")
		return
	}

	found := false
	for i, task := range tasks {
		if task.ID == intIDArg {
			tasks[i] = tasks[len(tasks)-1] 
			tasks = tasks[:len(tasks)-1] 
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Ошибка: задача с ID %d не найдена.", intIDArg)
		return
	}

	RewriteTasks(tasks)

	fmt.Printf("Задача с ID %d успешно удалена.\n", intIDArg)
}