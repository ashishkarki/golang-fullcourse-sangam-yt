package main

import (
	"02_go-modules/internal/greet"
	"fmt"
)

func main() {
	msg1 := greet.Hello("Gopher")
	fmt.Println(msg1)
}
