package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type Job struct {
	Line string
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func worker(id int, jobs <-chan Job, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		reversed := reverseString(job.Line)
		fmt.Printf("Worker %d processed line: %s\n", id, reversed)
		results <- reversed
	}
}

func main() {
	numWorkers := 5
	jobs := make(chan Job, 100)
	results := make(chan string, 100)

	var wg sync.WaitGroup

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Чтение из файла
	file, err := os.Open("C:\\Users\\quertyy\\go\\src\\6thGo-main\\6\\input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		jobs <- Job{Line: line}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	close(jobs)

	// Ждем завершения воркеров
	wg.Wait()
	close(results)

	// Вывод результатов
	fmt.Println("Reversed lines:")
	for result := range results {
		fmt.Println(result)
	}
}
