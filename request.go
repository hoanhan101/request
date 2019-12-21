package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

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
