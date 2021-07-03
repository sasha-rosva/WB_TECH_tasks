// Пример №1. Горутины постоянно опрашивают канал отмены. Мы посылаем значение в канал отмены, тем самым сигнализируем горутинам закрываться.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func Print(wg *sync.WaitGroup,i int,cancel chan bool){
	defer wg.Done()
	for{select {
	case <-cancel:
		fmt.Println("Горутина закрылась!")
		return
	default:
		fmt.Println(i)
		runtime.Gosched()
	}}
}

func main(){
	wg:=&sync.WaitGroup{}
	cancel:=make(chan bool)
	for i:=0;i<10;i++{
		wg.Add(1)
		go Print(wg,i,cancel)}

	time.Sleep(2*time.Second)
	for t:=0;t<10;t++{
		cancel<-true}
	wg.Wait()
	fmt.Println("Game over!")
}