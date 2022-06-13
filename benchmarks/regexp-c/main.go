package main

import "regexp"

const N = 1<<18 - 1

func main() {
	for i := 0; i < N; i++ {
		regexp.MustCompile(".*")
		regexp.MustCompile("[a-z]+")
		regexp.MustCompile("(.*)@(.*)")
	}
}
