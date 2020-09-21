package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "hello" }()

	msgs := <-messages
	fmt.Println(msgs)
}
