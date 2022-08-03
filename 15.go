package main

import (
	"fmt"
	"math/rand"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

*/

/*
Плохие вещи в коде:
1. Не понятна суть использования переменной v, после создания слайса из нее остальные 1024-100=924 элемента
являются неиспользуемыми, поэтому можно сразу от результата функции взять только первые 100 элементов
2. Глобальные переменные не рекомендуютяс к использованию, поэтому определение justString можно сделать
и в функции(либо передать по указателю в функцию)

*/

func createHugeString(num int) string {
	str := []rune("0123456789абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	resultString := ""
	for i := 0; i < num; i++ {
		resultString += string(str[rand.Intn(len(str))])
	}
	return resultString
}

func someFunc() string {
	return createHugeString(1 << 10)[:100]
}

func main() {
	fmt.Print(someFunc())
}
