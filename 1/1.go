package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция для расчета факториала числа
func factorial(n int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("Факториал %d = %d\n", n, result)
}

// Функция для генерации случайных чисел
func randomNumbers(count int) {
	for i := 0; i < count; i++ {
		num := rand.Intn(100)
		fmt.Printf("Случайное число: %d\n", num)
		time.Sleep(200 * time.Millisecond)
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
	go randomNumbers(5)
	go sumSeries(10)

	time.Sleep(3 * time.Second)

	fmt.Println("Все функции завершены.")
}
