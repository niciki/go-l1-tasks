package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора
количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы
всех воркеров.
*/

func worker(num int, ch <-chan int, wg *sync.WaitGroup) {
	for message := range ch {
		fmt.Printf("Worker №%d text: %d\n", num+1, message)
	}
	wg.Done()
}

func main() {
	var n int
	fmt.Println("Please, input number of workers")
	fmt.Scan(&n)
	ch := make(chan int)
	var wg sync.WaitGroup
	// запускаем горутины для записи в канал
	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(i, ch, &wg)
	}
	i := 0
	// канал, сообщающий о завершении программы по нажатию Ctrl+C.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	for {
		// считываем данные из канала, пока не будет сообщено о завершении программы
		select {
		case <-sigs:
			close(ch)
			wg.Wait()
			return
		default:
			ch <- i
			i++
		}
	}
}
