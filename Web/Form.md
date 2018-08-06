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

> r.Form["username"] 也可以用 r.FormValue("username") 代替，但當使用 r.FormValue 時可以不必呼叫 r.ParseForm

## 表單驗證

資料後端驗證

### 必填欄位

```go

if len(r.Form["username"][0])==0 {
    // 空處理
}

// or

if len(r.Form.Get("username"))==0 {
    // 空處理
}

```

### 數字驗證

```go

getint , err := strconv.Atoi(r.Form.Get("age"))

if err != nil {
    // 轉型失敗
}

if getint > 50 {
    // 數字邏輯處理
}
// 正則表示式
if m , _ := regxp.MathString("^[0-9]+$" , r.Form.Get("age")); !m {
    return false
}

```

### 下拉選單(select)驗證

```html

<select name = "fruit">
    <option value = "apple"> apple </option>
    <option value = "peer"> peer </option>
    <option value = "orange"> orange </option>
<select> 

```

```go

slice := []string{"apple" , "peer" , "orange"}
for _ , v := range slice {
    if v == r.Form.Get("fruit"){
        return true 
    }
}
return false

```

### 選項按鈕(Radio button)驗證

```html

<input type ="radio" name = "gender" value = "1"> male
<input type ="radio" name = "gender" value = "2"> female

```

```go

slice := []int{1,2}
for _ , v := range slice {
    if v == r.Form.Get("gender"){
        return true 
    }
}
return false

```

### 核取按鈕(checkbox)驗證


```html

<input type ="checkbox" name = "interest" value = "socker"> socker
<input type ="checkbox" name = "interest" value = "baseball"> baseball
<input type ="checkbox" name = "interest" value = "basketball"> basketball

```

```go

slice := []string{"socker","baseball", "basketball"}
check := Slice_diff(r.Form["interest"],slice) 
if check != nil {
    return false
}
return false

```

### 日期時間驗證

```go

t := time.Date(2009, time.November , 10 ,23 ,0,0,0,time.UTC)
fmt.Printf(t.Local())

```

## 預防跨站攻擊(Cross Site Scripting,XSS)

使用 `html/template` 

| 功能 | 說明 |
| ----- | ----- |
| HTMLEscape(w io.Wtiter , b []byte) | 將 b 進行逸出後寫到 w |
| HTMLEscapeString(s String) string | 逸出s之後回傳字串 |
| HTMLEscape(args ...interface{}) string | 支援多參數逸出，並傳回結果字串 |

```go

    fmt.Println("username", template.HTMLEscapeString(r.FormValue("userName")) )

```

## 防止多次送出表單

在 form 裡新增一個時間戳記

```html

<!DOCTYPE html>
<html lang="en">
    <head>
        <title>test</title>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link href="css/style.css" rel="stylesheet">
    </head>
    <body>
        <input type="checkbox" name="interest" value="football">足球
        <input type="checkbox" name="interest" value="basketball">籃球
        <input type="checkbox" name="interest" value="tennis">網球
        使用名稱： <input type="text" name="username" >
        密碼： <input type="password" name="password" >
        <!-- md5 時間戳記 -->
        <input type="hidden" name="token" value="{{.}}">
        <input type="submit" value="登入">
    </body>
</html>

```

```go

func login (w http.ResponseWriter , r *http.Request) {
    if r.Method == "GET" {
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h , strconv.FormatInt(crutime,10))
        token := fmt.Sprintf("%x",h.Sum(nil))
        t , _ := template.ParseFiles("index.html")
        t.Excute(w,token)
    }else{
        token := r.FormValue("token")
        if token != "" {
            // 驗證 token
        }else{
            // 驗證失敗
        }
        
    }
}

```

## 檔案上傳

