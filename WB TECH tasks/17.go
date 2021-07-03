package main

import (
	"fmt"
	"reflect"
)

func main(){
	// Инициализируем переменную любого типа
	variable:=make(chan float64)
	// Воспользуемся пакетом "reflect"
	xType := reflect.TypeOf(variable)
	// Выведем тип переменной в stdout
	fmt.Println(xType)
}
