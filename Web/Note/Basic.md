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

> ListenAndServe：初始化 Server 物件，呼叫 net.Listen("tcp",addr)<br>
> srv.Listen(net.Listener)函數，處理接受 Client 請求資訊<br>
> 使用者每一次請求都是由新的 goroutine 去福，並補互相影響


## Go Web 工作概念

| 名稱 | 說明 |
| ----- | ----- | 
| Request | 使用者請求資訊，用來解析使用者請求資訊 |
| Response | Server 回饋給 Client 的資訊 |
| Conn | 使用者請求連線 |
| Handler | Request & Response 處理邏輯 |


## http 封包執行機制

1. 監聽 Socket 等待Client請求
1. 接受 Client Request 
1. 處理 Client Request，並交給 Handler 處理

## Conn

用戶每一次請求都會建立 Conn ，這裡 Conn 儲存了該次請求的資訊，然後再傳遞給 handler ，該 handler 中便可以讀取對應的 handler 資訊，確保請求獨立性

```go

// 等待用戶請求

c, err := srv.newConn(rw)
if err != nil {
    continue
} 
go c.serve()

```
## ServeMux

```go

package main 

import (
    "fmt"
    "net/http"
)

type MyMux struct {

}

func (p *MyMux) ServeHTTP(w http.ResponseWriter , r *http.Request){
    if r.URL.Path == "/" {
        sayHelloName(w,r)
        return 
    }
    http.NotFound(w,r)
    return
}

func sayHelloName(w http.ResponseWriter , r *http.Request){
    fmt.Fprintf(w,"Hello My Route!!^^")
}
func main (){
    mux := &MyMux{}
    http.ListenAndServe(":9090",mux)
}

```