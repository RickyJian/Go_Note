# HTTP2

```go

package main

import (
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
    http.HandleFunc("/hello", nil)
    // 開啟 HTTP2
	http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServe()
}


```