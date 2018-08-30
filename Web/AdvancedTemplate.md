## 進階模板

## 參數

在動作中設置變數，須使用`$`開頭

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
    {{ range $key , $value := . }}
        鍵：{{ $key }}
        值：{{ $value }}
    {{ end }}
</body>
</html>

```

## 管道(pipeline)

有順序的將參數、函式、方法串接起來

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
    {{ 12.3456789 | printf "%.2f"}}
</body>
</html>

```

## 函式

* 只能返回一個值和錯誤
* template.FuncMap：綁定函式名稱(前端)與函式(後端)
* t.Funcs：將函式綁定模板中

### 後端

```go

package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func templateExample(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("index.html").Funcs(funcMap)
	t, _ = t.ParseFiles("index.html")
	t.Execute(w, time.Now())
}

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	mux.HandleFunc("/template", templateExample)
	server.ListenAndServe()
}

```

### 前端

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
    {{ . | fdate }}
    <br>
    {{  fdate . }}
</body>
</html>

```