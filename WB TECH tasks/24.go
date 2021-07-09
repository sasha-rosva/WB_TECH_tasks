package main

import "fmt"

func main(){
	slice:=make([]int,100)
	
	fmt.Printf("Наш слайс: %v\nЁмкость нашего слайса: %v",slice,cap(slice))
}
