package main

import (
	"fmt"
)

func QuickSort(v []int){
	// Напишем базовое условие рекурсии
	if len(v) <= 1 {
		return
	}
	split := Partition(v)

	QuickSort(v[:split])
	QuickSort(v[split:])}
 // Partition Функция итерируется по срезу и сравнивает значения с опорным значением. После первого полного прохода
/* функция возвращает индекс, с помощью которого массив будет разделен на 2 части и с помощью рекурсии
 в них будет продолжена сортировка.*/
func Partition(v []int) int {
// Для алгоритма быстрой сортировки нужен опорный элемент. Обычно выбирают где-то посередине.
pivot:=len(v)/2
// Индексы с левой и с правой стороны
l:=0
r:=len(v)-1
// Значение опорного элемента
value:=v[pivot]
for{
	for ; v[l]<value; l++{}
	for ; v[r]>value; r--{}
	if l >= r {
		return r}
	change(v, l, r)
}
}
// Функция для перемещения значений в ситуациях, при которых значение на которое указывает левый индекс больше
// опорного значения, а правого меньше
func change(v []int, l, r int) {
	tmp := v[l]
	v[l] = v[r]
	v[r] = tmp
}

func SearchInIntSlice(haystack []int, needle int) (result bool, iterationsCount int) {
	lowKey := 0 // первый индекс
	highKey := len(haystack) - 1 // последний индекс
	if (haystack[lowKey] > needle) || (haystack[highKey] < needle) {
		return // нужное значение не в диапазоне данных
	}
	for lowKey <= highKey {
		// уменьшаем список рекурсивно
		iterationsCount++
		mid := (lowKey + highKey) / 2 // середина
		if haystack[mid] == needle {
			result = true // мы нашли значение
			return
		}
		if haystack[mid] < needle {
			// если поиск больше середины - мы берем только блок с большими значениями увеличивая lowKey
			lowKey = mid + 1
			continue
		}
		if haystack[mid] > needle {
			// если поиск меньше середины - мы берем блок с меньшими значениями уменьшая highKey
			highKey = mid - 1
		}
	}
	return
}
func main (){
	// Инициализируем среза
	arr:=[]int{6,4,5,8,2,90,55,45,44,89,34,77}
	// Для бинарного поиска нам нужен отсортированный срез.
	QuickSort(arr)
	fmt.Println(arr)
	result,n:=SearchInIntSlice(arr,46)
	fmt.Println(result,n)
}
