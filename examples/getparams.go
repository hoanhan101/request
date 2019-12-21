package main

import (
	"fmt"

	"github.com/hoanhan101/request"
)

func main() {
	payload := map[string]string{"k1": "v1", "k2": "v2"}

	r, err := request.Get("https://httpbin.org/get", payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", r)
}
