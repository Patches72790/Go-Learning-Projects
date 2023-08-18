package main

import (
	"fmt"
	"sync"
	select_it "synchronization/select"
	timers "synchronization/timer"
	"synchronization/workerpool"
	"time"
)

func worker(done chan<- bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func wg_worker(i int) {
	fmt.Printf("Worker %d starting...", i)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", i)
}

func DoWork() {
	done := make(chan bool)
	var wg sync.WaitGroup

	go worker(done)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i

		go func() {
			defer wg.Done()
			wg_worker(i)
		}()
	}

	wg.Wait()
	<-done
}

func main() {
	workerpool.PoolIt()
	select_it.SelectIt()
	timers.TimeIt()
}
