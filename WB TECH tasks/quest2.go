package main

import (
	"fmt"
	"strings"
)

func main(){
	map1:=make(map[rune]int)
	var ourStr string
	var d rune
	sb:=new(strings.Builder)
	str:="AAABBCCCDDEEE"
	runes := []rune(str)
	for _,v:=range runes{
		func(v rune){
			if v!=d && map1[d]!=0{
				ourStr=fmt.Sprintf("%d%s",map1[d],string(d))
				sb.WriteString(ourStr)
			}
			map1[v]++
			d=v
		}(v)
	}
	sb.WriteString(fmt.Sprintf("%d%s",map1[runes[len(runes)-1]],string(d)))
	fmt.Println(sb.String())
}
