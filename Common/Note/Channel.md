# 並形程式設計 (channel)

* 是 `gorutine` 之間的連線
* 元素型別，元素型別為 int 的 channel 型別寫做 chan int
* 同型別兩個 channel 可以做比較
* channel 零值為 `nil`
* 運算子：`<-`
* 同型別 channel 可比較

## 建立及關閉

用 make 建構 channel，用close(ch)關閉 channel

```go

    ch := make(chan int) // 建立無緩衝 cahnnel
    ch2 := make (chan int , 3) // 建立有緩衝 cahnnel，容量為 3

    close(ch) // 關閉 channel

```

> channel 是 make 所建構資料結構的參考。複製 channel 或當作函式參數時會複製參考

## 通訊

發送與接送兩個基本操作的通稱。發送陳述將值從一個goroutine透過channel傳送到另一個執行相對應接收運算式的goroutine
<br>
使用 `<-` 運算子 


```go

    ch := make(chan int)

    ch <- x // 發送陳述
    x = <- ch // 指派陳述中的接收運算式
    <- ch // 接收運算式，拋棄結果

```

> channel 關閉後發送操作會產生 pannic

## 管道(pipeline)

一個 goroutine 的輸出可做為另一個 goroutine 的輸入

```go

func main() {
	na := make(chan int)
	sq := make(chan int)
	go func() {
		for x := 0; x < 100; x++ {
			na <- x
		}
		fmt.Println("na")
		close(na)
	}()
	go func() {
		for x := range na {
			sq <- x * x
		}
		close(sq)
	}()
	for x := range sq {
		fmt.Println(x)
	}
}

```

```go

// 檢查 channel 是否已抽乾或關閉

    x , ok := <- na

```

## 單向 channel 型別

防止 channel 接送與發送勿用，單向發送(chan<- )，單向接收(<-chan )

```go

func counter (out chan<- int){
    for x := 0 ; x <100 ; x++ {
        out <- x
    }
    close(out)
}

func squarer (out chan<- int , in <-chan int){
    for v := range in {
        out <- v * v
    }
    close(out)
}


func pointer (in <-chan int){
    for v := range in {
        fmt.Println(v)
    }
}

func main (){
    na := make(chan int)
    sq := make(chan int)
    go counter(na)
    go squarer(sq,na)
    pointer(sq)

}

```

## 無緩衝 channel

發送操作會阻斷發送方 goroutine 直到另一個 goroutine 在同一個 channel 執行相對應的接收<br>
接收操作則接收方會阻斷直到另一個 goroutine 在同一個 channel 上執行相對應的發送<br>
又稱 goroutine 同步化

```go

    ch := make(chan int) // 建立無緩衝 cahnnel

```

## 有緩衝 channel

buffer 空時，發送操作會將元素插入佇列的後方，接收操作會阻斷直到另一方 goroutine 發送值<br>
buffer 滿時，發送操作會阻斷它的 goroutine 直到另一方 goroutine 的接收產生空間，接收操作會將前面元素移除<br>
buffer 未滿時，接收及發送不會被阻斷


```go

    ch := make (chan int , 3) // 建立有緩衝 cahnnel，容量為 3
    cap(ch) // 檢查 channel 緩衝容量
    len(ch) // 檢查 channel 目前在緩衝中的元素數量

```

