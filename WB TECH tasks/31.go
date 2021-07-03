package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}
func NewPoint(x,y float64) *Point{
	return &Point{X:x,Y:y}
}
func EuclidDistance(p1,p2 *Point){
	fmt.Println(math.Sqrt(math.Pow(p1.X-p2.X,2)+math.Pow(p1.Y-p2.Y,2)))
}

func main(){
	point1:=NewPoint(5,-3)
	point2:=NewPoint(4,9)
	EuclidDistance(point1,point2)

}