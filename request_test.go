package request

import (
	"testing"

	"github.com/hoanhan101/requesttest"
	"github.com/stretchr/testify/assert"
)

type Reponse struct {
	Status string            `json:"status"`
	Query  map[string]string `json:"query"`
}

func TestGet(t *testing.T) {
	url, closer := requesttest.Serve("/get", `{"status":"ok"}`)
	defer closer()

	r, err := Get(url, nil)
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

	r, err := Get(url, map[string]string{"k1": "v1"})
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

	r := new(Reponse)
	err := GetJSON(url, nil, r)
	assert.NoError(t, err)
	assert.Equal(
		t,
		&Reponse{
			Status: "ok",
		},
		r,
	)
}

func TestGetJSONQuery(t *testing.T) {
	url, closer := requesttest.Serve("/get", `{"status":"ok","query":{"k1":"v1"}}`)
	defer closer()

	r := new(Reponse)
	err := GetJSON(url, nil, r)
	assert.NoError(t, err)
	assert.Equal(
		t,
		&Reponse{
			Status: "ok",
			Query:  map[string]string{"k1": "v1"},
		},
		r,
	)
}

func TestPost(t *testing.T) {
	url, closer := requesttest.Serve("/post", `{"status":"ok"}`)
	defer closer()

	r, err := Post(url, nil)
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

	r, err := Post(url, nil)
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

	r := new(Reponse)
	err := PostJSON(url, nil, r)
	assert.NoError(t, err)
	assert.Equal(
		t,
		&Reponse{
			Status: "ok",
		},
		r,
	)
}

func TestPostJSONQuery(t *testing.T) {
	url, closer := requesttest.Serve("/post", `{"status":"ok","query":{"k1":"v1"}}`)
	defer closer()

	r := new(Reponse)
	err := PostJSON(url, nil, r)
	assert.NoError(t, err)
	assert.Equal(
		t,
		&Reponse{
			Status: "ok",
			Query:  map[string]string{"k1": "v1"},
		},
		r,
	)
}
