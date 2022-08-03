package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает
две числовых переменных a,b, значение которых > 2^20.
*/

func main() {
	var num1, num2 string
	fmt.Print("Please input your numbers\n\tnumber 1: ")
	fmt.Scan(&num1)
	fmt.Print("\tnumber 2: ")
	fmt.Scan(&num2)
	var op rune
	fmt.Print("Please input operation: ")
	reader := bufio.NewReader(os.Stdin)
	op, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err.Error())
	}
	// создаем переменные big.Float переменной библиотеки big длинной арифметики
	a := new(big.Float)
	a.SetString(num1)
	b := new(big.Float)
	b.SetString(num2)
	c := big.NewFloat(0)
	// в зависимости от символа операции выполняем нужную операцию из библиотеки
	switch op {
	case '/':
		fmt.Printf("%f / %f = %f\n", a, b, c.Quo(a, b))
	case '*':
		fmt.Printf("%f / %f = %f\n", a, b, c.Mul(a, b))
	case '+':
		fmt.Printf("%f / %f = %f\n", a, b, c.Add(a, b))
	case '-':
		fmt.Printf("%f / %f = %f\n", a, b, c.Sub(a, b))
	default:
		fmt.Print("Unknown math operation\n")
	}
}
