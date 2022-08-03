package main

import (
	"fmt"
	"log"
)

/*
Удалить i-ый элемент из слайса.
*/

func deleteIthElement(arr []int, i int) []int {
	if i < 0 {
		log.Fatal("Please input positive index\n")
	}
	if i > len(arr)-1 {
		log.Fatal("Please input correct index!\n")
	}
	// воспользуемся функцией append для конкатенации слайсов
	return append(arr[:i], arr[i+1:]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(deleteIthElement(arr, 0))
}
