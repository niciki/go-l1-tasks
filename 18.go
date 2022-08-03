package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/
type ConcurrencCounter struct {
	cnt int
	mtx sync.RWMutex
}

func (c *ConcurrencCounter) Increment(wg *sync.WaitGroup) {
	// блокируем мьютекс на запись
	c.mtx.Lock()
	c.cnt++
	// разблокируем
	c.mtx.Unlock()
	wg.Done()
}

func (c *ConcurrencCounter) GetCounter() int {
	// блокируем мьютек на чтение
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.cnt
}

func main() {
	var cc ConcurrencCounter
	var wg sync.WaitGroup
	// добавляем в waitgtoup количество горутин, чтобы дождаться их выполнения
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go cc.Increment(&wg)
	}
	wg.Wait()
	fmt.Println(cc.GetCounter())
}
