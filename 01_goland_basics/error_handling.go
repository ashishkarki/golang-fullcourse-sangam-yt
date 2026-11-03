package main

import (
	"fmt"
	"log"
	"strconv"
)

func errorHandlingDemo() {
	fmt.Println("--- Testing a VALID input ---")
	if err := run("30"); err != nil {
		fmt.Println("error demo caught an error: ", err)
	}

	fmt.Println("--- Testing a INVALID output ---")
	if err := run("abc"); err != nil {
		// following will print the error and then crashes/exits instantly
		log.Fatal("error demo caught an error: ", err)
	}

	// if there is an error, following line will never be printed/executed
	fmt.Println("This line is unreachable if the input is invalid!!!!")
}

func run(input string) error {
	level, err := parseLevel(input)
	if err != nil {
		// if parseLevel fails, we stop and return error
		return err
	}

	// only if no error
	fmt.Println("selected level:", level)
	return nil
}

func parseLevel(s string) (int, error) {
	n, err := strconv.Atoi(s)

	if err != nil {
		// if its not number, we return a fallback value and a custom error
		return 0, fmt.Errorf("level must be a number: %w", err)
	}

	// if successful, we return the parsed integer and nil error
	return n, nil
}
