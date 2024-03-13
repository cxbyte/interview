package main

import (
	"fmt"
	"sync"
)

func main() {
	var numJobs = 10
	var numWorkers = 2

	var jobs = make(chan int, numJobs)
	var workers = make(chan int, numWorkers)

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			worker(i, jobs, workers)
		}(i)
	}

	wg.Wait()

	//go func() {
	//	wg.Wait()
	//	close(workers)
	//}()
}

func worker(i int, jobs chan int, workers chan int) {
	fmt.Println(`worker: `, i)
}
