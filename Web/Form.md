# 表單(Form)處理

## ParseForm()

解析 Form 表單

### 前端

```html

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
</body>
    <form action="http://127.0.0.1:8080/process" method="post">
        <input type="text" name="username" id="username">
        <input type="submit" value="送出">
    </form>
</html>

```

### 後端

```go

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// 解析 form
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, "index.html")
}

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	mux.HandleFunc("/process", process)
	mux.HandleFunc("/index", index)
	server.ListenAndServe()
}

```

## 表單輸入

> content-type：決定 POST 請求用何種格式傳送，由 form tag 的 enctype 指定，預設為 `application/x-www-form-urlencoded`

```html
<!-- 前端 -->
<html>
    <head>
        <title>
        </title>
    </head>
    <body>
        <form action =  "http://localhost:9090/Login" method = "post">
            username: <input type = "text" name = "username"> 
            password: <input type = "password" name = "password">
            <input type = "submit" value = "登入"> 
        </form>
    </body>
</html>

```

```go
// 後端
package main 

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)

func sayHelloName (w http.ResponseWriter , r *http.Request){
    r.ParseForm() // 參數解析，無此方法無法獲得 Form 參數
    fmt.Println(r.Form)
    fmt.Println("Path:",r.URL.Path)
    fmt.Println("Scheme:",r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k,v := range r.Form {
        fmt.Println("key:",k)
        fmt.Println("val:",strings.Join(v,""))
    }
    fmt.Fprintf(w,"Hello") // 輸出到用戶端
}

func login (w http.ResponseWriter , r *http.Request){
    fmt.Println("method:",r.Method) // 請求取得方法
    if r.Method == "GET" {
        t , _ := template.ParseFiles("login.html")
        t.Execute(w,nil)
    }else{
        // form 回傳商模處理
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}

func main (){
    http.HandleFunc("/",sayHelloName) // Route 設定
    http.HandleFunc("/Login",login) // Route 設定
    err := http.ListenAndServe(":9090",nil)
    if err != nil {
        log.Fatal("ListenAndServe:",err)
    }

}

```

> r.Form["username"] 也可以用 r.FormValue("username") 代替，但當使用 r.FormValue 時可以不必呼叫 r.ParseForm

-----

模板：[連結](Template.md)