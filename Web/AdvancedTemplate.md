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


