package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func Sleep(d time.Duration) {
	// создаем канал, в который будет совершена запись через указанное время
	// и тут же дожидаемся записи в него
	<-time.After(d)
}

func main() {
	fmt.Print("Before sleep function\n")
	Sleep(5 * time.Second)
	fmt.Print("After sleep function\n")
}
