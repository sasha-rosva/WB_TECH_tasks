package main

import "fmt"

func main() {
	slice := []string{"a", "a"}
	func(slice []string) {
		slice = append(slice, "a") // Функция append возвращает обновленный слайс, поэтому slice будет в локальной области видимости
		slice[0] = "b"
		slice[1] = "b"
		fmt.Println(slice)
	}(slice)
	fmt.Println(slice)
}
