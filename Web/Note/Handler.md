# Handler Method And Handler Function

## 處理器(Handler Method)

* 一個處理器就是實作 ServeHTTP 方法介面
* 參數一：http.ResponseWriter， 參數二：*http.Request

> ServeHTTP(http.ResponseWriter , *http.Reuqust)


```go


type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "MyHandler")
}

func main() {
    handler := MyHandler{}
	server := http.Server{
        Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}


```

### 多處理器

* 使用多處理器處理不同URL，因此我們會使用 DefaultServeMux 做為處理器。
* 透過 http.handle 函數將處理器綁定至 DefaultServeMux

```go

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HelloHandler")
}

type UserHandler struct{}

func (u *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UserHandler")
}

func main() {
	hello := HelloHandler{}
	user := UserHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.Handle("/hello", &hello)
	http.Handle("/user", &user)
	server.ListenAndServe()
}

```

## 處理函式(Handler Function)

* 與處理器擁有相同行為的函式
* 將一個正確簽名的函數傳換成方法，並將它與 DefaultServeMux 綁定

```go

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UserHandleFunc")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UserHandleFunc")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/user", user)
	server.ListenAndServe()
}

```