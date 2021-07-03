package main

import "fmt"

func main(){
// Создаем слайс с емкостью 2. Это означает, что в слайс можно добавить 2 значения без аллоцирования новой памяти.
	// При переполнения емкости, срез будет ссылаться на бОльший массив с копированием данных в него
	slice:=make([]int,0,2)

	slice=append(slice,88)
	fmt.Println(slice,cap(slice))
	slice=append(slice,88)
	fmt.Println(slice,cap(slice))
	slice=append(slice,88)
	fmt.Println(slice,cap(slice))
	slice=append(slice,88)
	fmt.Println(slice,cap(slice))
	slice=append(slice,88)
	fmt.Println(slice,cap(slice))
}
