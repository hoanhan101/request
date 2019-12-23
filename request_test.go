package request

import (
	"testing"

	"github.com/hoanhan101/requesttest"
	"github.com/stretchr/testify/assert"
)

type Reponse struct {
	Status string `json:"status"`
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
