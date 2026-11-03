package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("going to port 8090..")
	err := http.ListenAndServe(":8090", nil)

	fmt.Println("error creating server:", err)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	_, _ = w.Write([]byte("Hello, World! from net/http server"))
}
