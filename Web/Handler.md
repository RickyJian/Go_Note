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
* 透過 http.Handle 函數將處理器綁定至 DefaultServeMux

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

> http.HandleFunc("url",func)：把函式 f 轉換成帶有方法 f 的 Handler

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

## 串聯多個處理器和處理函式

常用於 LOG紀錄 、 安全檢查 和錯誤處理

### 處理器

```go


type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HelloHandler")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler Function call Log----" + name)
		h.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	hello := HelloHandler{}
	http.Handle("/hello", log(&hello))
	server.ListenAndServe()
}



```

### 處理函式

```go

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HelloHandleFunc")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler Function call Log----" + name)
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", log(hello))
	server.ListenAndServe()
}


```

> 這中串接方式又稱`管道處理(pipeline processing)`