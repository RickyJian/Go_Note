# Select

* case 指定一個通訊(某個 channel 的收發操作)與相關的陳述區塊
* select 會等到每個通訊就緒就執行相對應的 case
* 若多個 case 達成，select 會隨機挑一個

```go

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

```

