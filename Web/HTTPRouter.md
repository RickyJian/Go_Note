# HTTPRouter

* 第三方 Router 處理器
* 解決 ServeMux URL 無法動態擷取的問題

[連結](https://github.com/julienschmidt/httprouter)

## New()

開起 HTTPRouter

```go

package main 

import "github.com/julienschmidt/httprouter"

func main (){
    httprouter.New()
}

```