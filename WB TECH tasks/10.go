package main

import "fmt"

func main (){
	ourMap:=make(map[int] []float32)
	mas:=[...]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for _,v:= range mas{
		switch {
		case v<=-20 && v>-30:
			ourMap[-20]=append(ourMap[-20],v)
		case v<=-10 && v>-20:
			ourMap[-10]=append(ourMap[-10],v)
		case v<=0 && v>-10:
			ourMap[0]=append(ourMap[0],v)
		case v<=10 && v>0:
			ourMap[10]=append(ourMap[10],v)
		case v<=20 && v>10:
			ourMap[20]=append(ourMap[20],v)
		case v<=30 && v>20:
			ourMap[30]=append(ourMap[30],v)
		case v<=40 && v>30:
			ourMap[40]=append(ourMap[40],v)
		}
		}
		fmt.Println(ourMap)

	}



