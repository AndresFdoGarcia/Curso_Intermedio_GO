package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	c := make(chan int, 2)

	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go DoSomething(i, &wg, c)
	}
	wg.Wait()
}

func DoSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Id %d started\n", i)
	time.Sleep(time.Second * 5)
	fmt.Printf("Id %d finished\n", i)
	<-c
}
