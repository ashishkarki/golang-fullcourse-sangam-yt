package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// defer close the body, always
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("response Status is not okay:", resp.Status)
		return
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading body:", err)
		return
	}

	bodyString := string(bodyBytes)
	const maxChars = 250
	if len(bodyString) > maxChars {
		fmt.Println("bodyString is too large, max allowed is ", maxChars)
	} else {
		fmt.Println("bodyString is ", bodyString)
	}
}
