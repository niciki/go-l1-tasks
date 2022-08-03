package main

import "fmt"

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2* 2 + 3 * 3 + 4 * 4 …)
с использованием конкурентных вычислений.
*/

func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	done := make(chan int, len(arr))
	defer close(done)
	// выполняем запуск горутин, считающих произведение и записывающих его в канал
	for _, i := range arr {
		go func(i int) {
			done <- i * i
		}(i)
	}
	sum := 0
	// считывание данных из канала и их суммирование
	for i := 0; i < len(arr); i++ {
		sum += <-done
	}
	fmt.Print(sum, "\n")
}
