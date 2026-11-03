package main

import "fmt"

// 1. first define the blueprint
type Todo struct {
	Title    string
	IsDone   bool
	Priority int
}

type User struct {
	ID      int
	Name    string
	Age     int
	Address string
}

func structsDemo() {
	// 2. create instance using struct literal
	task1 := Todo{
		Title:    "Task 1",
		IsDone:   false,
		Priority: 10,
	}

	user1 := User{
		ID:      1,
		Name:    "John Doe",
		Age:     20,
		Address: "100 Fortune way, Heaven 9999",
	}

	// 3. access or update fields
	fmt.Println(task1, user1)
	task1.IsDone = true
	fmt.Println("task1 is done: ", task1)

	// anonymous structs
	usr := struct {
		Name string
		Age  int
	}{
		Name: "Jack",
		Age:  20,
	}
	fmt.Println("printing anonymous struct: ", usr)
}
