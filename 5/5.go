package main

import (
	"fmt"
	"log"
	"sync"
)

// Структура для запроса на выполнение операции
type CalculationRequest struct {
	A, B   float64
	Op     string
	Result chan float64
	Error  chan error
}

func calculator(requests chan CalculationRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	for req := range requests {
		var res float64
		var err error

		switch req.Op {
		case "+":
			res = req.A + req.B
		case "-":
			res = req.A - req.B
		case "*":
			res = req.A * req.B
		case "/":
			if req.B == 0 {
				err = fmt.Errorf("деление на ноль")
			} else {
				res = req.A / req.B
			}
		default:
			err = fmt.Errorf("неизвестная операция: %s", req.Op)
		}

		if err != nil {
			req.Error <- err
		} else {
			req.Result <- res
		}
	}
}

func main() {
	requests := make(chan CalculationRequest)

	var wg sync.WaitGroup

	wg.Add(1)
	go calculator(requests, &wg)

	// Создаём клиентские запросы
	operations := []CalculationRequest{
		{A: 10, B: 20, Op: "+", Result: make(chan float64), Error: make(chan error)},
		{A: 15, B: 5, Op: "-", Result: make(chan float64), Error: make(chan error)},
		{A: 7, B: 3, Op: "*", Result: make(chan float64), Error: make(chan error)},
		{A: 40, B: 0, Op: "/", Result: make(chan float64), Error: make(chan error)},
		{A: 16, B: 4, Op: "/", Result: make(chan float64), Error: make(chan error)},
	}

	// Используем WaitGroup для клиентских горутин
	var clientWg sync.WaitGroup

	// Запуск клиентских запросов в отдельных горутинах
	for _, op := range operations {
		clientWg.Add(1)
		go func(op CalculationRequest) {
			defer clientWg.Done()

			// Отправляем запрос на выполнение
			requests <- op

			// Ожидаем либо результат, либо ошибку
			select {
			case res := <-op.Result:
				fmt.Printf("Результат %f %s %f = %f\n", op.A, op.Op, op.B, res)
			case err := <-op.Error:
				log.Printf("Ошибка при выполнении операции %f %s %f: %s\n", op.A, op.Op, op.B, err)
			}
		}(op)
	}

	// Ожидаем завершения всех клиентских горутин
	clientWg.Wait()

	// Закрываем канал с запросами после того, как все клиентские горутины завершились
	close(requests)

	// Ожидаем завершения работы калькулятора
	wg.Wait()

	fmt.Println("Все операции завершены.")
}
