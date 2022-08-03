package main

import "fmt"

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.


Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
func MakeMap(arr []float64) map[int][]float64 {
	// строим мапу ключ - слайс float, обходя слайс температур, считая значение десятка
	data := make(map[int][]float64, 0)
	for _, i := range arr {
		data[int(i/10)*10] = append(data[int(i/10)*10], i)
	}
	return data
}

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Printf("%-#v\n", MakeMap(arr))
}