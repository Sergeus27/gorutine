package main

import (
	"fmt"
)

func main() {
	//создали интовый канал, каналы используются как для синхронизации)
	// так и для доступа к общим ресурсам(обменом данными между горутинами
	ch := make(chan int)

	for i := range 5 {
		go func() { ch <- i }()

	}

	for range 5 {
		val := <-ch

		fmt.Println(val)

	}

}
