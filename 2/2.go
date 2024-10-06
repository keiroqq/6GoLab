package main

import (
	"fmt"
)

func fibonacci(n int, ch chan<- int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch) // Закрываем канал после завершения генерации
}

func main() {
	n := 10
	ch := make(chan int)

	go fibonacci(n, ch) // Запускаем горутину для генерации чисел Фибоначчи

	// Чтение из канала и вывод на экран
	for num := range ch {
		fmt.Println(num)
	}
}
