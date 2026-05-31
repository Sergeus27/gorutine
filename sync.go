package main

import (
	"fmt"
	"sync"
)

type Service struct {
	db    IDatabase
	cache ICache
}

type ICache interface {
	Get(int) int
}

type IDatabase interface {
}

func main() {

}
