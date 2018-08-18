# HTTPS

```go

package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	// cert.pem : ssl 認證
	// key.pem : server 私鑰
	server.ListenAndServeTLS("cert.pem", "key.pem")
}


```