package main

import (
	"bytes"
	"fmt"
)

const N = 1<<31 - 1

func main() {
	var data []uint32
	data = make([]uint32, 32)
	for i := 0; i < 32; i++ {
		data[i] = uint32(i)
	}

	var sum uint32
	for i := 0; i < N; i++ {
		switch data[i%32] {
		case 0:
			sum += data[i%32] + 0
		case 1:
			sum += data[i%32] + 1
		case 2:
			sum += data[i%32] + 2
		case 3:
			sum += data[i%32] + 3
		case 4:
			sum += data[i%32] + 4
		case 5:
			sum += data[i%32] + 5
		case 6:
			sum += data[i%32] + 6
		case 7:
			sum += data[i%32] + 7
		case 8:
			sum += data[i%32] + 8
		case 9:
			sum += data[i%32] + 9
		case 10:
			sum += data[i%32] + 10
		case 11:
			sum += data[i%32] + 11
		case 12:
			sum += data[i%32] + 12
		case 13:
			sum += data[i%32] + 13
		case 14:
			sum += data[i%32] + 14
		case 15:
			sum += data[i%32] + 15
		}
	}

	var b bytes.Buffer
	fmt.Fprintf(&b, "%d", sum)
	return
}
