package main

import "fmt"

func main() {
	generator := make(chan int)
	doubled := make(chan int)

	go Generator(generator)
	go Double(generator, doubled)
	Printer(doubled)
}

func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- value * 2
	}
	close(out)
}

func Printer(c chan int) {
	for value := range c {
		fmt.Println(value)
	}
}
