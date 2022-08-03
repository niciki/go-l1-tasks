package main

import (
	"fmt"
	"strconv"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func setIthBit(i int, bit int, number int64) int64 {
	// устанавливаем бит с помощью битовых операций
	if bit == 1 {
		return number | (1 << i)
	} else {
		return number &^ (1 << i)
	}
}

func main() {
	var i, bit int
	var number int64
	fmt.Println("Please input your number: ")
	fmt.Scan(&number)
	fmt.Println("Please input your bit: ")
	fmt.Scan(&bit)
	fmt.Println("Please input your i: ")
	fmt.Scan(&i)
	fmt.Printf("Before: %s\n", strconv.FormatInt(number, 2))
	number = setIthBit(i, bit, number)
	fmt.Printf("Before: %s\n", strconv.FormatInt(number, 2))

}
