# 表單(Form)處理

## 表單輸入

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