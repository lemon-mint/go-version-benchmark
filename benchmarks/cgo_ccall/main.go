package main

/*
int Add(int a, int b) {
	return a+b;
}
*/
import "C"

const N = 50000000

func main() {
	for i := (0); i < N; i++ {
		_ = C.Add(C.int(i), C.int(i+1))
	}
}
