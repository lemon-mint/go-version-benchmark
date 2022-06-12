package main

import "sort"

const N = 1<<10 - 1
const DataSize = 1 << 16

func main() {
	var data []int = make([]int, DataSize)
	for i := 0; i < N; i++ {
		for i := 0; i < DataSize; i++ {
			data[i] = N - i
		}
		sort.Ints(data)
	}
}
