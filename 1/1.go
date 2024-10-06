package main

import (
	"fmt"
	"time"
)

// Функция для расчета факториала числа
func factorial(n int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		fmt.Printf("Факториал %d = %d\n", i, result)
		time.Sleep(100 * time.Millisecond)
	}

}

// Функция для вывода ряда Фибоначчи до номера n
func fibonacci(n int) {
	f1, f2 := 0, 1
	for i := 2; i < n; i++ {
		f1, f2 = f2, f1+f2
		fmt.Printf("Число ряда Фибоначчи под номером %d = %d\n", i, f2)
		time.Sleep(150 * time.Millisecond)
	}
}

// Функция для вычисления суммы числового ряда
func sumSeries(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		time.Sleep(150 * time.Millisecond)
	}
	fmt.Printf("Сумма ряда до %d = %d\n", n, sum)
}

func main() {
	// Запуск горутин
	go factorial(5)
	go fibonacci(6)
	go sumSeries(10)

	time.Sleep(3 * time.Second)

	fmt.Println("Все функции завершены.")
}
