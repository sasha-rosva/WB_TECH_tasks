package main
//Конкурентная запись в map
import (
	"fmt"
	"sync"
)

func main(){
	// Map является конкурентно не безопасным типом данных,
	// поэтому воспользуемся sync.Mutex для надежного обращения с map.
	mu:=sync.RWMutex{}
	map1:=make(map[int]int)
	mas:=[...]int{0,1,2,3,4,5,6,7,8,9}
	n:=len(mas)
	for i,v:=range mas{
		go func(i,v int){
			mu.Lock() // Блокировка
			map1[i]=v
			mu.Unlock() // Отключение блокировки
		}(i,v)
	}
	// Ждем оканчания записи в map1
	for{
		if len(map1)==n{break}
	}
	mu.RLock()
fmt.Println(map1[3])
	mu.RUnlock()
}
