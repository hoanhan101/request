package request

import (
	"testing"

	"github.com/hoanhan101/requesttest"
	"github.com/stretchr/testify/assert"
)

// Response describes a response from requesttest.Echo().
type Response struct {
	Method  string `json:"method"`
	Payload string `json:"payload"`
}

func TestGetJSON(t *testing.T) {
	url, closer := requesttest.Echo()
	defer closer()

	r := new(Response)
	err := GetJSON(
		&Options{
			URL: url,
		},
		r,
	)
	assert.NoError(t, err)
	assert.Equal(
		t,
		&Response{
			Method:  "GET",
			Payload: "",
		},
		r,
	)
}

func TestGetJSONQuery(t *testing.T) {
	url, closer := requesttest.Echo()
	defer closer()

	r := new(Response)
	err := GetJSON(
		&Options{
			URL:     url,
			Payload: map[string]string{"k1": "v1", "k2": "v2"},
		},
		r,
	)
	assert.NoError(t, err)
	assert.Equal(
		t,
		&Response{
			Method:  "GET",
			Payload: "k1=v1&k2=v2",
		},
		r,
	)
}

func TestPostJSON(t *testing.T) {
	url, closer := requesttest.Echo()
	defer closer()

	r := new(Response)
	err := PostJSON(
		&Options{
			URL: url,
		},
		r,
	)
	assert.NoError(t, err)
	assert.Equal(
		t,
		&Response{
			Method:  "POST",
			Payload: "",
		},
		r,
	)
}

func TestPostJSONQuery(t *testing.T) {
	url, closer := requesttest.Echo()
	defer closer()

	r := new(Response)
	err := PostJSON(
		&Options{
			URL:     url,
			Payload: map[string]string{"k1": "v1", "k2": "v2"},
		},
		r,
	)
	assert.NoError(t, err)
	assert.Equal(
		t,
		&Response{
			Method:  "POST",
			Payload: "k1=v1&k2=v2",
		},
		r,
	)
}
