package main

import (
	"fmt"
	"sync"
)

// main - функция, которая запускается первой, когда программа стартует
func main() {

	// создаём каналы
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	// запускаем горутину, которая будет писать в канал a
	go func() {
		// пишем в канал a
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		// закрываем канал, чтобы получатель знал, что данные закончились
		close(a)
	}()

	// запускаем горутину, которая будет писать в канал b
	go func() {
		// пишем в канал b
		for _, num := range []int{20, 10, 30} {
			b <- num
		}
		// закрываем канал, чтобы получатель знал, что данные закончились
		close(b)
	}()

	// запускаем горутину, которая будет писать в канал c
	go func() {
		// пишем в канал c
		for _, num := range []int{200, 300, 100} {
			c <- num
		}
		// закрываем канал, чтобы получатель знал, что данные закончились
		close(c)
	}()

	// читаем из канала, который является результатом merge
	for num := range merge(a, b, c) {
		// печатаем полученное из канала значение
		fmt.Println(num)
	}
}

// функция, которая объединяет несколько каналов в один
func merge(cs ...<-chan int) <-chan int {
	// создаём канал, который будет результатом merge
	res := make(chan int)

	// запускаем горутину, которая будет работать с каналами
	go func() {
		// создаем wait group, чтобы дождаться, когда все горутины
		// закончат свою работу
		wg := &sync.WaitGroup{}

		// добавляем в wait group столько задач, сколько каналов
		wg.Add(len(cs))

		// запускаем для каждого канала свою горутину
		for _, ch := range cs {
			go func(ch <-chan int) {
				// когда мы закончим свою работу, то уведомляем wait group
				defer wg.Done()
				// читаем из канала и пишем в result
				for v := range ch {
					res <- v
				}
			}(ch)

		}

		// дожидаемся, когда все задачи закончат свою работу
		wg.Wait()
		// закрываем result, чтобы получатель знал, что данные закончились
		close(res)
	}()
	// возвращаем result
	return res
}
