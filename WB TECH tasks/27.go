package main

import "fmt"

func Reverse(s string) string {
	r:=[]rune(" ")
	u:=r[0]
	run:=[]rune("")
	run2:=[]rune("")
	runes := []rune(s)
	j2:=len(runes)
	for j:=len(runes)-1;j>0;j--{
		if t:=runes[j]; t==u{
			run=append(run,runes[j+1:j2]...)
			run=append(run,u)
			j2=j
		}
	}
	for i:=0;i<len(runes);i++{
		if t2:=runes[i]; t2==u{
			run2=append(run2,runes[:i+1]...)
			break
	}}
	run=append(run,run2...)
	return string(run)
}


func main (){
	d:="Britain Great The Of Capital A Is London"
	dd:=Reverse(d)
	fmt.Println(dd)
}
