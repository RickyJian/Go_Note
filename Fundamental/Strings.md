# 字串工具

```go

import "strings"

```

## Contains

判斷字串內是否包含，回傳型態布林(bool)

```go

fmt.Println(strings.Contains("seafood", "foo")) // 回傳 true

```

## Split

分割字串，回傳型態[]string

```go

fmt.Println(strings.Split("a,b,c", ",")) // 回傳 ["a" "b" "c"]

```

## Join

連接字串，回傳型態string

```go

s := []string{"apple","peach","banana"}
// 連接 '、' 進 s，輸出：apple、peach、banana
fmt.Println(strings.Join(s, "、"))

```

## Trim

刪除前後字，參數 1 為原字串，參數 2 為要刪除的字串

```go

fmt.Println(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))

```

## TrimSpace

刪除前後空白

```go

fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))

```

## TrimLeft

刪除字串最左方的字，會追蹤

```go

fmt.Println(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "¡")) // 回傳 Hello, Gophers!!!

```

## TrimPrefix

刪除字串前綴字，不會追蹤

```go

fmt.Println(strings.TrimPrefix("¡¡¡Hello, Gophers!!!", "¡")) // 回傳 ¡¡Hello, Gophers!!!

```

## TrimRight

刪除字串最右方的字，會追蹤

```go

fmt.Println(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!")) // 回傳¡¡¡Hello, Gophers

```

## TrimSuffix

刪除字串後綴字，不會追蹤

```go

fmt.Println(strings.TrimSuffix("¡¡¡Hello, Gophers!!!", "!")) // 回傳 ¡¡¡Hello, Gophers!!

```

## Title

將字串中單字開頭轉大寫

```go

fmt.Println(strings.Title("hello everyone."))   // 回傳 Hello Everyone.

```

## ToLower

字串轉小寫

```go

fmt.Println(strings.ToLower("HELLO EVERYONE."))   // 回傳 hello everyone.

```

## ToUpper

字串轉大寫

```go

fmt.Println(strings.ToUpper("hello everyone."))   // 回傳 HELLO EVERYONE.

```
