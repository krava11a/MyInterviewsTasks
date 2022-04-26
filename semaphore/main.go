package main

import (
	"fmt"
	"time"
)

type sema chan struct{}

func New(n int) sema {
	return make(sema, n)
}

func (s sema) Inc(k int) {
	for i := 0; i < k; i++ {
		s <- struct{}{}
		fmt.Println("Increment",i)
	}
	close(s)

}

func (s sema) Decr(k int) {
	for i := 0; i < k; i++ {
		<-s
		fmt.Println("Decrement",i)
	}
}

func main() {

	s :=  New(5)

	for i := 0; i < 5; i++ {
		go func(semka sema) {
			semka.Inc(1)
		}(s)
	}

	time.Sleep(60*time.Second)

	//for i := 0; i < 5; i++ {
	//	go func(semka sema) {
	//
	//		semka.Decr(1)
	//	}(s)
	//}

	//s.Decr(5)

}
