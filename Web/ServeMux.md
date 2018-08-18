# 處理器

將 url 綁定至 處理器

## DefaultServeMux

* 預設路由器
* 實作 ServeMux Struct 與 Handler
* 是個處理器

```go

import (
	"net/http"
)

func mian (){
	http.ListenAndServe("",nil)
}

```

## NewServeMux

自訂路由器

```go

package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // 宣告自訂路由器
	mux.HandleFunc("/", index) // 當 url 為 / 導到 index func
	mux.HandleFunc("/hello", hello) // 當 url 為 /hello 導到 hello func
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Index Page")
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Hello Page")
}


```