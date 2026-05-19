package main

import (
	"fmt"
	"sync"
)

func foo(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello from foo")
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("hello from gorutine")
	}() //анонимная функция внутри которой все будет запущено конкурентно
	//анонимные функции?
	//конкурентные функции?
	//горутина это функция которая может запускаться конкурентно и работать конкурентно вместе с другими функциями

	wg.Add(1)
	go foo(wg)

	wg.Wait() //ждет пока не отработают все Done
	fmt.Println("hello from main")
}
