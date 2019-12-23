package main

import (
	"fmt"

	"github.com/hoanhan101/request"
)

type Response struct {
	Status string `json:"status"`
}

func main() {
	r := new(Response)

	err := request.GetJSON("http://localhost:1323/test", nil, r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", r)
}
