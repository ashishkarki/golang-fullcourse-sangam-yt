package main

import (
	"fmt"
	"strconv"
)

func rangeUsage() {
	// 1. slices and arrays returns both index and the current item
	fruits := []string{"peach", "banana", "kiwi"}

	for index, fruit := range fruits {
		fmt.Println("The " + strconv.Itoa(index+1) + " fruit is " + fruit)
	}

	// 2. maps returns key value pairs
	ages := map[string]int{"Bob": 20, "Thomas": 80}
	for key, value := range ages {
		fmt.Printf("Key: %s, Age: %d\n", key, value)
	}

	// 3. works with for loop
	for k := range 5 {
		fmt.Printf("Printing with range: %s\n", strconv.Itoa(k))
	}
}
