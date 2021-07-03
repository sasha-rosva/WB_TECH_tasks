package main

import "fmt"

func DeleteAnyValue(val []int, i int) []int {
	if i>len(val)-1 {fmt.Println("В данном срезе нет такого индекса!")
	return nil}
	if len(val)==i+1{
		val=val[:i]
		return val
	}
	val=append(val[:i],val[i+1:]...)
	return val
}

func main(){
	// Инициализируем срез
	val:=[]int{1,2,3,4,5,6,7,8}
	// Выбираем i-й элемент из слайса.
	n:=4
	newVal:=DeleteAnyValue(val,n)
	// Выводим результат
	if newVal!=nil{fmt.Println(newVal)}

}