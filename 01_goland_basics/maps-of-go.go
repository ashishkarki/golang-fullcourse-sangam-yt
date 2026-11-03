package main

import "fmt"

func mapsDemo() {
	// map[key-type]value-type
	ages := map[string]int{
		"James":  32,
		"Miss":   36,
		"Jhandu": 100,
	}

	// delete from map
	delete(ages, "James")
	fmt.Println(ages)

	// get from map
	age, ok := ages["James"]
	if ok {
		fmt.Printf("%d is an existing age\n", age)
	} else {
		fmt.Printf("%d is NOT an existing age\n", age)
	}

	//scores = make(map[string]int)
	//scores["math"] = 200
}
