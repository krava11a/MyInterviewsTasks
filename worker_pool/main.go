package main

import "fmt"

func worker(id int, f func(int) int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- f(j)
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	testFunction := func(x int) int {
		return x
	}

	for w := 1; w <= 2; w++ {
		go worker(w, testFunction, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		fmt.Println(<-results)
	}
}
