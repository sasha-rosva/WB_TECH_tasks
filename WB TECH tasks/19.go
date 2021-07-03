package main

import "fmt"

func createHugeString(size int) string{
	var v string
	for i:=0;i<size;i++{
		v+="A"
	}
	return v
}

var justString string
func someFunc() {
	// Получаем длинную строку с помощью функции createHugeString. Предполагая. что мы не знаем, что будет записано
	// в этой строке можем получить не ожидаемое нами поведение программы. А именно: возможно в нашей строке
	// будет содержаться символ, занимающий больше одного байта. Строка в Go - это слайс байт, соответственно можем
	// потерять часть символов. Лучше использовать руны.
	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100])
	fmt.Println(justString)
}
func main() {
	someFunc()
}
