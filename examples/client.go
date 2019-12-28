package main

import (
	"fmt"

	"github.com/hoanhan101/request"
)

// Response is a sample response.
type Response struct {
	Status  string      `json:"status"`
	Method  string      `json:"method"`
	Payload interface{} `json:"payload"`
}

func main() {
	r := new(Response)

	_ = request.GetJSON(
		&request.Options{
			URL:     "http://localhost:8000/get",
			Payload: map[string]string{"k1": "v1", "k2": "v2"},
		},
		r,
	)
	fmt.Printf("%+v\n", r)

	_ = request.PostJSON(
		&request.Options{
			URL:     "http://localhost:8000/post",
			Payload: map[string]string{"k1": "v1", "k2": "v2"},
		},
		r,
	)

	fmt.Printf("%+v\n", r)
}
