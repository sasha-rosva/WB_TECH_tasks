package main

import "fmt"

// Object Адаптер нужен для того, чтобы наделить объект определенным интерфейсом
// Создаем структуру
type Object struct{
	Str int
}

// Update Экземпляры этой структуры могут обновлять свое поле новым значением
func (o *Object) Update(i int){
	o.Str=i
}
// Но, к сожалению, не могут показать значения своего поля, как это делают другие структуры
type Object1 struct{
	Str int
}
func (o *Object1) Show(){
	fmt.Println(o.Str)
}
func (o *Object1) Update(i int){
	o.Str=i
}
// В эту функцию могут попасть все, кто удовлетворяют интерфейсу Shower
func ShowToUs(s Shower){
	s.Show()
}
// Интерфейс Shower
type Shower interface{
	Show()
}
// Т.к. наша структура Object не удовлетворяет данному интерфейсу, нам понадобится адаптер.
type Adapter struct{
	Obj *Object
}
func (a Adapter) Show(){
	fmt.Println(a.Obj.Str)
}

func main(){
	object:=new( Object)
	object.Update(77)
	object1:=new(Object1)
	object1.Update(88)
	var adapter=Adapter{Obj: object}
	ShowToUs(object1)
	ShowToUs(adapter)


}


