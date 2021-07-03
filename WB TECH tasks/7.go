package main
//Конкурентная запись в map
import (
	"fmt"
	"sync"
)

func main(){
	// Map является конкурентно не безопасным типом данных,
	// поэтому воспользуемся sync.Mutex для надежного обращения с map.
	mu:=&sync.Mutex{}
	map1:=make(map[int]int)
	mas:=[...]int{0,1,2,3,4,5,6,7,8,9}
	for i,v:=range mas{
		go func(i,v int){
			mu.Lock() // Блокировка
			map1[i]=v
			mu.Unlock() // Отключение блокировки
		}(i,v)
	}
	mu.Lock()
fmt.Println(map1[2])
	mu.Unlock()
}
