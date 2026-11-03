package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const url = "https://catfact.ninja/fact"

type CatFactResponse struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func writeJson(w http.ResponseWriter, status int, data map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func fetchCatFact() (CatFactResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return CatFactResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return CatFactResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return CatFactResponse{}, err
	}

	var data CatFactResponse
	if err = json.Unmarshal(bodyBytes, &data); err != nil {
		return CatFactResponse{}, err
	}

	return data, nil
}

func externalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJson(w, http.StatusMethodNotAllowed, map[string]any{
			"message": "only POST method is allowed",
			"ok":      false,
		})
		return
	}

	data, err := fetchCatFact()
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]any{
			"message": "failed to fetch cat fact!!",
			"ok":      false,
		})
		return
	}

	writeJson(w, http.StatusOK, map[string]any{
		"ok":   true,
		"time": time.Now().UTC(),
		"external": map[string]any{
			"source": "Catfact.ninja",
			"fact":   data.Fact,
			"length": data.Length,
		},
	})
}

func main() {
	const url = "https://catfact.ninja/fact"

	http.HandleFunc("/external", externalHandler)

	err := http.ListenAndServe(":8092", nil)
	fmt.Println(err)
}
