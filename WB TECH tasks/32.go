package main

import (
	"fmt"
	"time"
)

func After(i int,ch chan bool){
	now:=time.Now().Second()
	if 60-time.Now().Second()>i{
	for {

		if time.Now().Second()-now==i{
			break
		} }

		}else{for{
		if time.Now().Second()==now+i-60{
			break
		}
	}



	}
ch<-true
}
func MySleep(i int){
	ch1:=make(chan bool)
	go After(i,ch1)
	<-ch1
}
func main(){
	sec:=5
	//fmt.Println(time.Now().Second())
	fmt.Printf("будем спать %v секунд!\n",sec)
	MySleep(sec)
	fmt.Println("Кол-во секунд, при которых программа была в спячке:",sec)
	//fmt.Println(time.Now().Second())
}
