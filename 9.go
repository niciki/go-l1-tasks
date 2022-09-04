package main

import "fmt"

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	chan1 := make(chan int)
	chan2 := make(chan int)
	// записываем параллельно числа из массива в канал
	go func() {
		for _, i := range arr {
			chan1 <- i
		}
		close(chan1)
	}()
	//считываем данные из канала chan1, умножаем на 2 и пишем в chan2
	go func() {
		for val := range chan1 {
			chan2 <- val * 2
		}
		close(chan2)
	}()
	// выводим данные из канала chan2
	for i := range chan2 {
		fmt.Println(i)
	}
}
