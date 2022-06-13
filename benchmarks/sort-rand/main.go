package main

import (
	"math/rand"
	"sort"
)

const N = 1<<10 - 1
const DataSize = 1 << 16

func main() {
	s := rand.NewSource(42)
	var data []int = make([]int, DataSize)
	for i := 0; i < N; i++ {
		for i := 0; i < DataSize; i++ {
			data[i] = int(s.Int63())
		}
		sort.Ints(data)
	}
}
