# request

Simple HTTP library.

## `GET`

### String response

[No query parameters](examples/get.go)
```go
r, err := request.Get("https://httpbin.org/get", nil)
```

[With query parameters](examples/getparams.go)
```go
payload := map[string]string{"k1": "v1", "k2": "v2"}
r, err := request.Get("https://httpbin.org/get", payload)
```

### JSON response

[No query parameters](examples/getjson.go)
```go
r := new(BinResponse)
err := request.GetJSON("https://httpbin.org/get", nil, r)
```

[With query parameters](examples/getjsonparams.go)
```go
r := new(BinResponse)
err := request.GetJSON("https://httpbin.org/get", payload, r)
```
