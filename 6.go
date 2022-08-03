package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func UsingChanClose(ch <-chan int, wg *sync.WaitGroup) {
	// читаем данные из канала, пока он не будет закрыт
	for i := range ch {
		fmt.Print(i, " ")
	}
	fmt.Print("goroutine death\n")
	wg.Done()
	return
}

func UsingChan(exit <-chan interface{}, wg *sync.WaitGroup) {
	// выводим данные, пока в канал выхода не будет совершена запись
	for {
		select {
		case <-exit:
			fmt.Print("goroutine death\n")
			wg.Done()
			return
		default:
			fmt.Print("Goroutine is alive\n")
			time.Sleep(1 * time.Second)
		}
	}
}

func UsingContext(ctx context.Context, wg *sync.WaitGroup) {
	// выводим данные, пока контекст не будет закрыт
	for {
		select {
		case <-ctx.Done():
			fmt.Print("goroutine death\n")
			wg.Done()
			return
		default:
			fmt.Print("goroutine with context is alive\n")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// остановка выполнения горутины с использованием закрытия канала
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int)
	go UsingChanClose(ch, &wg)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
	// остановка выполнения горутины с использованием канала закрытия
	wg.Add(1)
	exit := make(chan interface{})
	go UsingChan(exit, &wg)
	time.Sleep(1 * time.Second)
	exit <- struct{}{}
	wg.Wait()
	// остановка выполнения горутины с использованием контекстов
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go UsingContext(ctx, &wg)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
}
