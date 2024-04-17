package main

import "fmt"

func main() {
	input := []int{1, 2, 3}

	doubledChannel := make(chan int)
	go func() {
		defer close(doubledChannel)
		for _, val := range input {
			doubledChannel <- val * 2
		}
	}()

	squaredChannel := make(chan int)
	go func() {
		defer close(squaredChannel)
		for val := range doubledChannel {
			squaredChannel <- val * val
		}
	}()

	for value := range squaredChannel {
		fmt.Println(value)
	}
}
