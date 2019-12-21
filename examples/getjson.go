package main

import (
	"fmt"

	"github.com/hoanhan101/request"
)

type BinResponse struct {
	Args map[string]string `json:"args"`
	URL  string            `json:"url"`
}

func main() {
	r := new(BinResponse)

	err := request.GetJSON("https://httpbin.org/get", nil, r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", r)
}
