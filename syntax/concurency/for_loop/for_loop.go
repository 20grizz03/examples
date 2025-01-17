package main

import (
	"fmt"
	"sync"
)

func main() {
	var mx int

	for i := 1000; i > 0; i-- {
		go func() { // если до go 1.22 i передаётся по ссылке,
			// какое значение первое взяли то и будет максимумом
			if i > mx {
				mx = i
			}
		}()
	}
	fmt.Println("max = ", mx)
}

// ==========================================================
func main1() {
	var mx int
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for i := 1000; i > 0; i-- {
		// i:= i shadowing
		wg.Add(1)
		go func(i int) { // если до go 1.22 i передаётся по ссылке,
			// какое значение первое взяли то и будет максимумом
			if i > mx {
				mu.Lock()
				mx = i
				mu.Unlock()
			}
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("max = ", mx)
}
