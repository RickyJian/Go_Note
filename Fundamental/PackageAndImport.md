# PackaAndImport

以 `Package` 開頭，後面接著 `import`

```go

package main 

import "fmt"

func main (){
    fmt.Println("Hello World!!")
}


```

## 套件、檔案

* 原始碼能放在一個或多個 `.go`的檔案中
* 放`package`前為**文件註解**，用來說明套件，每個套件只有一個文件註解
* 大量文件放在`doc.go`中

### package

* 決定該套件由其他套件匯入時的預設識別名稱(套件名稱)
* 路徑獨一無二

### 匯出(exported)

名稱以大寫字母開頭為

### 匯入(import)

* 路徑，字串，獨一無二
* 匯入宣告將短名稱與匯入的套件綁定已透過檔案來參考它的內容
* 沒有用到匯入會造成編譯錯誤  
* 將檔案放置到不同的資料夾內，必須匯入才能使用，若是放在同樣的資料夾內可以彼此共享

> "gopl.io/ch2/tempconv"

```go

// 第一種

import "a/b" //套件必須增加在 GOPATH 中

// 第二種

import (
    "a/b"
    "c"
)

// 第三種 新增別名，若套件名稱相同可用此方法呼叫

import (
    . "d" // 直接呼叫相依套件
    a "a" // 依別名呼叫相依套件
    _ "c" // 指初始套件但不用套件內的任何程式
)


```



