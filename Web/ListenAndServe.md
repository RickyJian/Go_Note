# ListenAndServe

* `Go` 提供完整的 `net/http` 封包
* 直接監聽 `tcp` 通訊埠
* 默認 port：80
* 處理器參數為 nil ，使用默認 DefaultServeMux

```go

// 方法一

import (
	"net/http"
)

func mian (){
	http.ListenAndServe("",nil)
}

```

```go

// 方法二

import (
	"net/http"
)

func mian (){
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServe("",nil)
}

```

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
