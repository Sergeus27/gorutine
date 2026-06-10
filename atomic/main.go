package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := &sync.WaitGroup{}
	var money atomic.Int32
	const c = 1000

	wg.Add(c)
	for range c {
		go func() {
			defer wg.Done()

			money.Add(1)

		}()
	}

	wg.Wait()
	fmt.Println(money.Load())
}
