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

## context-aware

根據內容在檔案中的位置做相對應的內容逸出(escape)，可防止XSS攻擊

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
</body>
    {{ . }}
    <div>{{ . }}</div>
    <a href="/{{ . }}">test</a>
    <a href="/?q={{ . }}">test</a>
    <a onclick="f('{{ . }}')">test</a>
</html>

```

#### 逸出(escape)

```HTML

<html style="" lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body style="margin-bottom: 8px !important;">

    &lt;i&gt; Escape &lt;/i&gt;
    <div>&lt;i&gt; Escape &lt;/i&gt;</div>
    <a href="/%3ci%3e%20Escape%20%3c/i%3e">test</a>
    <a href="/?q=%3ci%3e%20Escape%20%3c%2fi%3e">test</a>
    <a onclick="f('\x3ci\x3e Escape \x3c\/i\x3e')">test</a>
</body></html>

```

### 後端

```go

package main

import (
	"html/template"
	"net/http"
)

func templateExample(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	content := "<i> Escape </i>"
	t.Execute(w, content)
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