package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 2; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func printFibonacci(c chan int) {
	for n := range c {
		fmt.Println(n)
	}
}

func main() {
	c := make(chan int)

	go fibonacci(10, c)  // Запускаем горутину для генерации чисел Фибоначчи
	go printFibonacci(c) // Запускаем горутину для вывода чисел на экран

	time.Sleep(1 * time.Second)
}
