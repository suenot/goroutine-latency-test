package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем канал
	ch := make(chan int, 1)

	// Измеряем задержку
	startTime := time.Now()
	
	// Отправка в горутине
	go func() {
		ch <- 1
	}()
	
	// Чтение
	<-ch
	
	// Вывод задержки в миллисекундах
	duration := time.Since(startTime)
	fmt.Printf("Задержка передачи данных по каналу: %v\n", duration)
}
