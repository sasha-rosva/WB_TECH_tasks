package main

import "fmt"
// В данном примере мы найдем пересечение двух неупорядоченных массивов
func main(){
	// Создаем 2 массива, срез для хранения пересекающихся значений
	mas1:=[...]int{1,2,4,5,6,6,8,9,3,3,2,10,3}
	mas2:=[...]int{2,2,6,6,7,12,17,5,3,3,3,3,3,3,3}
	var intersection []int
	// Создаем map. В качестве ключа используем значение, а частота появления в первом массиве в качестве значения
	map1:=make(map[int]int)
	for _,v1:= range mas1{
		map1[v1]++
	}
	// Ищем пересечение и при обнаружении значения уменьшаем значение в map-е
	for _, v2:= range mas2{
		if _,ok:=map1[v2]; ok && map1[v2]>0{
				intersection=append(intersection,v2)
			map1[v2]--
			}
		}

fmt.Println(intersection)
}
