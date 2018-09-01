# 表單(Form)處理

## 表單解析

### ParseForm()

* 解析 Form 表單，包含 URL Query 及 form 內容
* 結構為 Map，key 型態為 string，value 型態為 string slice
* HTML Form enctype 為 application/x-www-form-urlencoded

#### 前端

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
    <form action="http://127.0.0.1:8080/process?gender=male" method="post" enctype = "application/x-www-form-urlencoded">
        <input type="text" name="username" id="username">
        <input type="submit" value="送出">
    </form>
</html>

```

#### 後端

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

### ParseMultipartForm()

* 解析 MultipartForm 表單
* 結構為2個 Map，
    * 第1個 Map：key 型態為 string，value 型態為 string slice，存放一般 Form 內容
    * 第2個 Map：key 型態為 string，value 型態為 string slice，存放 上傳檔案 內容
* HTML Form enctype 為 multipart/form-data

#### 前端

```HTML

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
    <form action="http://127.0.0.1:8080/process" method="post" enctype="multipart/form-data">
        <input type="file" name="uploadFile">
        <input type="submit" value="送出">
    </form>
</html>

```

#### 後端

```go

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// 解析 MultipartForm，上傳容量為 1024
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.MultipartForm)
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

## 表單讀取

### 全

| 方法 | 解析 | 編碼方式 |
| ----- | ----- | ----- |
| Form | URL、Form | URL Encode |
| PostForm | Form | URL Encode |
| MultipartForm | Form | Multipart Encode |

#### Form

解析 URL Query 和 Form 

```go

func process(w http.ResponseWriter, r *http.Request) {
	// 解析 form，解析結果 map[username:[test] gender:[male]]
    r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

```

#### PostForm

解析 Form 

```go

func process(w http.ResponseWriter, r *http.Request) {
	// 解析 form，解析結果 map[username:[test]]
	r.ParseForm()
	fmt.Fprintln(w, r.PostForm)
}

```

#### MultipartForm

解析 Form 

```go

func process(w http.ResponseWriter, r *http.Request) {
	// 解析 MultipartForm，上傳容量為 1024，結果為&{map[] map[uploadFile:[0xc04212e000]]}
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.MultipartForm)
}

```

### 表單值

#### FormValue()

解析 URL Query 和 Form 

```go

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// 讀取 Form username 的值
	username := r.FormValue("username")
	// 讀取 Form email 的值
	email := r.FormValue("email")
	// 讀取 URL Query gender 的值
	gender := r.FormValue("gender")
	fmt.Fprintln(w, username, email, gender)
}

```

#### PostFormValue()

解析 Form 

```go

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// 讀取 Form username 的值
	username := r.PostFormValue("username")
	// 讀取 Form email 的值
	email := r.PostFormValue("email")
	fmt.Fprintln(w, username, email)
}

```

-----

模板：[連結](Template.md)