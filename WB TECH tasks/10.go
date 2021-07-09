package main

import "fmt"

func main (){
	ourMap:=make(map[int] []float32)
	mas:=[...]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -13.8}
	var intValue int
	for _,v:= range mas{
		intValue=int(v/10)
		ourMap[10*intValue]=append(ourMap[10*intValue],v)
	}
	fmt.Println(ourMap)

}



