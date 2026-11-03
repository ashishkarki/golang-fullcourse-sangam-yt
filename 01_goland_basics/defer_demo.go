package main

import (
	"errors"
	"fmt"
)

func deferStatementsDemo() {
	fmt.Println("Case 1: success")
	if err := doWork(true); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Case 2: failure")
	if err := doWork(false); err != nil {
		fmt.Println("failure error:", err)
	}
}

func doWork(success bool) error {
	// resource related
	// start msg -> resource acquired
	// cleanup msg -> resource released
	fmt.Println("start: resource acquired")

	defer fmt.Println("end: resource released")

	if !success {
		return errors.New("something bad happened")
	}

	fmt.Println("work: doing some work")
	fmt.Println("work: this is done")

	return nil
}
