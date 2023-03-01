package main

import (
	"fmt"
	"sync"
	"time"
)

func countDown(n int) {
	for n >= 0 {
		fmt.Println(n)
		n--
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	var wg sync.WaitGroup // Blocker for goroutine using wait group
	wg.Add(1)
	go func() {
		countDown(5)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		countDown(10)
		wg.Done()
	}()

	wg.Wait()
} // <- main goroutines ... terminated all other goroutines
