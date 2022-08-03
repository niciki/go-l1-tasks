package main

import "fmt"

/*
	Реализовать бинарный поиск встроенными методами языка.
*/

func binarySearch(arr []int, key int) bool {
	l := 0
	r := len(arr) - 1
	for l < r {
		m := (l + r) / 2
		if arr[m] < key {
			l = m + 1
		} else if arr[m] > key {
			r = m - 1
		} else {
			return true
		}
	}
	if arr[l] == key {
		return true
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	fmt.Println(binarySearch(arr, 1), " ", binarySearch(arr, 16), " ", binarySearch(arr, 25), " ")
	fmt.Println(binarySearch(arr, 7))
}
