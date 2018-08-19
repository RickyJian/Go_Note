# Response

`net/http` 的 `http.ResponseWriter`

## Write

將資料寫入 HTTP 主體中

```go

package main

import (
	"net/http"
)

func write(w http.ResponseWriter, r *http.Request) {
	str := "<h1>Hello World</h1>"
	w.Write([]byte(str))
}

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	mux.HandleFunc("/writeexample", write)
	server.ListenAndServe()
}

```

## WriteHeader 

以 HTTP Status 作為參數，並作為 reaponse 回應狀態

```go

package main

import (
	"fmt"
	"net/http"
)

func write(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprintf(w, "404 NOT FOUND")
}

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	mux.HandleFunc("/writeheader", write)
	server.ListenAndServe()
}

```