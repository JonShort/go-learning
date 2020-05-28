package main

import "fmt"

func main() {
	// we can give the channel a 'capacity' after which it becomes blocking
	c := make(chan string, 2)
	// sending two is non-blocking
	c <- "hello"
	c <- "world"

	// if a third message is sent, it becomes blocking
	// c <- "!"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}
