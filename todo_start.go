package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Создаем структуру
type Task struct {
	ID        int
	Title     string
	Done      bool
	CreatedAt string
}

// printTasks - Показывает задачи
func printTasks(tasks []Task) {
	fmt.Println("\n=== Список задач ===")

	for i, task := range tasks {
		status := "[ ]"
		if task.Done {
			status = "[✓]"
		}
		fmt.Printf("%v %d. %s [ Создано: %s]\n", status, i+1, task.Title, task.CreatedAt)
	}
}

// completeTask - Показывает выполненные задачи
func completeTask(tasks []Task, id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			if tasks[i].Done {
				fmt.Printf("Задача '%s' уже выполнена\n", tasks[i].Title)
				return
			}
			tasks[i].Done = true
			fmt.Printf("Задача '%s' отмечена как выполненная\n", tasks[i].Title)
			return
		}
	}
	fmt.Printf("Задача с ID %d не найдена\n", id)
}

// uncompleteTask - Показывает не выполненные задачи
func uncompleteTask(tasks []Task, id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			if !tasks[i].Done { // Если уже не выполнена
				fmt.Printf("Задача '%s' уже не выполнена\n", tasks[i].Title)
				return
			}
			tasks[i].Done = false // Меняем на false
			fmt.Printf("Задача '%s' отмечена как НЕ выполненная\n", tasks[i].Title)
			return
		}
	}
	fmt.Printf("Задача с ID %d не найдена\n", id)
}

// addTask - Добавляет новую задачу
func addTask(tasks []Task, title string) []Task {
	// ПРОВЕРКА: если название пустое
	if title == "" {
		fmt.Println("Ошибка: название задачи не может быть пустым!")
		return tasks // Возвращаем задачи без изменений
	}

	newID := 1
	if len(tasks) > 0 {
		// Находим последний ID
		lastTask := tasks[len(tasks)-1]
		newID = lastTask.ID + 1
	}

	currentDate := time.Now().Format("02.01.2006")

	newTask := Task{
		ID:        newID,
		Title:     title,
		Done:      false,
		CreatedAt: currentDate,
	}

	return append(tasks, newTask)
}

// readLine - читает строку с консоли
func readLine() string {
	var input string
	fmt.Scanln(&input)
	return input
}

// deleteTask - Удаляет задачу
func deleteTask(tasks []Task, id int) []Task {
	// 1. Ищем задачу
	for i, task := range tasks {
		if task.ID == id {
			// 2. Удаляем задачу (самый простой способ)
			// tasks[:i] - все элементы до i
			// tasks[i+1:]... - все элементы после i
			newTasks := append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Задача '%s' удалена\n", task.Title)
			return newTasks
		}
	}

	// 3. Если не нашли
	fmt.Printf("Задача с ID %d не найдена\n", id)
	return tasks
}

func editTask(tasks []Task, id int, newTitle string) []Task {
	// 1. Ищем задачу по ID
	for i := range tasks {
		if tasks[i].ID == id {
			if newTitle == "" {
				fmt.Println("Ошибка: новое название не может быть пустым!")
				return tasks
			}

			// 2. Меняем название задачи
			oldTitle := tasks[i].Title
			tasks[i].Title = newTitle

			// 3. Сообщаем пользователю
			fmt.Printf("Задача '%s' переименована в '%s'\n", oldTitle, newTitle)
			return tasks
		}
	}

	// Если дошли сюда - задача не найдена
	fmt.Printf("Задача с ID %d не найдена\n", id)
	return tasks
}

// showTaskByID - Показать задачу по ID
func showTaskByID(tasks []Task) {
	fmt.Print("Введите ID задачи для просмотра: ")
	var id int
	fmt.Scan(&id)

	taskPtr := getTaskByID(tasks, id)

	if taskPtr == nil {
		fmt.Println("Задача не найдена")
	} else {
		status := "Не выполнена"
		if taskPtr.Done {
			status = "Выполнена"
		}
		fmt.Printf("Задача #%d: %s (%s)\n", taskPtr.ID, taskPtr.Title, status)
	}
}

// showStats - Показывает статистику задач (Всего, выполненные, невыполненные, прогресс)
func showStats(tasks []Task) {
	total := len(tasks)
	completed := 0

	for _, task := range tasks {
		if task.Done {
			completed++
		}
	}

	fmt.Println("\n=== СТАТИСТИКА ===")
	fmt.Printf("Всего задач: %d\n", total)
	fmt.Printf("Выполнено: %d\n", completed)
	fmt.Printf("Осталось: %d\n", total-completed)

	if total > 0 {
		percent := float64(completed) / float64(total) * 100
		fmt.Printf("Прогресс: %.1f%%\n", percent)
	}
}

// getTaskByID - Возвращает *Task - указатель на Task
func getTaskByID(tasks []Task, id int) *Task {
	// Ищем задачу
	for i := range tasks {
		if tasks[i].ID == id {
			// Возвращаем АДРЕС задачи в памяти, а не её копию
			// & - оператор "взять адрес"
			return &tasks[i]
		}
	}

	// Задача не найдена - возвращаем nil (особое значение "ничего" для указателей)
	return nil
}

// saveTasksToFile - сохраняет задачи в файл (простой текст)
func saveTasksToFile(tasks []Task, filename string) {
	// 1. Открываем файл для записи
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close() // Закроем файл при выходе из функции

	// 2. Записываем задачи в файл
	for _, task := range tasks {
		status := "Не выполнена"
		if task.Done {
			status = "Выполнена"
		}
		line := fmt.Sprintf("ID:%d | %s | %s | Создано: %s\n",
			task.ID, task.Title, status, task.CreatedAt)
		file.WriteString(line)
	}

	fmt.Printf("Задачи сохранены в файл: %s\n", filename)
}

// loadTasksSimple - Загружает задачи
func loadTasksSimple(filename string) []Task {
	file, err := os.Open(filename)
	if err != nil {
		return []Task{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tasks := []Task{}
	id := 1

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Создаём задачу, где вся строка - это название
		task := Task{
			ID:        id,
			Title:     line, // Вся строка как название
			Done:      false,
			CreatedAt: time.Now().Format("02.01.2006"),
		}

		tasks = append(tasks, task)
		id++
	}

	return tasks
}

// showMenu - Показывает меню приложения
func showMenu() {
	fmt.Println("\n=== TO-DO МЕНЕДЖЕР ===")
	fmt.Println("1. Показать все задачи")
	fmt.Println("2. Добавить задачу")
	fmt.Println("3. Редактировать задачу")
	fmt.Println("4. Отметить задачу выполненной")
	fmt.Println("5. Отметить задачу НЕ выполненной")
	fmt.Println("6. Удалить задачу")
	fmt.Println("7. Показать статистику")
	fmt.Println("8. Показать задачу по ID")
	fmt.Println("9. Сохранить задачи в файл")
	fmt.Println("10. Загрузить задачи из файла")
	fmt.Println("0. Выход")
	fmt.Print("Выберите действие: ")
}

func main() {

	tasks := []Task{
		{ID: 1, Title: "Сходить в магазин", Done: false, CreatedAt: "17.01.2026"},
		{ID: 2, Title: "Купить яйца", Done: false, CreatedAt: "17.01.2026"},
		{ID: 3, Title: "Сходить на OZON", Done: true, CreatedAt: "17.01.2026"},
	}

	for {
		showMenu()

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			printTasks(tasks)

		case 2:
			fmt.Print("Введите название задачи: ")
			title := readLine() // Теперь читаем так

			// Добавляем валидацию
			if title == "" {
				fmt.Println("Ошибка: название не может быть пустым!")
			} else {
				tasks = addTask(tasks, title)
				fmt.Println("Задача добавлена!")
			}
		case 3:
			printTasks(tasks)
			fmt.Print("Введите ID задачи для изменения: ")
			var id int
			fmt.Scan(&id)
			fmt.Print("Введите новое название: ")
			var newTitle string
			fmt.Scanln() // Очищаем буфер
			fmt.Scanln(&newTitle)
			tasks = editTask(tasks, id, newTitle)
		case 4:
			printTasks(tasks)
			fmt.Print("Введите ID задачи для отметки: ")
			var id int
			fmt.Scan(&id)
			completeTask(tasks, id)
		case 5:
			printTasks(tasks)
			fmt.Print("Введите ID задачи для отметки как НЕ выполненной: ")
			var id int
			fmt.Scan(&id)
			uncompleteTask(tasks, id)
		case 6:
			printTasks(tasks)
			fmt.Print("Введите ID задачи для удаления: ")
			var id int
			fmt.Scan(&id)
			tasks = deleteTask(tasks, id)

		case 7:
			showStats(tasks)
		case 8:
			showTaskByID(tasks)
		case 9:
			saveTasksToFile(tasks, "tasks.txt")
		case 10:
			loadedTasks := loadTasksSimple("tasks.txt")
			if len(loadedTasks) == 0 {
				fmt.Println("Не удалось загрузить задачи или файл пуст")
			} else {
				tasks = loadedTasks
				fmt.Printf("Загружено %d задач\n", len(tasks))
			}
		case 0:
			// Спросим, хочет ли сохранить перед выходом
			fmt.Print("Сохранить задачи в файл перед выходом? (y/n): ")
			answer := readLine()
			if answer == "y" || answer == "Y" {
				saveTasksToFile(tasks, "tasks.txt")
			}
			fmt.Println("Выход из программы")
			return
		default:
			fmt.Println("Неверный выбор")
		}
	}
}
