package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Get issues a GET request to a given URL address and formats the response in
// string.
//
// For example:
//  // No query parameters
//  r, err := request.Get("http://localhost:8000", nil, nil)
//
//  // With query parameters
//  r, err := request.Get("http://localhost:8000", map[string]string{"k1": "v1"}, nil)

//  // With authentication
//  r, err := request.Get("http://localhost:8000", map[string]string{"k1": "v1"}, map[string]string{"username":"password"})
func Get(address string, params, auth map[string]string) (string, error) {
	res, err := get(address, params, auth)
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
//  // No query parameters
//  r := new(Response)
//  err := request.GetJSON("http://localhost:8000", nil, nil, r)
//
//  // With query parameters
//  r := new(Response)
//  err := request.GetJSON("http://localhost:8000", map[string]string{"k1": "v1"}, nil, r)

//  // With authentication
//  r := new(Response)
//  err := request.GetJSON("http://localhost:8000", map[string]string{"k1": "v1"}, map[string]string{"username":"password"}, r)
func GetJSON(address string, params, auth map[string]string, scheme interface{}) error {
	res, err := get(address, params, auth)
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
//  r, err := request.Post("http://localhost:8000", nil, nil)
//
//  // With post body
//  r, err := request.Post("http://localhost:8000", map[string]string{"k1": "v1"}, nil)
//
//  // With authentication
//  r, err := request.Post("http://localhost:8000", map[string]string{"k1": "v1"}, map[string]string{"username":"password"})
func Post(address string, body, auth map[string]string) (string, error) {
	res, err := post(address, body, auth)
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
//  // No post body
//  r := new(Response)
//  err := request.PostJSON("http://localhost:8000", nil, nil, r)
//
//  // With post body
//  r := new(Response)
//  err := request.PostJSON("http://localhost:8000", map[string]string{"k1": "v1"}, nil, r)
//
//  // With authentication
//  r := new(Response)
//  err := request.PostJSON("http://localhost:8000", map[string]string{"k1": "v1"}, map[string]string{"username":"password"}, r)
func PostJSON(address string, body, auth map[string]string, scheme interface{}) error {
	res, err := post(address, body, auth)
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

func get(address string, params, auth map[string]string) (*http.Response, error) {
	u, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	for username, password := range auth {
		req.SetBasicAuth(username, password)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func post(address string, body, auth map[string]string) (*http.Response, error) {
	q := url.Values{}
	for k, v := range body {
		q.Set(k, v)
	}

	req, err := http.NewRequest("POST", address, strings.NewReader(q.Encode()))
	if err != nil {
		return nil, err
	}

	for username, password := range auth {
		req.SetBasicAuth(username, password)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil

	// res, err := http.PostForm(address, q)
	// if err != nil {
	// 	return nil, err
	// }

	// return res, nil
}
