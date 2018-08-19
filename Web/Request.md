# Request

`net/http` 的 `http.Request`

## Header 

* 獲取 Header 參數
* map 型態

```go

package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
    // Header 全參數
    h := r.Header
    // Header 指定參數(Accept-Encoding)，回傳 string slice 
	h2 := r.Header["Accept-Encoding"]
    // Header 指定參數(Accept-Encoding)，回傳 string
	h3 := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, h, h2, h3)
}

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	mux.HandleFunc("/headers", headers)
	server.ListenAndServe()
}


```

## Body

讀取 request body

```go

package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	// 獲取 body 長度
	len := r.ContentLength
	body := make([]byte, len)
	// 讀取 body 內容
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	mux.HandleFunc("/body", body)
	server.ListenAndServe()
}

```