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
//  r, err := request.Get("https://httpbin.org/get", nil)
//
//  // With query parameters
//  payload := map[string][]string{"k1": []string{"v1"}, "k2": []string{"v2", "v3"}}
//  r, err := request.Get("https://httpbin.org/get", payload)
func Get(address string, params map[string][]string) (string, error) {
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
func GetJSON(address string, params map[string][]string, scheme interface{}) error {
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
func Post(address string, body map[string][]string) (string, error) {
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
func PostJSON(address string, body map[string][]string, scheme interface{}) error {
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

func get(address string, params map[string][]string) (*http.Response, error) {
	u, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for k, values := range params {
		for _, v := range values {
			q.Add(k, v)
		}
	}

	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	return res, nil
}

func post(address string, body map[string][]string) (*http.Response, error) {
	q := url.Values{}
	for k, values := range body {
		for _, v := range values {
			q.Add(k, v)
		}
	}

	res, err := http.PostForm(address, q)
	if err != nil {
		return nil, err
	}

	return res, nil
}
