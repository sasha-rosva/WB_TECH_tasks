package main

import "fmt"



func main (){
	map1:=make(map[string]bool)
	var set []string
	arr:=[]string{"cat", "cat", "dog", "cat", "tree"}
	for _, v:= range arr{
		map1[v]=true
	}
	for k:= range map1{
		set=append(set,k)
	}
fmt.Println(set)
}