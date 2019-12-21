package main

import (
	"fmt"

	"github.com/hoanhan101/request"
)

type BinResponse struct {
	Args map[string]interface{} `json:"args"`
	URL  string                 `json:"url"`
}

func main() {
	r := new(BinResponse)
	payload := map[string][]string{"k1": []string{"v1"}, "k2": []string{"v2", "v3"}}

	err := request.GetJSON("https://httpbin.org/get", payload, r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", r)
}
