# cookie

* 儲存在 client 端的 browser 裡
* 時間短暫
* 克服 HTTP 的無狀態性

```go

// Cookie struct

type Cookie struct {
	Name       string       // cookie 名稱，能隨意設置
	Value      string       // 儲存在 browser 裡的唯一ID
	Path       string       // 選用
	Domain     string       // 選用
	Expires    time.Time    // 若設置 cookie ， cookie 會在過期後除，若沒設置 cookie 會在 browser 關閉時移除，選用
	RowExpires string       // ，選用
	MaxAge     int          // 指定 cookie 在 browser 存在秒數
	Secure     bool
	HttpOnly   bool         // 設為 true 時 cookie 只能通過 HTTP 或 HTTPS 訪問
	Raw        string
	Unparsed   []string
}

```

## cookie 創建

將 cookie 加在回應 header 裡

```go

package main

import (
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first",
		Value:    "first",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second",
		Value:    "second",
		HttpOnly: true,
	}
	c3 := http.Cookie{
		Name:     "third",
		Value:    "third",
		HttpOnly: true,
	}
	// 設置 Cookie
	w.Header().Set("Set-Cookie", c1.String())
	// 增加 Cookie
	w.Header().Add("Set-Cookie", c2.String())
	// 設置 Cookie
	http.SetCookie(w, &c3)
}

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	mux.HandleFunc("/setcookie", setCookie)
	server.ListenAndServe()
}

```

## cookie 讀取

```go

package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first",
		Value:    "first",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second",
		Value:    "second",
		HttpOnly: true,
	}
	c3 := http.Cookie{
		Name:     "third",
		Value:    "third",
		HttpOnly: true,
	}
	// 設置 Cookie
	w.Header().Set("Set-Cookie", c1.String())
	// 增加 Cookie
	w.Header().Add("Set-Cookie", c2.String())
	// 設置 Cookie
	http.SetCookie(w, &c3)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	// 回傳 slice
	first := r.Header["Cookie"]
	// 回傳 slice
	cookies := r.Cookies()
	// 回傳指定 cookie
	second, _ := r.Cookie("first")
	fmt.Fprintln(w, first)
	fmt.Fprintln(w, cookies)
	fmt.Fprintln(w, second)
}

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	mux.HandleFunc("/setcookie", setCookie)
	mux.HandleFunc("/getcookie", getCookie)
	server.ListenAndServe()
}


```