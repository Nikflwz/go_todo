package main

import "fmt"

type Task struct {
	ID        int
	Title     string
	Done      bool
	CreatedAt string
}

func main() {

	tasks := []Task{
		{
			ID:        1,
			Title:     "Сходить в магазин",
			Done:      false,
			CreatedAt: "17.01.2026",
		},
		{
			ID:        2,
			Title:     "Купить яйца",
			Done:      false,
			CreatedAt: "17.01.2026",
		},
		{
			ID:        3,
			Title:     "Сходить на OZON",
			Done:      true,
			CreatedAt: "17.01.2026",
		},
		{
			ID:        4,
			Title:     "Позаниматься с языком GO",
			Done:      false,
			CreatedAt: "17.01.2026",
		},
	}

	fmt.Println("=== Список задач ===")

	for i, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[✓]"
		}
		fmt.Printf("%v %d. %s [ Создано: %s]\n", status, i+1, task.Title, task.CreatedAt)
	}

	tasks[0].Done = true

	newTask := Task{
		ID:        5,
		Title:     "Вынести мусор",
		Done:      true,
		CreatedAt: "17.01.2026",
	}

	tasks = append(tasks, newTask)

	fmt.Println("\n\nВыполненные задачи: ")
	for i, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[✓]"
			fmt.Printf("%v %d. %s [ Создано: %s]\n", status, i+1, task.Title, task.CreatedAt)
		}

	}

	fmt.Println("\n\nНевыполненные задачи: ")
	for i, task := range tasks {
		status := "[ ]"
		if !task.Done {
			status = "[ ]"
			fmt.Printf("%v %d. %s [ Создано: %s]\n", status, i+1, task.Title, task.CreatedAt)
		}

	}
}
