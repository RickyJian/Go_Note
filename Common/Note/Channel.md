# 並形程式設計 (Channel)

* 是 `gorutine` 之間的連線
* 元素型別，元素型別為 int 的 channel 型別寫做 chan int
* 同型別兩個 channel 可以做比較
* channel 零值為 `nil`
* 運算子：`<-`

```go

    ch := make(chan int) // 建立無緩衝 cahnnel
    ch := make (chan int , 3) // 建立有緩衝 cahnnel，容量為 3

    ch <- x // 發送陳述
    x = <- ch // 指派陳述中的接收運算式
    <- ch // 接收運算式，拋棄結果

    close(ch) // 關閉 channel

```

> channel 是 make 所建構資料結構的參考。複製 channel 或當作函式參數時會複製參考 

## 通訊

發送與接送兩個基本操作的通稱。發送陳述將值從一個goroutinee透過channel傳送到另一個執行相對應接收運算式的goroutine<br>
使用 `<-` 運算子 
