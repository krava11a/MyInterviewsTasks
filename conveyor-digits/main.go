package main

import "fmt"

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x <= 10; x++ {
			naturals <- x
		}
		close(naturals) //NB!!!!
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares) //NB!!!!
	}()

	for x := range squares {
		fmt.Println(x)
	}

}
