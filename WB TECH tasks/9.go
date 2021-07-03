package main
// Реализация конвейера чисел
import (
	"fmt"
)

func main(){
	// Создаем 2 канала и массив
	ch1:=make(chan int)
	ch2:=make(chan int)
	mas:=[...]int{0,1,2,3,4,5,6,7,8,9}
	// Запускаем первую горутину для итерирования массива с последующей отправкой в канал
	go func(ch chan<- int){
		for _,value:=range mas{
			ch<-value
		}
		close(ch1)  // закрываем первый канал
	}(ch1)
	// Запускаем вторую горутину для получения значений из 1-го канала с последующей отправкой во 2-й канал
	go func(ch chan<- int){
		for v1:=range ch1{
			ch<-v1*2
		}

		close(ch2) // закрываем второй канал
	}(ch2)
	// Выводим в stdout значения из второго канала
	for v2:=range ch2{
		fmt.Println(v2)
	}
}
