package main

type request struct {
	Name string `json:"name"`
}

type response struct {
	Message string `json:"message"`
}
