# 互斥鎖



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