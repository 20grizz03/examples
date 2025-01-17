package main

import (
	"fmt"
	"sync"
	"time"
)

func joinChannels(cs ...<-chan int) <-chan int {
	// создаём канал, который будет результатом merge
	res := make(chan int)
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

	go func() {
		// дожидаемся, когда все задачи закончат свою работу
		wg.Wait()
		// закрываем result, чтобы получатель знал, что данные закончились
		close(res)
	}()

	// возвращаем result
	return res
}

func worker() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 21
	}()
	return ch
}

func main() {
	timeStart := time.Now()

	/*_, _ = <-worker(), <-worker()
	аналогично
	<-worker()
	<-worker()
	*/

	<-joinChannels(worker(), worker()) // выполняется параллельно

	fmt.Println(int(time.Since(timeStart).Seconds()))
}
