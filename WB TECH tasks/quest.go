package main

import (
	"fmt"
	"strings"
	"sync"
)

func main(){
	mu:=sync.RWMutex{}
	map1:=make(map[rune]int)
	str:="AAABBCCCDDEEE"
	runes := []rune(str)
	for _,v:=range runes{
		go func(v rune){
			mu.Lock() // Блокировка
			map1[v]++
			mu.Unlock() // Отключение блокировки
		}(v)
	}
	var ourStr string
	sb:=new(strings.Builder)
	for ii,vv:= range map1 {
		ourStr=fmt.Sprintf("%d%s",vv,string(ii))
		sb.WriteString(ourStr)
	}

fmt.Println(sb.String())
}
