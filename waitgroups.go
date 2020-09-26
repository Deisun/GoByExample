package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	println("running main...")
	go foo()
	go bar()

	wg.Wait()
}

func foo() {
	time.Sleep(time.Second * 3)
	println("running foo...")
	wg.Done()
}

func bar() {
	time.Sleep(time.Second * 5)
	println("running bar...")
	wg.Done()
}
