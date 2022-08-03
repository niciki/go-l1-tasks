package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов
чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

// создаем WaitGroup, увеличиваем ее счетчик на число элементов массива.
// Далее запускаме горутины, в каждой из которых декрементируем счетчик WaitGroup
// После запуска всех горутин ждем, пока счетчик WaitGroup не станет
// равен нулю(все горутины не выполнятся) с помощью метода .Wait()

func v1() {
	fmt.Print("Variant 1: using WaitGroup\n")
	arr := [...]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	wg.Add(len(arr))
	for _, i := range arr {
		go func(i int) {
			fmt.Print(i*i, "\n")
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// Создаем буферизированный канал с буфером равным числу элементов массива, запускаем горутины,
// в каждой из которых пишем в канал элемент. После запуска всех горутин считываем из канала элементы

func v2() {
	fmt.Print("Variant 2: using Buffured Channel\n")
	arr := [...]int{2, 4, 6, 8, 10}
	done := make(chan interface{}, len(arr))
	defer close(done)
	for _, i := range arr {
		go func(i int) {
			fmt.Print(i*i, "\n")
			done <- struct{}{}
		}(i)
	}
	for i := 0; i < len(arr); i++ {
		<-done
	}
}

func main() {
	v1()
	v2()
}
