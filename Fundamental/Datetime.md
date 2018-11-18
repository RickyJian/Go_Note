# 日期時間處理

```go

import "time"

```

## 函式&方法

### Now

回傳現在時間

```go

t := time.Now()
// 輸出：2018-11-18 15:33:40.1275359 +0800 CST m=+0.006983300
fmt.Println(t)

```

### 時區設定

```go

t := time.Now()

// local
fmt.Println(t.Location())

// 時區設定，參數為位置，若位置為 nil 則會產生 error
location, err := time.LoadLocation("Asia/Taipei")
t = t.In(location)

if err != nil {
    fmt.Println(err)
}

// 位置 :  Asia/Taipei  時間 :  2018-11-18 15:39:38.7503566 +0800 CST
fmt.Println("位置 : ", location, " 時間 : ", t)

```