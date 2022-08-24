package main

/*
extern void TF();

void Go_Test(int n) {
	int i;
	for (i=0;i<n;i++) {
		TF();
	}
}
*/
import "C"

const N = 5000000

//export TF
func TF() {}

func main() {
	C.Go_Test(C.int(N))
}
