package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var money atomic.Int32
	var donationsCount atomic.Int32
	const c = 1000

	go func() {
		for {
			m := money.Load()
			dc := donationsCount.Load()

			if m != dc {
				fmt.Println("money=", m, "donation=", dc)
				break
			}

		}
	}()

	wg := &sync.WaitGroup{}

	wg.Add(c)
	for range c {
		go func() {
			defer wg.Done()

			money.Add(1)
			donationsCount.Add(1)

		}()
	}

	wg.Wait()
	fmt.Println(money.Load())
}
