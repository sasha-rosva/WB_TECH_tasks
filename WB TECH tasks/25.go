package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)
type Counter struct {
	Total int32
}
// Увеличивает счетчик на 1
func (c *Counter)incRequests() {
	atomic.AddInt32(&c.Total, 1)
}

// Выводит в stdout текущее значение счетчика
func (c *Counter)getRequests()  int32{
	val:=atomic.LoadInt32(&c.Total)
	fmt.Println(val)
	return val
}

func main() {
	wg:=&sync.WaitGroup{}

	counter := new(Counter)
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(c *Counter) {
			defer wg.Done()
			counter.incRequests()
			counter.getRequests()
			runtime.Gosched()
		}(counter)
		//time.Sleep(time.Second)
	}
	wg.Wait()

}
