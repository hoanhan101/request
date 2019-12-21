package main

import (
	"fmt"

	"github.com/hoanhan101/request"
)

type Response struct {
	Args    map[string]interface{} `json:"args"`
	Headers map[string]interface{} `json:"headers"`
	Origin  string                 `json:"origin"`
	URL     string                 `json:"url"`
}

func main() {
	r := new(Response)

	err := request.GetJSON("https://httpbin.org/get", nil, r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", r)
}
