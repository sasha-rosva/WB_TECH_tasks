package main
// Пример №2. Воспользуемся пакетом context.
import (
	"context"
	"fmt"
	"time"
)
// Функция worker, запущенная в горутине, закрывается при получения сигнала от контекста
func worker(ctx context.Context, num int,out chan<- int){
	select{
	case <-ctx.Done(): // ждем сигнал о закрытии
		return
		case <-time.After(3*time.Second):
			out<-num
	}
}

func main(){
	ch1:=make(chan int,1)
	ctx,finish:=context.WithCancel(context.Background())
for i:=0;i<10;i++{
	go worker(ctx,i,ch1)
}

fmt.Println(<-ch1)
// Отправляем сигнал в контекст о закрытии
finish()
}
