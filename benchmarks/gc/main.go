package main

import (
	"runtime"
)

const PTR_N = 10000000
const N = 30

func main() {
	m := make([]***int, PTR_N)
	for i := 0; i < PTR_N; i++ {
		m[i] = new(**int)
	}
	for i := 0; i < N; i++ {
		runtime.GC()
	}
	runtime.KeepAlive(m)
}
