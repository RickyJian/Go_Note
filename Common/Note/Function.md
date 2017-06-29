# 函式

包裝陳述式，對使用者隱藏實作細節

```go

func main (parameter-list) (result-list){
    do something
}

func Count (no int) int

```

> 未定義函數內容代表由其他語言實作

## 參數列 (parameter-list)

名稱+型別

```go

func add (x int , y int) int {}
func sub (x , y int) int {}
func first (x int , _ int) int {}
func zero (int , int) int {}

```

### 參數

* 函式內的區域變數，初始由呼叫提供
* 傳值，接收參數的拷貝，修改拷貝不影響呼叫方，帶有參考(例：指標、slice、map、channel、func)會間接影響傳入的參考



### 型別

又稱 signature

## 回傳值(result-list)

指定函式的回傳的值與型別，結果可以具名

## 遞迴