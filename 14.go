package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func detectType(variable interface{}) {
	// используем оператор switch, чтобы определить динамический тип переменной
	switch variable.(type) {
	case int:
		fmt.Printf("Type int: %d\n", variable)
	case string:
		fmt.Printf("Type string: %s\n", variable)
	case bool:
		fmt.Printf("Type bool: %t\n", variable)
	case chan int:
		fmt.Printf("Type chan int: %v\n", variable)
	default:
		fmt.Printf("Unexpected type: %v\n", variable)
	}
}

func main() {
	detectType(interface{}(1))
	detectType(interface{}("mama"))
	detectType(interface{}(true))
	detectType(interface{}(make(chan int)))
	detectType(interface{}(8.9))
}
