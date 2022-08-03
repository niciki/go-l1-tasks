package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

// Создаем set строк, далее возвращаем слайс уникальных строк с помощью построенного map
func MakeSet(arr []string) []string {
	data := make(map[string]bool)
	for _, i := range arr {
		data[i] = true
	}
	resSet := make([]string, 0, len(data))
	for key := range data {
		resSet = append(resSet, key)
	}
	return resSet
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(MakeSet(arr))
}
