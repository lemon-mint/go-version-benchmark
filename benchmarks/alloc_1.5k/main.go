package main

const N = 1<<22 - 1
const N1_5K = 1024 * 1.5

func alloc_1_5k() []byte {
	mem := make([]byte, N1_5K)
	for j := range mem {
		mem[j] = byte(j)
	}
	return mem
}

func main() {
	for i := 0; i < N; i++ {
		func() {
			mem := alloc_1_5k()
			_ = mem
		}()
	}
}
