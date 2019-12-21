package main

import (
	"fmt"

	"github.com/hoanhan101/request"
)

func main() {
	r, err := request.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r)
}
