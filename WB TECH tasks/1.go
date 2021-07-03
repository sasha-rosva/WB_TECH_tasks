package main
/*В данном примере с помощью композиции и встраивания
создадим Humans, поддерживающие свои любимые команды!*/
import "fmt"

// Action Cтруктура Action   хранит в себе название любимой команды!
type Action struct{
	Team string
}

// Shout Создаем метод Shout для структуры Action
func (a *Action) Shout(){fmt.Printf("%s - чемпион!!!\n",a.Team)}
// Sing Создаем метод Sing для структуры Action
func (a *Action) Sing(){fmt.Printf("Оле-оле-оле-оле-е-е-е! %s - чемпион!!!\n",a.Team)}
// Human Определяем  структуру Human, которая является композицией!
type Human struct{
	Name string    // Имя
	Age uint8      // Возраст
	Weight uint16  // Вес
	Height uint8   // Рост
	Action         // Анонимное поле Action
}

func main(){
	//Имплементируем структуру Human, т.е. создаем первого болельщика!
	human1:=Human{Age: 20,Weight: 71,Height: 170, Action: Action{Team: "Спартак"}}

	// Воспользуемся методом Sing
    human1.Action.Sing() //output: Оле-оле-оле-оле-е-е-е! Спартак - чемпион!!!

    // С помощью встаивания вызов метода можно упростить!
	human1.Shout()       //output: Спартак - чемпион!!!


	//Создаем второго болельщика!
	action:=Action{"Зенит"}
	human2:=Human{Age: 20,Weight: 70,Height: 170, Action: action}

	// Аналогично работает и для второго болельщика!
	human2.Sing()        //output: Оле-оле-оле-оле-е-е-е! Зенит - чемпион!!!
	human2.Shout()       //output: Зенит - чемпион!!!
}
