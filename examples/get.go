package main

import (
	"fmt"

	"github.com/hoanhan101/request"
)

func main() {
	r, err := request.Get("https://httpbin.org/get", nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", r)
}
