package request

import (
	"testing"

	"github.com/hoanhan101/requesttest"
	"github.com/stretchr/testify/assert"
)

type Response struct {
	Status  string            `json:"status"`
	Payload map[string]string `json:"payload"`
}

func TestGetJSON(t *testing.T) {
	url, closer := requesttest.Serve("/get", `{"status":"ok"}`)
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
			Status: "ok",
		},
		r,
	)
}

func TestGetJSONQuery(t *testing.T) {
	url, closer := requesttest.Serve("/get", `{"status":"ok","payload":{"k1":"v1"}}`)
	defer closer()

	r := new(Response)
	err := GetJSON(
		&Options{
			URL: url,
			// FIMXE - make test server echo payload
			// Payload: map[string]string{"k1": "v1"}
		},
		r,
	)
	assert.NoError(t, err)
	assert.Equal(
		t,
		&Response{
			Status:  "ok",
			Payload: map[string]string{"k1": "v1"},
		},
		r,
	)
}

func TestPostJSON(t *testing.T) {
	url, closer := requesttest.Serve("/post", `{"status":"ok"}`)
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
			Status: "ok",
		},
		r,
	)
}

func TestPostJSONQuery(t *testing.T) {
	url, closer := requesttest.Serve("/post", `{"status":"ok","payload":{"k1":"v1"}}`)
	defer closer()

	r := new(Response)
	err := PostJSON(
		&Options{
			URL: url,
			// Payload: map[string]string{"k1": "v1"},
		},
		r,
	)

	assert.NoError(t, err)
	assert.Equal(
		t,
		&Response{
			Status:  "ok",
			Payload: map[string]string{"k1": "v1"},
		},
		r,
	)
}
