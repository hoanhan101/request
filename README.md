# request

[![Godoc](https://godoc.org/github.com/hoanhan101/request?status.svg)](https://godoc.org/github.com/hoanhan101/request)

Simple HTTP library.

## `GET`

String response
```go
// No query parameters
r, err := request.Get("http://localhost:8000", nil)

With query parameters
r, err := request.Get("http://localhost:8000",  map[string]string{"k1": "v1"})
```

JSON response
```go
// No query parameters
r := new(Response)
err := request.GetJSON("http://localhost:8000", nil, r)

// With query parameters
r := new(Response)
err := request.GetJSON("http://localhost:8000", map[string]string{"k1": "v1"}, r)
```

## `POST`

String response
```go
// No post body
r, err := request.Post("http://localhost:8000", nil)

// With post body
r, err := request.Post("http://localhost:8000", map[string]string{"k1": "v1"})
```

JSON response
```go
// No post body
r := new(Response)
err := request.PostJSON("http://localhost:8000", nil, r)

// With post body
r := new(Response)
err := request.PostJSON("http://localhost:8000", map[string]string{"k1": "v1"}, r)
```
