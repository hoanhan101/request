package main

import (
	"fmt"

	"github.com/hoanhan101/request"
)

type BinResponse struct {
	URL string `json:"url"`
}

func main() {
	r := new(BinResponse)
	err := request.GetJSON("https://httpbin.org/get", r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r.URL)
}
