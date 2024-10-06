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
		for i := 0; i < 10; i++ {
			numberChannel <- rand.Intn(100)
			time.Sleep(time.Millisecond)
		}
		close(numberChannel)
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

	for {
		select {
		case num, ok := <-numberChannel:
			if !ok {
				fmt.Println("Канал numberChannel закрыт")
				return
			}
			fmt.Printf("Случайное число: %d\n", num)
		case parity, ok := <-parityChannel:
			if !ok {
				fmt.Println("Канал parityChannel закрыт")
				return
			}
			fmt.Printf("Чётность/Нечётность: %s\n", parity)
		}
	}
}
