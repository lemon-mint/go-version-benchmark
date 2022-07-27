package main

import "sync"

const N = 5000000

func main() {
	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		wg.Add(1)
		go gofunc(&wg)
	}
	wg.Wait()
}

func gofunc(wg *sync.WaitGroup) {
	var stack [4096]byte
	for i := range stack {
		stack[i] = byte(i)
	}
	_ = stack
	wg.Done()
}
