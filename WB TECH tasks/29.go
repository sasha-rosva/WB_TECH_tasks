package main

import "fmt"

func Sum (a int64,b int64){
	fmt.Println("Сумма чисел равна: ",a+b)
}
func Prod (a int64,b int64){
	fmt.Println("Произведение чисел равно: ",a*b)
}
func Diff (a int64,b int64){
	fmt.Println("Разность чисел равна: ",a-b)
}
func Div (a float64,b float64){
	fmt.Println("Частное чисел равно: ",a/b)
}
func main(){
	var a,b int64
	a,b=345367798,453563453
	Sum(a,b)
	Prod(a,b)
	Diff(a,b)
	var aa,bb float64
	aa,bb=34536778,453563453
	Div(aa,bb)
}