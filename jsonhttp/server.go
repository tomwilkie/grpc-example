package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func server(addr string) {
	http.Handle("/greeter", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := response{
			Message: "Hello " + req.Name,
		}

		if err := json.NewEncoder(w).Encode(&res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))
	log.Fatalf("listen err: %v", http.ListenAndServe(addr, nil))
}
