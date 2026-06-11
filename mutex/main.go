package main

import (
	"fmt"
	"sync"
)

func main() {

	var money int32
	var donationsCount int32
	const c = 1000
	mutex := &sync.RWMutex{}
	//атомик самый низкоуровневый примитив(не знает ничего о состоянии горутины)
	//мьютекс чуть более высокоуровневый (мьютекс построкны на атомике)
	//канал самый выскоуровенвый примитив (тоже построен на атомиках)
	go func() {
		for {

			mutex.RLock()        //мьютекс только для чтения,
			m := money           // чтобы можно было читать сразу нескольким пользователям.
			dc := donationsCount //Писать во время этой блокировки никто не будет но читать могут сразу несколько
			mutex.RUnlock()

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

			mutex.Lock()
			money++
			donationsCount++
			mutex.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(money)
}
