package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Get issues a GET request to a given URL address and formats the response in
// string.
//
// For example:
//  // No query parameters
//  r, err := request.Get("http://localhost:8000", nil)
//
//  // With query parameters
//  r, err := request.Get("http://localhost:8000",  map[string]string{"k1": "v1", "k2": "v2"})
func Get(address string, params map[string]string) (string, error) {
	res, err := get(address, params)
	if err != nil {
		return "", err
	}

	raw, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}

	return string(raw), nil
}

// GetJSON issues a GET request to a given URL address and formats the response
// in JSON.
//
// For example:
//  type Response struct {
//      Status string `json:"status"`
//  }
//
//  // No query parameters
//  r := new(Response)
//  err := request.GetJSON("http://localhost:8000", nil, r)
//
//  // With query parameters
//  r := new(Response)
//  err := request.GetJSON("http://localhost:8000", map[string]string{"k1": "v1", "k2": "v2"}, r)
func GetJSON(address string, params map[string]string, scheme interface{}) error {
	res, err := get(address, params)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&scheme); err != nil {
		return err
	}

	return nil
}

// Post issues a POST request to a given URL address and formats the response
// in string.
//
// For example:
//  // No post body
//  r, err := request.Post("http://localhost:8000", nil)
//
//  // With post body
//  r, err := request.Post("http://localhost:8000", map[string]string{"k1": "v1", "k2": "v2"})
func Post(address string, body map[string]string) (string, error) {
	res, err := post(address, body)
	if err != nil {
		return "", err
	}

	raw, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}

	return string(raw), nil
}

// PostJSON issues a POST request to a given URL address and formats the response
// in JSON.
//
// For example:
//  type Response struct {
//      Status string `json:"status"`
//  }
//
//  // No post body
//  r, err := request.PostJSON("http://localhost:8000", nil)
//
//  // With post body
//  r, err := request.PostJSON("http://localhost:8000", map[string]string{"k1": "v1", "k2": "v2"})
func PostJSON(address string, body map[string]string, scheme interface{}) error {
	res, err := post(address, body)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&scheme); err != nil {
		return err
	}

	return nil
}

func get(address string, params map[string]string) (*http.Response, error) {
	u, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func post(address string, body map[string]string) (*http.Response, error) {
	q := url.Values{}
	for k, v := range body {
		q.Set(k, v)
	}

	res, err := http.PostForm(address, q)
	if err != nil {
		return nil, err
	}

	return res, nil
}
