package main

import (
	"fmt"
	"sync"
	"time"
)

func counter(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println(name, i)
	}
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	go counter("one")
	go counter("two")

	time.Sleep(time.Second)
	fmt.Println("done")

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

}
