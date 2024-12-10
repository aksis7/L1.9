package main

import (
	"fmt"
)

func main() {
	// Массив чисел
	numbers := []int{1, 2, 3, 4, 5}

	// Создаем два канала
	inputChannel := make(chan int)
	outputChannel := make(chan int)

	// Запускаем горутину для записи чисел в первый канал
	go func() {
		for _, num := range numbers {
			inputChannel <- num
		}
		close(inputChannel) // Закрываем канал после записи всех чисел
	}()

	// Запускаем горутину для обработки чисел и записи в outputChannel
	go func() {
		for num := range inputChannel {
			outputChannel <- num * 2
		}
		close(outputChannel) // Закрываем канал после обработки всех чисел
	}()

	// Считываем и выводим числа из второго канала
	for result := range outputChannel {
		fmt.Println(result)
	}
}
