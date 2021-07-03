package main

import "fmt"

func Unique(s string)  {
	runes := []rune(s)
	map1:=make(map[rune]int)
	for _,run:=range runes{
		if map1[run]>0{
			fmt.Println("Символы в строке не уникальны!")
			return
		}
		map1[run]++
	}
	fmt.Println("Символы в строке уникальны!")
	return  
}

func main(){
	str:="abcdefghijklmnopqrstuvwxyz"
	str1:="abcdefghijkldmnopqrstuvwxyz"
	Unique(str)
	Unique(str1)
}
