# 服務靜態檔案

```go

package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
    mux.Handle("/static/", http.StripPrefix("/static/", files))
    // 當 URL 為：http://localhost:8080/static/js/jQuery.js 時，
    // 他會在 public 目錄找到 js/jQuery.js 檔案
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	// index prcess
}


```

## FileServer

指定路徑中靜態檔案的處理器

## Dir

檔案放置目錄

## StripPrefix

移除 URL 前綴字