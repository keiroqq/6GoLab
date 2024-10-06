package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		// Блокировка мьютекса для предотвращения гонки данных
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup

	// Запускаем 10 горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()

	fmt.Println("Final counter value with mutex:", counter)

	// Теперь уберем мьютекс для сравнения
	counter = 0
	wg.Add(0)

	// Запускаем 10 горутин без мьютекса
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value without mutex:", counter)
}
