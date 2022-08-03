package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func main() {
	var n int
	fmt.Print("Please input number of seconds:")
	fmt.Scan(&n)
	ch := make(chan int)
	// запускаем горутину на отправку данных в канал
	go func() {
		i := 0
		for {
			ch <- i
			i++
		}
	}()
	// создаем канал таймер, в который будет совершена запись спустя n секунд
	timer := time.After(time.Duration(n) * time.Second)
	// считываем данне из канала, пока не будет совершена запись в канал таймера
	for {
		select {
		case <-timer:
			return
		default:
			fmt.Printf("Receive message: %d\n", <-ch)
		}
	}
}
