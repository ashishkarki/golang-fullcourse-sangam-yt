package main

import (
	"fmt"
)

func prodAndDivide(a, b int) (int, float64) {
	prod := a * b
	dividend := float64(a) / float64(b)

	return prod, dividend
}

func functionsDemo() {
	prod, dividend := prodAndDivide(10, 4)
	fmt.Println(prod, dividend)
}

func namedReturn() (result int) {
	//result := rand.Intn(100)
	return
}
