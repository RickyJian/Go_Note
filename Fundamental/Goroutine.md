# 並形程式設計 (goroutine)

* 並行執行的活動稱為 goroutine
* main函式為一個 goroutine ，又稱主 goroutine
* goroutine 無法由其他 goroutine 以程式終止

```go

// 建構新的 Goroutine
go f() 

// 執行匿名函式
go func (){
    fmt.Println("Goroutine")
}()

```

## 參數傳遞

```go

names := []string{"john", "dora", "jack", "doris", "bob"}
for _, name := range names {
    go func(who string) {
        fmt.Println(who)
    }(name)
}


```

## 暫停

### Sleep

暫停當前的 Goroutine

```go

// 暫停 1 毫秒
time.Sleep(time.Millisecond)

```

### Gosched

若當前 Goroutine 中尚有其餘 Goroutine 在執行時，則不會暫停當前 Goroutine

```go

go println("New GO")
fmt.Println("Current GO")
runtime.Gosched()

```


## sync.WaitGroup

goroutine 計數器

```go

func main() {
	demo()
}

func demo() {
	sizes := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sizes <- i
		}(i)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int
	for size := range sizes {
		total += size
		fmt.Println(size)
	}
}


```

### Add()

增加計數器

### Done()

減少計數器，等同於 Add(-1)，搭配 defer 使用能確保在 panic 發生時也會繼續遞減

### Wait()

等待 goroutine 結束