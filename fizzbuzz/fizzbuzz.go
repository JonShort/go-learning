package main

import (
	"fmt"
)

func main() {
	for out := range fizzbuzz(100) {
		fmt.Println(out)
	}
}

func fizzbuzz(amount int) <-chan string {
	out := make(chan string, amount)

	go func() {
		for i := 1; i <= amount; i++ {
			result := ""

			if i%3 == 0 {
				result += "fizz"
			}
			if i%5 == 0 {
				result += "buzz"
			}
			if result == "" {
				result = fmt.Sprint(i)
			}

			out <- result
		}

		close(out)
	}()

	return out
}
