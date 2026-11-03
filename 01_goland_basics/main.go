package main

import (
	"fmt"
	"strconv"
	"time"
)

var bigName = "Big Guy!!"

func main() {
	//switchStatements()
	//rangeUsage()
	//arraySlicesDemo()
	//mapsDemo()
	//functionsDemo()
	//errorHandlingDemo()
	//deferStatementsDemo()
	//pointersDemo()
	//structsDemo()
	methodsValReceiverDemo()
}

func switchStatements() {
	const day = 24 * time.Hour

	switch day {
	case 24 * time.Hour:
		fmt.Println("Day 24h")
	case 7 * 24 * time.Hour:
		fmt.Println("Day 7h")
	default:
		fmt.Println("Day 24h")
	}
}

func starter() {
	name := "Ashish"
	age, weight := 25, 97

	fmt.Println("name is "+name, ", the age is "+strconv.Itoa(age), ", weight is "+strconv.Itoa(weight))
	//fmt.Println("name is " + name + ", the age is " + strconv.Itoa(age) + ", weight is " + strconv.Itoa(weight))

	// classic for loop
	for i := 0; i < 10; i++ {
		fmt.Println("classic loop i is ", i)
	}

	// for with range
	for j := range 10 {
		fmt.Println("range loop i is ", j)
	}
}
