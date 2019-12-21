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
payload := map[string][]string{"k1": []string{"v1"}, "k2": []string{"v2", "v3"}}
r, err := request.Get("https://httpbin.org/get", payload)
```

### JSON response

[No query parameters](examples/getjson.go)
```go
r := new(Response)
err := request.GetJSON("https://httpbin.org/get", nil, r)
```

[With query parameters](examples/getjsonparams.go)
```go
r := new(Response)
payload := map[string][]string{"k1": []string{"v1"}, "k2": []string{"v2", "v3"}}
err := request.GetJSON("https://httpbin.org/get", payload, r)
```

## `POST`

### String response

[](examples/post.go)
```go
payload := map[string][]string{"k1": []string{"v1"}, "k2": []string{"v2", "v3"}}
r, err := request.Post("https://httpbin.org/get", payload)
```

### JSON response

[](examples/postjson.go)
```go
r := new(Response)
payload := map[string][]string{"k1": []string{"v1"}, "k2": []string{"v2", "v3"}}
err := request.PostJSON("https://httpbin.org/get", payload, r)
```
