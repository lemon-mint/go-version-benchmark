package main

import "regexp"

const N = 1<<16 - 1

func main() {
	var res []*regexp.Regexp = []*regexp.Regexp{
		regexp.MustCompile(".*"),
		regexp.MustCompile("[a-z]+"),
		regexp.MustCompile("(.*)@(.*)"),
	}
	var tests []string = []string{
		"",
		"a",
		"abc",
		"a@b",
		"abcd@efgh",
		"abcdefghijklmnopqrstuvwxyz",
		"abcdefghijklmnopqrstuvwxyz@abcdefghijklmnopqrstuvwxyz",
		"abcdefghijklmnopqrstuvwxyz@abcdefghijklmnopqrstuvwxyz@abcdefghijklmnopqrstuvwxyz",
		"abcdefghijklmnopqrstuvwxyz@abcdefghijklmnopqrstuvwxyz@abcdefghijklmnopqrstuvwxyz@abcdefghijklmnopqrstuvwxyz",
		"abcdefghijklmnopqrstuvwxyz@abcdefghijklmnopqrstuvwxyz@abcdefghijklmnopqrstuvwxyz@abcdefghijklmnopqrstuvwxyz@abcdefghijklmnopqrstuvwxyz",
	}
	for i := 0; i < N; i++ {
		for _, test := range tests {
			for _, res := range res {
				res.FindAllString(test, -1)
			}
		}
	}
}
