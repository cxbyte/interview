package main

import "fmt"

func main() {
	var a = 1
	var b = 1

	fmt.Println(Equals(a, b))
}

func Equals[T any](a, b T) bool {
	return a == b
}
