package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var a int
	a = 10
	y := 8
	fmt.Println(a)
	fmt.Println(y)
	myValue, err := strconv.ParseInt("NaN", 0, 64)
	if err != nil {
		fmt.Println("Error al convertir el string a int\n", err)
	} else {
		fmt.Println(myValue)
	}
	//Forma de uso de la estructura map
	m := make(map[string]int)
	m["key"] = 6
	fmt.Println(m["key"])

	delete(m, "key")
	fmt.Println(m["key"])

	//Forma de uso de la estructura slice
	s := []int{19, 24, 32}
	for index, v := range s {
		fmt.Println(index)
		fmt.Println(v)
	}

	c := make(chan int)
	go doSomething(c)
	<-c
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
