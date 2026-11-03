package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type TestRequest struct {
	Name string `json:"name"`
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	// 1. first check if it is POST request
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "method not allowed",
		})
		return
	}

	// at the end close the stream?
	defer r.Body.Close()

	// 2. if post method, read the request body
	var req TestRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Invalid json format",
		})
		return
	}

	req.Name = strings.TrimSpace(req.Name)

	if req.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "name is required!!!",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"ok":        true,
		"data":      req.Name,
		"timestamp": time.Now().UTC(),
	})
}

func main() {
	http.HandleFunc("/test", testHandler)

	err := http.ListenAndServe(":8092", nil)
	if err != nil {
		log.Println(err)
	}
}
