package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	//f1()
	f2()
}

func f1() {
	ch := make(chan int)

	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Printf("%d\n", i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // zero memory allocation used

func f2() {
	go selectLogger()

	//defer func() { close(logCh) }()

	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App shutting down"}

	time.Sleep(100 * time.Millisecond)

	doneCh <- struct{}{}
}

func logger() {
	for entry := range logCh {
		fmt.Printf("%v - [%v] %v\n", entry.time, entry.severity, entry.message)
	}
}

func selectLogger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v] %v\n", entry.time, entry.severity, entry.message)
		case <-doneCh:
			break
		default:
		}
	}
}
