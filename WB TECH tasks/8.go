package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main(){
	// случайно генерируем значение бита (0 или 1)
	rand.Seed(time.Now().UTC().UnixNano())
    needNum:=rand.Intn(10)%2
    fmt.Println("Значение бита, полученное рандомно: ",needNum)
    // указываем номер бита, который мы поменяем
	n:=0
	// указываем число типа int64
	num := int64(-1234)

	// получаем число в битовом представлении в виде строки
    s:=strconv.FormatInt(num, 2)

    // меняем n-й бит на полученный рандомно
    news:=s[:n] + strconv.Itoa(needNum) + s[n+1:]
    // получаем итоговое число
    b,_:=strconv.ParseInt(news,2,64)
    fmt.Println("Изначальное число и число, полученное в результате возможного изменения n-ого бита: ",num, " : ",b)
}
