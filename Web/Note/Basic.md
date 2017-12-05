# Go Web Basic

* `Go` 提供完整的 `net/http` 封包
* 直接監聽 `tcp` 通訊埠

```go

// 建立 Web 伺服器

package main 

import (
    "fmt"
    "net/http"
    "strings"
    "log"
)

func sayHelloName(w http.ResponseWriter , r *http.Request){
    r.ParseForm() //解析參數
    fmt.Println(r.Form) // 輸出伺服器資訊
    fmt.Println("path",r.URL.Path)
    fmt.Println("scheme",r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k , v := ranger.Form {
        fmt.Println("key:",k)
        fmt.Println("val:",stings.Join(v,""))
    }
    fmt.Fprintf(w,"Hello") // 將資料輸出到用戶端
}

func main (){
    http.HandleFunc("/",sayHelloName) // 設定存取路由
    err := http.ListenAndServe(":9090" , nil) // 設定 port
    if err != nil {
        log.Fatal("ListenAndServe: "err)
    }
}

```

## Go Web 工作概念

| 名稱 | 說明 |
| ----- | ----- | 
| Request | 使用者請求資訊，用來解析使用者請求資訊 |
| Response | Server 回饋給 Client 的資訊 |
| Conn | 使用者請求連線 |
| Handler | Request & Response 處理邏輯 |