# 互斥鎖

## sync.Mutex

```go

import "sync"

var (
    mu sync.Mutex
    balance int
)

func Deposite(amount int ){
    // 取得互斥鎖，會阻斷另一個 goroutine 直到 unlock
    mu.Lock()
    balance = balance + amount
    mu.Unlock()
}

func Balance() int {
    mu.Lock()
    b := balance
    mu.Unlock()
    return b
}

```

## sync.RWMutex

允許讀取並行寫入獨佔


```go

import "sync"

var (
    mu sync.RWMutex
    balance int
)

func Deposite(amount int ){
    // 取得互斥鎖，會阻斷另一個 goroutine 直到 unlock
    mu.Lock()
    defer mu.Unlock()
    balance = balance + amount
}

func Balance() int {
    // 取得讀取鎖
    mu.RLock()
    defer mu.RUnlock()
    return balance
}

```
