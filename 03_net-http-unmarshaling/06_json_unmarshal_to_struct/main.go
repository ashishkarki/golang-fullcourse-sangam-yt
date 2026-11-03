package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CatFactResponse struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func main() {
	const url = "https://catfact.ninja/fact"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error closing resp.Body: ", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("bad status:", resp.Status)
		return
	}

	// json unmarshal takes bodybytes to struct
	var data CatFactResponse

	// Method 1: directly read from the stream
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("failed to decode response body: ", err)
		return
	}

	// Method 2: load into memory and then convert to json
	//bodyBytes, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("failed to read response body: ", err)
	//	return
	//}
	//if err := json.Unmarshal(bodyBytes, &data); err != nil {
	//	fmt.Println("failed to unmarshal response body: ", err)
	//	return
	//}

	fmt.Println("data fact:", data.Fact)
}
