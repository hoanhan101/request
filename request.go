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
	u, err := url.Parse(opts.URL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for k, v := range opts.Payload {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()

	return request(opts, "GET", u.String(), nil)
}

// GetJSON issues a GET to the specified URL and marshals the output in JSON.
func GetJSON(opts *Options, output interface{}) error {
	resp, err := Get(opts)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&output); err != nil {
		return err
	}

	return nil
}

// Post issues a POST to the specified URL
func Post(opts *Options) (*http.Response, error) {
	q := url.Values{}
	for k, v := range opts.Payload {
		q.Set(k, v)
	}

	return request(opts, "POST", opts.URL, strings.NewReader(q.Encode()))
}

// PostJSON issues a POST to the specified URL and marshals the output in JSON.
func PostJSON(opts *Options, output interface{}) error {
	resp, err := Post(opts)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&output); err != nil {
		return err
	}

	return nil
}

func request(opts *Options, method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	for username, password := range opts.Auth {
		req.SetBasicAuth(username, password)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
