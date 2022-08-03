package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func SetIntersection(set1, set2 []int) []int {
	set3 := []int{}
	data := make(map[int]bool, len(set1))
	// Строим мапу за О(n)
	for _, i := range set1 {
		data[i] = true
	}
	// Обходим массив на поиск одинаковых элементов за O(n)
	for _, i := range set2 {
		if data[i] {
			set3 = append(set3, i)
		}
	}
	return set3
}

func main() {
	set1 := []int{13, 10, 8, 12, 6, 4, 15, 13, 5, 3, 7, 2, 9, 1}
	set2 := []int{15, 12, 2, 20, 1, 21}
	fmt.Println(SetIntersection(set1, set2))
}
