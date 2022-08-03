package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

type ConcurrentMap struct {
	data map[int]int
	// RWmutex для реализации объектов, которые нельзя параллельно писать, но можно читать
	m sync.RWMutex
}

func NewConcurrentMap() ConcurrentMap {
	return ConcurrentMap{data: make(map[int]int)}
}

func (c *ConcurrentMap) Add(key, value int, wg *sync.WaitGroup) {
	// блокируем мьютекс при добавлении элемента
	c.m.Lock()
	c.data[key] = value
	// разблокируем мьютекс после добавления нового элемента
	c.m.Unlock()
	wg.Done()
}

func (c *ConcurrentMap) Get(key int) int {
	// блокируем  мьютекс на чтение для поиска элемента по ключу
	c.m.RLock()
	defer c.m.RUnlock()
	return c.data[key]
}

func (c *ConcurrentMap) PrintAll() {
	// блокируем мьютекс на чтение
	c.m.RLock()
	defer c.m.RUnlock()
	for key, value := range c.data {
		fmt.Println(key, " ", value)
	}
}

var N int = 100

func main() {
	mapa := NewConcurrentMap()
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i <= N; i++ {
		go mapa.Add(i, i*i, &wg)
	}
	wg.Wait()
	mapa.PrintAll()
}
