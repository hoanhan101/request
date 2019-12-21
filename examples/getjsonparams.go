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
	payload := map[string][]string{"k1": []string{"v1"}, "k2": []string{"v2", "v3"}}

	err := request.GetJSON("https://httpbin.org/get", payload, r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", r)
}
