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

### Format

時間格式設定

```go

t := time.Now()
// 一般，輸出：2018-11-18 16:05:41.5624879 +0800 CST m=+0.006980100
fmt.Println("一般：\t" + t.String())
// YYYY-MM-DD，輸出：2018-11-18
fmt.Println("YYYY-MM-DD：\t" + t.Format("2006-01-02"))
// 自製 Format，輸出：2018#11#18
fmt.Println("YYYY#MM#DD：\t" + t.Format("2006#01#02"))
// YYYY-MM-DD hh:mm:ss，輸出：2018-11-18 16:10:14
fmt.Println("YYYY-MM-DD hh:mm:ss：\t" + t.Format("2006-01-02 15:04:05"))
// YYYY-MM-DD hh:mm:ss，輸出：2018-11-18 16:10:14 PM
fmt.Println("YYYY-MM-DD hh:mm:ss：\t" + t.Format("2006-01-02 03:04:05 PM"))
// YYYY-MM-DD hh:mm:ss，輸出：2018-11-18 16:10:14.441385
fmt.Println("YYYY-MM-DD hh:mm:ss：\t" + t.Format("2006-01-02 15:04:05.000000"))
// YYYY-MM-DD hh:mm:ss，輸出：2018-11-18 16:10:14.441385100
fmt.Println("YYYY-MM-DD hh:mm:ss：\t" + t.Format("2006-01-02 15:04:05.000000000"))
// YYYY-MM-DD 月份文字輸出，輸出：2018-November-18
fmt.Println("YYYY-MM-DD：\t" + t.Format("2006-January-02"))
// YYYY-MM-DD 月份文字輸出，輸出：2018-Nov-18
fmt.Println("YYYY-MM-DD：\t" + t.Format("2006-Jan-02"))
// YYYY-MM-DD 星期輸出，輸出：2018-11-18 Sunday
fmt.Println("YYYY-MM-DD：\t" + t.Format("2006-01-02 Monday"))
// YYYY-MM-DD 星期輸出，輸出：2018-11-18 Sun
fmt.Println("YYYY-MM-DD：\t" + t.Format("2006-01-02 Mon"))

```

### 時區設定

```go

t := time.Now()

// local
fmt.Println(t.Location())

// 時區設定，地區時間設定，參數為位置，若位置為 nil 則會產生 error
location, err := time.LoadLocation("Asia/Taipei")
t = t.In(location)

if err != nil {
    fmt.Println(err)
}

// 位置 :  Asia/Taipei  時間 :  2018-11-18 15:39:38.7503566 +0800 CST
fmt.Println("位置 : ", location, " 時間 : ", t)

// 時區設定，時間格式設定，參數為位置，若位置為 nil 則會產生 error
location, err := time.LoadLocation("UTC")
t = t.In(location)

if err != nil {
    fmt.Println(err)
}

// 位置 :  UTC  時間 :  2018-11-18 07:48:04.7495329 +0000 UTC
fmt.Println("格式 : ", UTC, " 時間 : ", t)

```

### 抓取日期&時間

```go

t := time.Now()
// 獲取 年
fmt.Println(t.Year())
// 獲取 年月日
fmt.Println(t.Date())
// 獲取 月
fmt.Println(t.Month())
// 獲取 日
fmt.Println(t.Day())
// 獲取 時
fmt.Println(t.Hour())
// 獲取 分
fmt.Println(t.Minute())
// 獲取 秒
fmt.Println(t.Second())

```

### 日期運算

```go

t := time.Now()
// 增加 年月日，參數一：年，參數二：月，參數三：日
p := t.AddDate(1, 1, 1)
fmt.Println(p.String())

// 增加 時間
p = t.Add(10 * time.Hour)
fmt.Println(p.String())
p = t.Add(10 * time.Minute)
fmt.Println(p.String())
p = t.Add(10 * time.Second)
fmt.Println(p.String())
p = t.Add(10 * time.Millisecond)
fmt.Println(p.String())
p = t.Add(10 * time.Microsecond)
fmt.Println(p.String())
p = t.Add(10 * time.Nanosecond)
fmt.Println(p.String())

```

### 日期比較

```go

t := time.Now()
p := time.Now().AddDate(-1, -1, -1)
// 比較 t 和 p 之間所相差的時間
diff := t.Sub(p)
fmt.Println(diff)

```