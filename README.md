# request

Simple HTTP library.

## `GET`

String response
```go
r, err := request.Get("https://httpbin.org/get")
```

JSON response
```go
r := new(BinResponse)
err := request.GetJSON("https://httpbin.org/get", r)
```
