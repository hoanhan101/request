package main

import (
	"fmt"

	"github.com/hoanhan101/request"
)

// Response is a sample response.
type Response struct {
	Status string      `json:"status"`
	Method string      `json:"method"`
	Query  interface{} `json:"query"`
}

func get(url string, param map[string]string) (*Response, error) {
	r := new(Response)

	err := request.GetJSON(url, param, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func post(url string, body map[string]string) (*Response, error) {
	r := new(Response)

	err := request.PostJSON(url, body, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func main() {
	r, err := get("http://localhost:8000/get", map[string]string{"k1": "v1", "k2": "v2"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", r)

	r, err = post("http://localhost:8000/post", map[string]string{"k1": "v1", "k2": "v2"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", r)
}
