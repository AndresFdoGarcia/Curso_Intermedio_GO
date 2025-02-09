package main

import (
	"fmt"
	"time"
)

func main() {
	// x := 5
	// // funcion anonima
	// y := func() int {
	// 	return x * 2
	// }()
	// fmt.Println(y)

	//Funcion anonima concurrente
	c := make(chan int)
	go func() {
		fmt.Println("Starting function")
		time.Sleep(time.Second * 5)
		fmt.Println("Ending function")
		c <- 1
	}()
	<-c
}
