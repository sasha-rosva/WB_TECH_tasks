package main

import (
	"fmt"
	"time"
)
const n=5 // Количество секунд, по истечению которых программа завершится
func main(){
	// Создадим канал
	ch1:=make(chan int)
	// Создадим горутину, которая будет читать из канала
	go func(ch <-chan int){
		for v:= range ch{
			time.Sleep(500*time.Millisecond)
			fmt.Println(v)
		}
	}(ch1)
	// Создадим горутину, которая будет писать в канал
	go func(ch chan<- int){for i:=0;;i++{
		ch<-i
	}}(ch1)
	// Установим время на таймере
	timer1 := time.NewTimer(time.Second * n)
	// Ждем выполнения таймера
	<-timer1.C
	// Закрываем канал, тем самым звавершая программу
	close(ch1)

}
