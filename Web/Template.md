# Template

* `text/template` 和 `html/template` 為預設 library
* 模板引擎套過將資料和模板組合成最終 HTML
* 模板又分成兩種
    * 無邏輯模板(logic-less template engine)：只進行字串替換不進行邏輯處理
    * 嵌入邏輯模板(embedded logic template engine)：將程式寫入模板中並做字串替代

## 前端

### 動作(action)

模板中的默認動作須使用`{{`和`}}`包圍

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
    {{ . }}
</body>
</html>

```

### 條件判斷

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
    {{ if . }}
    if condition
    {{ else }}
    else condition
    {{ end }}
</body>
</html>

```

### 迭代

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
    {{ range . }}
        <li>{{ . }}</li>
    {{ end }}
</body>
</html>

```

#### fallback

當迭代為空時，則執行 fallback 結果

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
    {{ range . }}
        <li>{{ . }}</li>
    {{ else }}}
    <li>沒有資料顯示</li>
    {{ end }}
</body>
</html>

```

### 設值


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
    {{ with "apple" }}
        Set Value is {{ . }}
    {{ end }}
</body>
</html>

```

### 模板嵌入

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
    {{ template "t2.html" }}
</body>
</html>

```

#### 設值


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
    {{ template "t2.html" . }}
</body>
</html>

```

### 定義模板名稱並嵌入

須使用 `{{ defined "name" }}`，name 為定義模板名稱

嵌入模板名稱

```html

{{ defined "layout" }}

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    {{ template "content" }}
</body>
</html>

{{ end }}

```

被嵌入模板名稱

```html

{{ defined "content" }}

<div>HELLO</div>

{{ end }}

```

## 後端

### ParseFiles() & Execute()

ParseFiles()：解析模板，當使用者使用此函式時，Go 會創建一個新模板，並指定此文件名稱為新模板名稱

Execute()：將資料綁定模板中，模板在生成 HTML 之後會將該 HTML 傳給 ResponseWriter

```go

package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
    // 新建模板
    t := template.New("index.html")
    // 新建&解析模板
    t, _ := template.ParseFiles("index.html")
    // Must()：當有錯誤會導致 panic 產生
    t, _ := template.Must(template.ParseFiles("index.html"))
    // 將資料綁釘模板
	t.Execute(w, "Hello")
}

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	mux.HandleFunc("/process", process)
	server.ListenAndServe()
}


```

## ParseGlob()

對所有指定到的檔案視為模板

```go

template.ParseGlob("*.html")

```

## ExecuteTemplate()

執行模板集合或其他模板

```go

t , _ := template.ParseFiles("t1.html","t2.html")
t.ExecuteTemplate(w,"t2.html","t2 template")


```