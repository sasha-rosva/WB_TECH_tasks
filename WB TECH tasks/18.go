package main

import "fmt"

func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}
func main() {
	// Инициализируем срез
	var a = []int8{1, 2, 3, 4, 5}
	/*Выполним функцию someAction. Т.к. срез является модифицированным указателем на массив,
	первая операция в функции изменит первое значение в массиве на который ссылается слайс
	Функция append возвращает обновленный слайс, который записывается в переменную с уже
	локальной видимостью. */
	someAction(a, 6)
    // Поэтому в stdout выведется [100,2,3,4,5]
	fmt.Println(a)
}