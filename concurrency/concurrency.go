package main

import (
	"fmt"
	"time"
)

// ----------- using waitgroup to ensure all work done
// func main() {
// 	// wg can be thought of as a 'counter', which is incremented by us
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	go func() {
// 		count("sheep")

// 		// Done() decrements the counter by 1
// 		wg.Done()
// 	}()

// 	wg.Wait()
// }

// ----------- using channels to send messages
// func main() {
// 	// make initializes an object of slice, map, or chan
// 	// returns type passed, not a pointer
// 	c := make(chan string)
// 	go count("sheep", c)

// 	for {
// 		// blocking - waits for sender - open is bool
// 		msg, open := <-c

// 		if !open {
// 			break
// 		}

// 		fmt.Println(msg)
// 	}
// }

func main() {
	c := make(chan string)
	go count("sheep", c)

	// syntactic sugar - call until closed
	for msg := range c {
		fmt.Println(msg)
	}
}

// channel argument allows communication with caller
// chan [type] tells us what type will be sent via. the channel
func count(s string, c chan string) {
	for i := 1; i <= 5; i++ {
		// arrow syntax is 'send message' this is blocking - waits for receiver
		c <- fmt.Sprintf("%d %s", i, s)
		time.Sleep(time.Millisecond * 500)
	}

	// we are finished, tell the receiver
	// ONLY close channel as a sender
	close(c)
}
