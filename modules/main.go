package main

import (
	"github.com/donvito/hellomod"
	hellomodv2 "github.com/donvito/hellomod/v2"
)

func main() {
	hellomod.SayHello()
	hellomodv2.SayHello("Andres\n")
}
