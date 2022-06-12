package main

const N = 1<<10 - 1
const DataSize = 1 << 16

func main() {
	a := make([]int, DataSize)
	ch := make(chan []int)

	for i := 0; i < N; i++ {
		for i := 0; i < len(a); i++ {
			a[i] = DataSize - i
		}
		go MergeSort(a, ch)
		a = <-ch
	}
}

// The following codes are from: @snowmerak
// https://github.com/snowmerak

func InsertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		for j := i; j > 0 && data[j] < data[j-1]; j-- {
			data[j], data[j-1] = data[j-1], data[j]
		}
	}
}

func MergeSort(data []int, ch chan<- []int) {
	if len(data) <= 15 {
		InsertionSort(data)
		ch <- data
		return
	}
	subCh := make(chan []int)
	mid := len(data) / 2
	go MergeSort(data[:mid], subCh)
	go MergeSort(data[mid:], subCh)
	left := <-subCh
	right := <-subCh
	close(subCh)
	newData := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			newData[k] = left[i]
			i++
		} else {
			newData[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		newData[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		newData[k] = right[j]
		j++
		k++
	}
	ch <- newData
}
