package main

import (
	"fmt"
	"slices"
)

func arraySlicesDemo() {
	// 1. array have fixed sizes
	var marks [5]int
	marks[0] = 1
	marks[1] = 2
	marks[2] = 3
	marks[3] = 4
	marks[4] = 5

	for i := range 5 {
		marks[i] = i + 100
	}
	fmt.Println(marks)

	// 1.1 array literals
	otherMarks := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array with literals: %d\n", otherMarks)

	//// 2. Slices
	// 2.1 first way to declare slice in go
	veggies := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println("This is a slice of veggies:", veggies)
	veggies[0] = "a real vegetable"
	fmt.Println("This is a slice of veggies updated:", veggies)
	fmt.Printf("Length of veggies: %d\n", len(veggies))

	// 2.2 slice: appending
	veggies = append(veggies, "d", "e", "f", "g")
	fmt.Println("This is a slice of veggies appended:", veggies)

	// 2.3 now about length and capacity of slices
	// make([]T, len,  cap)
	scores := make([]int, 0, 5)
	fmt.Printf("1 => scores, its Length and capacity: %d %d %d\n", scores, len(scores), cap(scores))
	scores = append(scores, 100)
	scores = append(scores, 200, 300)
	scores = append(scores, 400, 500)
	fmt.Println("scores after appending exactly 5 items:", scores, len(scores), cap(scores))

	scores = append(scores, 600)
	fmt.Println("scores after appending over cap:", scores, len(scores), cap(scores))

	// older example of deleting element in a slice
	const indexToDelete = 2
	scores = append(scores[:indexToDelete], scores[indexToDelete+1:]...)
	fmt.Println("scores after deleting index 2 item:", scores, len(scores), cap(scores))

	// newer method to deleting element in slice
	// first lets add back the element "300" we deleted above
	scores = slices.Insert(scores, indexToDelete, 300)
	fmt.Println("scores after using slices module at index 2 item:", scores, len(scores), cap(scores))
	scores = slices.Delete(scores, 2, 3)
	fmt.Println("scores after using slices module at index 2 item:", scores, len(scores), cap(scores))

	// range over slices example
	views := []int{10, 20, 30, 40, 50}
	total := 0
	for _, view := range views {
		total += view
	}
	fmt.Println("Total views: ", total)
}
