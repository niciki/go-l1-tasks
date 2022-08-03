package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func reverseWords(str string) string {
	// получаем слайс строк из исходной строки по пробелу
	strArray := strings.Split(str, " ")
	resString := ""
	// пробегаем его в обратном порядке и складвыем с результирующей строкой
	for i := len(strArray) - 1; i >= 1; i-- {
		resString += strArray[i] + " "
	}
	resString += strArray[0]
	return resString
}

func main() {
	fmt.Println(reverseWords("snow dog sun"))
}
