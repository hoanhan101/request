package request_test

import (
	"fmt"

	"github.com/hoanhan101/request"
	"github.com/hoanhan101/requesttest"
)

func ExampleGetJSON() {
	url, closer := requesttest.Echo()
	defer closer()

	r := new(Response)

	err := request.GetJSON(
		&request.Options{
			URL: url,
		},
		r,
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", r)
	// Output: &{Method:GET Payload:}
}

func ExampleGetJSON_query() {
	url, closer := requesttest.Echo()
	defer closer()

	r := new(Response)

	err := request.GetJSON(
		&request.Options{
			URL:     url,
			Payload: map[string]string{"k1": "v1", "k2": "v2"},
		},
		r,
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", r)
	// Output: &{Method:GET Payload:k1=v1&k2=v2}
}

func ExamplePostJSON() {
	url, closer := requesttest.Echo()
	defer closer()

	r := new(Response)

	err := request.PostJSON(
		&request.Options{
			URL: url,
		},
		r,
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", r)
	// Output: &{Method:POST Payload:}
}

func ExamplePostJSON_query() {
	url, closer := requesttest.Echo()
	defer closer()

	r := new(Response)
	err := request.PostJSON(
		&request.Options{
			URL:     url,
			Payload: map[string]string{"k1": "v1", "k2": "v2"},
		},
		r,
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", r)
	// Output: &{Method:POST Payload:k1=v1&k2=v2}
}
