package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func client(address, name string) {
	req := request{
		Name: name,
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(req); err != nil {
		log.Fatalf("encoding error: %v", err)
	}

	resp, err := http.Post(address, "application/json", &buf)
	if err != nil {
		log.Fatalf("failed to post: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		log.Fatalf("request error: %d", resp.StatusCode)
	}

	var res response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Fatalf("decoding error: %v", err)
	}

	fmt.Println(res.Message)
}
