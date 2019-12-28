package request

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Options is the config options for a request.
type Options struct {
	URL     string
	Payload map[string]string
	Auth    map[string]string
}

// Get issues a GET to the specified URL
func Get(opts *Options) (*http.Response, error) {
	return request(opts, "GET")
}

// GetJSON issues a GET to the specified URL and marshals the output in JSON.
func GetJSON(opts *Options, output interface{}) error {
	resp, err := Get(opts)
	if err != nil {
		return err
	}

	return decode(resp, output)
}

// Post issues a POST to the specified URL
func Post(opts *Options) (*http.Response, error) {
	return request(opts, "POST")
}

// PostJSON issues a POST to the specified URL and marshals the output in JSON.
func PostJSON(opts *Options, output interface{}) error {
	resp, err := Post(opts)
	if err != nil {
		return err
	}

	return decode(resp, output)
}

func request(opts *Options, method string) (*http.Response, error) {
	var body io.Reader

	u, err := url.Parse(opts.URL)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	for k, v := range opts.Payload {
		q.Set(k, v)
	}

	if method == "GET" {
		u.RawQuery = q.Encode()
	} else if method == "POST" {
		body = strings.NewReader(q.Encode())
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for username, password := range opts.Auth {
		req.SetBasicAuth(username, password)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func decode(resp *http.Response, output interface{}) error {
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&output); err != nil {
		return err
	}

	return nil
}
