# ListenAndServe

* `Go` 提供完整的 `net/http` 封包
* 直接監聽 `tcp` 通訊埠


```go

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello World")
}


```

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
