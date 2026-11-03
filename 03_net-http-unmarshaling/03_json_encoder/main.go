package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func successHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	res := map[string]any{
		"ok":       true,
		"message":  "JSON encode success",
		"datetime": time.Now().UTC(),
	}

	_ = json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/ok", successHandler)

	err := http.ListenAndServe(":8092", nil)
	fmt.Println(err)
}
