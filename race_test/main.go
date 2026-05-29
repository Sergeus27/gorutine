package main

import (
	"fmt"
	"slices"
	//"sort"
	"sync"
)

func main() {
	var arr []int
	for range 100 {
		arr = append(arr, dupel_main())
	}
	slices.Sort(arr)
	for i := range arr {
		fmt.Println(arr[i])
	}

}

func dupel_main() int {
	var money int = 0
	wg := &sync.WaitGroup{}

	wg.Add(100)
	for range 100 {
		go func() {
			defer wg.Done()
			money++
		}()
	}
	wg.Wait()
	return money
}
