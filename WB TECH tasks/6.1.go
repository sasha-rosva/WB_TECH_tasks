package main
// Пример №2. Воспользуемся пакетом context.
import (
	"context"
	"fmt"
	"runtime"
	"time"
)
// Функция worker, запущенная в горутине, закрывается при получения сигнала от контекста
func worker(ctx context.Context, num int){
	for {select{
	case <-ctx.Done(): // ждем сигнал о закрытии
		fmt.Println("Worker ",num," закрылся!")
		return
	case <-time.After(time.Second):
		fmt.Println("Worker ",num)
	}
	}}

func main(){
	ctx,finish:=context.WithCancel(context.Background())
	for i:=0;i<10;i++{
		go worker(ctx,i)
	}

	time.Sleep(3*time.Second)

	// Отправляем сигнал в контекст о закрытии
	finish()
	// Ждем пока все горутины не закроются
	for {
		if runtime.NumGoroutine()==1{break}
	}
	fmt.Println("Все горутины успешно закрылись!")
}
