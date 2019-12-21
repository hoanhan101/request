package main

import (
	"fmt"

	"github.com/hoanhan101/request"
)

type Response struct {
	Args    map[string]string      `json:"args"`
	Data    string                 `json:"data"`
	Files   map[string]string      `json:"files"`
	Form    map[string]interface{} `json:"form"`
	Headers map[string]string      `json:"headers"`
	JSON    interface{}            `json:"json"`
	Origin  string                 `json:"origin"`
	URL     string                 `json:"url"`
}

func main() {
	r := new(Response)
	payload := map[string][]string{"k1": []string{"v1"}, "k2": []string{"v2", "v3"}}

	err := request.PostJSON("https://httpbin.org/post", payload, r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", r)
}
