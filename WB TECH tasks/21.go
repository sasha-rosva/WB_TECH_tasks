package main

import (
	"fmt"
	"sync"
)

// Функция будет печатать данные в stdout
func printer( i interface{},wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
}

func main() {
	// Воспользуемся sync.WaitGroup для ожидания завершения нескольких goroutine
	wg:=&sync.WaitGroup{}

	// Создадим массив с данными
	months := [...]string{"Январь", "Февраль", "Март",
		"Апрель", "Май", "Июнь",
		"Июль", "Август", "Сентябрь",
		"Октябрь", "Ноябрь", "Декабрь",
	}
	// Итерируемся по массиву
	for _, monthName := range months {
		wg.Add(1)
		go printer(monthName,wg)
	}
	// Ждем  выполнения работы всех горутин
	wg.Wait()
}
