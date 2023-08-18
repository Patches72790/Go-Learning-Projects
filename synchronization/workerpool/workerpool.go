package workerpool

import (
	"fmt"
	"math/rand"
)

func worker(id int, jobs <-chan int, results chan<- int, done chan bool) {
	for j := range jobs {

		fmt.Printf("worker %d started job %d\n", id, j)

		fmt.Printf("worker %d finished job %d\n", id, j)

		results <- j * rand.Int()
	}

	done <- true

}

func PoolIt() {
	const numJobs = 5
	done := make(chan bool)
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		w := w
		go worker(w, jobs, results, done)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	<-done

	for r := 1; r <= numJobs; r++ {
		fmt.Printf("Result %d: %d\n", r, <-results)
	}
}
