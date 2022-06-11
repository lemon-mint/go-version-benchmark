package main

const N = 1<<30 - 1

func main() {
	v := 0
	for i := 0; i < N; i++ {
		v += i
	}
	_ = v
}
