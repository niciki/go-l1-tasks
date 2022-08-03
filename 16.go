package main

import (
	"fmt"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

// берем самый правый элемент массива, все элементы которые меньше его, ставим в начало массива,
// таким образом, все, что больше, окажутся справа от переменной m. Дальше элемент, относительно
// которого мы распределяли ставим на позицию m
func Partition(arr []int, low, high int) int {
	pivot := arr[high]
	m := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[m], arr[j] = arr[j], arr[m]
			m++
		}
	}
	arr[m], arr[high] = arr[high], arr[m]
	return m
}

// пока левая и правая границы сортируемых массивов не совпадают, продолжаем сортировку.

func quickSort(arr []int, low, high int) {
	if low < high {
		p := Partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func main() {
	arr := []int{13, 10, 8, 12, 6, 4, 15, 13, 5, 3, 7, 2, 9, 1}
	fmt.Printf("Before sorting: %v\n", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Printf("After sorting: %v\n", arr)
}
