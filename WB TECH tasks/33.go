package main
// Реализация конвейера чисел
import (
"fmt"
"math/rand"
"time"
)

func main(){
	// Устанавливаем seed
	rand.Seed(time.Now().UnixNano())
	// Создаем 2 канала и массив
	ch1:=make(chan int)
	ch2:=make(chan string)
	// Запускаем первую горутину для итерирования массива с последующей отправкой в канал
	go func(ch chan<- int){
		for i:=0;i<10;i++{
			ch<-rand.Intn(100)
		}
		close(ch1)  // закрываем первый канал
	}(ch1)
	// Запускаем вторую горутину для получения значений из 1-го канала с последующей отправкой во 2-й канал
	go func(ch chan<- string){
		for v1:=range ch1{
			if v1%2==0 {
				ch <-fmt.Sprintf("%d - четное число!",v1)
			} else {
				ch <-fmt.Sprintf("%d - нечетное число!",v1)
			}
		}

		close(ch2) // закрываем второй канал
	}(ch2)
	// Выводим в stdout значения из второго канала
	for v2:=range ch2{
		fmt.Println(v2)
	}
}
