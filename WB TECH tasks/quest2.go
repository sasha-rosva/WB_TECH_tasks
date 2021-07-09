package main

import "fmt"

func Number(s string) int {
	runes := []rune(s)
	map1:=make(map[rune]bool)
	for _,run:=range runes{
		map1[run]=true
	}
	return len(map1)
}

func main(){
	mas:=[...] string{"abcdefghijklmnopqrstuvwxyz","abcdefghijklmnopqrstuvwxy","vccc"}
for i,v:= range mas{
	num:=Number(v)
	fmt.Printf("Количество уникальных символов у строки на позиции %d: %d\n",i,num)
}
}
