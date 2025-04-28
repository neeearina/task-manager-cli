package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct{
	ID int 
	Title string
	Done bool
}

// RewriteTasks полностью перезаписывает задачи в файл tasks.json после их обновления.
func rewriteTasks(tasks []Task) {
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
func addTask(tasks []Task, nextID int, fullTitle string) {
	fmt.Println("Добавляем задачу...")

	newTask := Task{ID: nextID, Title: fullTitle, Done: false}
	tasks = append(tasks, newTask)
	rewriteTasks(tasks)

	fmt.Printf("Добавлена новая задача '%s' с ID %d\n", fullTitle, nextID)
}

// ListTasks выводит список всех существующих задач.
func listTasks(tasks []Task) {
	fmt.Println("Выводим список задач...")

    for _, task := range tasks {
		if task.Done {
			fmt.Printf("[%d] %s - [x] Выполнена\n", task.ID, task.Title)
		} else {
			fmt.Printf("[%d] %s - [ ] Не выполнена\n", task.ID, task.Title)
		}
	}
}

// MarkTaskDone отмечает задачу выполненной по её ID и перезаписывает файл tasks.json.
func markTaskDone(tasks []Task, intIDArg int) {
	fmt.Println("Отмечаем задачу как выполненную...")

	found := false
	for i := range tasks {
		if tasks[i].ID == intIDArg {
			tasks[i].Done = true
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Ошибка: задача с ID %d не найдена", intIDArg)
		return
	}

	rewriteTasks(tasks)

	fmt.Printf("Задача с ID %d выполнена\n", intIDArg)
}

// DeleteTask удаляет задачу по её ID из списка задач и перезаписывает файл tasks.json.
func deleteTask(tasks []Task, intIDArg int) {
	fmt.Println("Удаляем задачу...")

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
		fmt.Printf("Ошибка: задача с ID %d не найдена", intIDArg)
		return
	}

	rewriteTasks(tasks)

	fmt.Printf("Задача с ID %d успешно удалена\n", intIDArg)
}