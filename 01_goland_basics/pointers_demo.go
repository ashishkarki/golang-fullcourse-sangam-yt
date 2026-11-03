package main

import "fmt"

func pointersDemo() {
	// &: address operator -> makes a pointer of address
	// *: deref operator -> enables looking into and manipulating value
	score := 10
	fmt.Println("before:", score)
	addScore(&score)
	fmt.Println("after:", score)
}

func addScore(score *int) {
	*score = *score + 5
}
