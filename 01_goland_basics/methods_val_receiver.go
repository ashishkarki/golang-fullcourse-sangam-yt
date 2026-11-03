package main

import "fmt"

type user struct {
	Name string
	Age  int
}

func methodsValReceiverDemo() {
	//// case A: standard value variable
	standardUser := User{
		ID:      0,
		Name:    "John Doe",
		Age:     20,
		Address: "1 Mountain way",
	}

	// Quirk 2 (Auto-referencing):
	// 'Birthday' requires a pointer receiver (*User), but standardUser is just a value.
	// Go automatically interprets this as: (&standardUser).Birthday()
	standardUser.Birthday()
	fmt.Println("New Age: ", standardUser.Age)

	// --- CASE B: Starting with a POINTER variable ---
	// Using '&' to create a pointer to a user directly
	pointerUser := &User{Name: "John Doe", Age: 33}
	// Quirk 3 (Auto-dereferencing):
	// 'Greets' requires a value receiver (User), but pointerUser is a pointer (*User).
	// Go automatically interprets this as: (*pointerUser).Greets()
	message := pointerUser.Greets()
	fmt.Println(message)
}

// Greets 1. Greets val receiver means receives a copy of the data
func (u User) Greets() string {
	return fmt.Sprintf("Hi, I am %s", u.Name)
}

// Birthday Pointer receiver
func (u *User) Birthday() {
	u.Age += 1 // Quirk 1: We don't have to type (*u).Age! Go handles it.
}
