package main

import (
	"sync"
	"time"
)

const N = 1<<22 - 1

func BenchmarkTimeAfter(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		targetTime := (time.Duration(i) * time.Second / time.Duration(n)) + time.Millisecond*50
		wg.Add(1)
		time.AfterFunc(targetTime, func() {
			wg.Done()
		})
	}
	wg.Wait()
}

func main() {
	BenchmarkTimeAfter(N)
}
