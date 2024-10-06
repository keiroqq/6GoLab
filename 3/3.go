package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numberChannel := make(chan int)
	parityChannel := make(chan string)

	// Горутина для генерации случайных чисел
	go func() {
		for {
			num := rand.Intn(100) // Генерируем случайное число от 0 до 99
			numberChannel <- num
			time.Sleep(time.Second) // Замедляем генерацию
		}
	}()

	// Горутина для проверки четности/нечетности
	go func() {
		for {
			num := <-numberChannel
			if num%2 == 0 {
				parityChannel <- fmt.Sprintf("%d: четное", num)
			} else {
				parityChannel <- fmt.Sprintf("%d: нечетное", num)
			}
		}
	}()

	// Горутина для получения и вывода сообщений
	go func() {
		for msg := range parityChannel {
			fmt.Println(msg)
		}
	}()

	// Бесконечный цикл, чтобы основная программа не завершалась
	select {}
}
