# Request

`net/http` 的 `http.Request` 方法

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
    // Header 指定參數(Accept-Encoding)
	h2 := r.Header["Accept-Encoding"]
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