package request

import (
	"testing"

	"github.com/hoanhan101/requesttest"
	"github.com/stretchr/testify/assert"
)

type Response struct {
	Status string            `json:"status"`
	Query  map[string]string `json:"query"`
}

func TestGet(t *testing.T) {
	url, closer := requesttest.Serve("/get", `{"status":"ok"}`)
	defer closer()

	r, err := Get(url, nil, nil)
	assert.NoError(t, err)
	assert.Equal(
		t,
		`{"status":"ok"}`,
		r,
	)
}

func TestGetQuery(t *testing.T) {
	url, closer := requesttest.Serve("/get", `{"status":"ok","query":{"k1":"v1"}}`)
	defer closer()

	r, err := Get(url, map[string]string{"k1": "v1"}, nil)
	assert.NoError(t, err)
	assert.Equal(
		t,
		`{"status":"ok","query":{"k1":"v1"}}`,
		r,
	)
}

func TestGetJSON(t *testing.T) {
	url, closer := requesttest.Serve("/get", `{"status":"ok"}`)
	defer closer()

	r := new(Response)
	err := GetJSON(url, nil, nil, r)
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
	url, closer := requesttest.Serve("/get", `{"status":"ok","query":{"k1":"v1"}}`)
	defer closer()

	r := new(Response)
	err := GetJSON(url, nil, nil, r)
	assert.NoError(t, err)
	assert.Equal(
		t,
		&Response{
			Status: "ok",
			Query:  map[string]string{"k1": "v1"},
		},
		r,
	)
}

func TestPost(t *testing.T) {
	url, closer := requesttest.Serve("/post", `{"status":"ok"}`)
	defer closer()

	r, err := Post(url, nil, nil)
	assert.NoError(t, err)
	assert.Equal(
		t,
		`{"status":"ok"}`,
		r,
	)
}

func TestPostQuery(t *testing.T) {
	url, closer := requesttest.Serve("/post", `{"status":"ok","query":{"k1":"v1"}}`)
	defer closer()

	r, err := Post(url, nil, nil)
	assert.NoError(t, err)
	assert.Equal(
		t,
		`{"status":"ok","query":{"k1":"v1"}}`,
		r,
	)
}

func TestPostJSON(t *testing.T) {
	url, closer := requesttest.Serve("/post", `{"status":"ok"}`)
	defer closer()

	r := new(Response)
	err := PostJSON(url, nil, nil, r)
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
	url, closer := requesttest.Serve("/post", `{"status":"ok","query":{"k1":"v1"}}`)
	defer closer()

	r := new(Response)
	err := PostJSON(url, nil, nil, r)
	assert.NoError(t, err)
	assert.Equal(
		t,
		&Response{
			Status: "ok",
			Query:  map[string]string{"k1": "v1"},
		},
		r,
	)
}
