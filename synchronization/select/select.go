package select_it

import (
	"fmt"
	"time"
)

func SelectIt() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Printf("received: %s\n", msg1)
		case msg2 := <-c2:
			fmt.Printf("received: %s\n", msg2)
		}
	}

}
