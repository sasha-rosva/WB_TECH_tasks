package main

import (
	"fmt"
	"sync"
)

func main() {
	// Структура WaitGroup имеет поле noCopy, которое дает понять, что данную структуру не стоит копировать.
	// А при передачи аргументов в функцию они как раз и копируются, что недопустимо.
	//Структура WaitGroup должна создаваться по указателю
	wg := sync.WaitGroup{} // т.е. wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) { // и go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
