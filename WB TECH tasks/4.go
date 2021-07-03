package main

import (
	"fmt"
	"sync"
)

// Воркер будет печатать данные из канала в stdout
func worker(id int, jobs <-chan interface{},wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "данные из канала: ", j)
	}
}

const n = 19 // Устанавливаем нужное количество воркеров
func main() {
	// Воспользуемся sync.WaitGroup для ожидания завершения нескольких goroutine
	wg:=&sync.WaitGroup{}
	// Создадим канал, который будет использовать значения типа interface{}
	jobs := make(chan interface{}, 100)
	// Запускаем  нужное количество воркеров
	for w := 0; w < n; w++ {
		wg.Add(1)
		go worker(w, jobs,wg)
	}
	// Отправляем данные в канал
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	// Создадим срез с данными
	months := []string{"Январь", "Февраль", "Март",
		"Апрель", "Май", "Июнь",
		"Июль", "Август", "Сентябрь",
		"Октябрь", "Ноябрь", "Декабрь",
	}
	//Отправляем другие данные в канал
	for _, monthName := range months {
		jobs <- monthName
	}

	// Закрываем канал, чтобы показать, что больше задач нет.
	close(jobs)
	// Ждем  выполнения работы всех горутин
	wg.Wait()
}
