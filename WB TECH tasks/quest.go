package main

import (
	"fmt"
	"strings"
)

func Join( words ...string) (string,string){
	sb:=new(strings.Builder)
	for _,v:= range words{
		sb.WriteString(v)
	}
	a1:=sb.String()
	sb.Reset()
	for i,v:= range words{
		if i%2==0{
		sb.WriteString(v)
	}}

	return a1,sb.String()
}

func main(){
w1:="London"
w2:=" capital"
q1,q2:=Join(w1," is"," a",w2," of the Great Britain!")
fmt.Println(q1)
fmt.Println(q2)
}
