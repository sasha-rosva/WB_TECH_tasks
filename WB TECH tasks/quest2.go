package main

import (
	"fmt"
	"sync"
)

func Number(s string) int {
	runes := []rune(s)
	map1:=make(map[rune]bool)
	for _,run:=range runes{
		map1[run]=true
	}
	return len(map1)
}

func main(){
	wg:=&sync.WaitGroup{}
	mas:=[...] string{"abcdefghijklmnopqrstuvwxyz","abcdefghijklmnopqrstuvwxy","vccc"}
for i,v:= range mas{
	wg.Add(1)
	go func(i int,v string){
		defer wg.Done()
		num:=Number(v)
	fmt.Printf("Количество уникальных символов у строки на позиции %d: %d\n",i,num)
}(i,v)}
wg.Wait()
}
