# 函式

* 包裝陳述式，對使用者隱藏實作細節
* 零值為nil
* 不可比較
* 在套件層級宣告

```go

func main (parameter-list) (result-list){
    do something
}

func Count (no int) int

```

> 未定義函數內容代表由其他語言實作

## 不具名函式

* 在運算式中寫出函式
* func 關鍵字後沒有名稱
* 內層函式可以存取與修改外層函式

> 閉包(closure)：是參照了自由變數的函式。這個被參照的自由變數將和這個函式一同存在，即使已經離開了創造它的環境也不例外。

```go

// 回傳一個函式，型別為 func() int ，為不具名函式
func squares() func() int {
    var x int
    // 回傳平方數
    return func() int{
        x++
        return x * x
    }
}

```

## 參數列 (parameter-list)

* 名稱+型別
* 函式也可當參數

```go

func add (x int , y int) int {}
func sub (x , y int) int {}
func first (x int , _ int) int {}
func zero (int , int) int {}

func funcToBeParameter  (c int , post func(n int)){}

```

### 參數

* 函式內的區域變數，初始由呼叫提供
* 傳值，接收參數的拷貝，修改拷貝不影響呼叫方，帶有參考(例：指標、slice、map、channel、func)會間接影響傳入的參考

### 型別

又稱 signature

## 回傳值(result-list)

指定函式的回傳的值與型別，結果可以具名

### 多回傳值

* 回傳一個以上結果
* 多直呼叫的結果本身可以是呼叫其他多值函式的結果
* 函示有具名結果，return 陳述的運算子可以省略，稱為 僅回傳 (bare return)

```go

func add (x int , y int) (int , int) {}


```

bare return

```go

func operate (x int , y int )(plus , minus , multiple , devide){
    plus := x + y
    minus := x - y
    multiple := x * y
    devide := x / y

    return
}


```


## 遞迴(recursion)

直接或間接呼叫自身function
