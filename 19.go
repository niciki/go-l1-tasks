package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func reverse(str string) string {
	resString := ""
	// Преобразуем строку для корректного итерирования по символам строки в слайс рун
	strArray := []rune(str)
	//обходим слайс рун в обратном порядке, записываем их в результирующую строку
	for i := len(strArray) - 1; i >= 0; i-- {
		resString += string(strArray[i])
	}
	return resString
}

func main() {
	fmt.Print(reverse("главрыба"))
}
