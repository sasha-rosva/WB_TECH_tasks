package main

import (
	"fmt"
	"strings"
	"sync"
)

func main(){
	mu:=sync.RWMutex{}
	wg:=&sync.WaitGroup{}
	map1:=make(map[rune]int)
	str:="AAABBCCCDDEEE"
	runes := []rune(str)
	for _,v:=range runes{
		wg.Add(1)
		go func(v rune){
			defer wg.Done()
			mu.Lock() // Блокировка
			map1[v]++
			mu.Unlock() // Отключение блокировки
		}(v)
	}
	wg.Wait()
	var ourStr string
	sb:=new(strings.Builder)
	for ii,vv:= range map1 {
		ourStr=fmt.Sprintf("%d%s",vv,string(ii))
		sb.WriteString(ourStr)
	}

fmt.Println(sb.String())
}
