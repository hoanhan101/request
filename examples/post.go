package main

import (
	"fmt"

	"github.com/hoanhan101/request"
)

func main() {
	payload := map[string][]string{"k1": []string{"v1"}, "k2": []string{"v2", "v3"}}

	r, err := request.Post("https://httpbin.org/post", payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", r)
}