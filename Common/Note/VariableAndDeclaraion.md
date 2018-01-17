# 宣告

* 以字母底線開頭
* 駝峰字(camel case)
* 原縮寫字維持原大小

> 四種宣告 ： var 、 const 、 type 、 func

## 區域 (local)

* 函式內宣告
* 傾向短名稱
* 內層宣告(區域變數)會將外層宣告(全域變數)`遮蔽`


## 全域

* 函式外宣告
* 所屬的套件的全部檔案皆能看見
* 第一個字母決定他的可見度
* 名稱以大寫字母開頭為`匯出(exported)`
* 套件名稱永遠是小寫
* 範圍愈大名稱愈長

## 變數 (Variable) 宣告

值的儲存體，又稱有位置的值

### 一般

保留給須明確設定參數與型別，或之後被指派值且初值不重要的變數

> var name type = expression <br>
> var name type 

```go

var i, j, k int  // 三個變數都是 int 型態
var b, f, s = true , 2.3 , "four"  // boolean , float , string

```

### 簡短宣告

* 用於初始化大部分的區域變數，且不能用在全域變數，在使用時至少要宣告一個新變數
* 需要知道範圍

>  name := expression

```go

i := 0
i , j := 0,1

```

> `:=` 宣告 <br>
> `=`  指派

```go

var cwd string

func init (){
    cwd , err := os.Getwd() //編譯錯誤，cwd 未使用
    if err != nil{
        log.FatalF("os.Getwd failed: %v",err)
    }
}

```

```go

var cwd string

func init (){    
    cwd , err := os.Getwd() //編譯錯誤，cwd 未使用
    if err != nil{
        log.FatalF("os.Getwd failed: %v",err)
    }
}

```

```go

var cwd string

func init (){
    cwd , err := os.Getwd() //雖編譯成功，但還是錯誤，編譯成功原因是因為， log.Fatalf("os.Getwd failed : %s", cwd)，它會抓到 cwd 的錯誤
    if err != nil{
        log.FatalF("os.Getwd failed: %v",err)
    }
    log.Fatalf("os.Getwd failed : %s", cwd) 

```

>  log.Fatalf() ： 會呼叫 os.Exit(1)

#### type

可以省略，若省略會運算式自行決定它的型態

#### expression

可以省略，若省略初值會是該`type`的零值

> 零值機制 ： 確保變數總是保持正確的值，`Go`沒有未定義的值。

#### 變數生命週期

* 套件層級：整個執行期間
* 區域變數：動態，當不用就會被回收

## 型別宣告

* 自訂義型別
* 不能直接進行比較及並行運算

> type name underlying-type

```go

type Celsius float64

```

### name

新具名型別

### underlying-type

底層型別

## 指派 `=`

```go

x = 1       // 具名變數 
*p = true   // 間接變數
person.name // struct 欄位
count[x] = count[x] * scale //陣列 slice map 元素

```

### 指派運算

```go

count[x] *= scale  // count[x] = count[x] * scale

v++ // v = v + 1

v-- // v = v - 1

```

### 資料組指派

可以一次指派多個變數

```go

x , y = y , x

```

## 轉型

* 底層型別相同
* 指向相同底層型別的變數指標
* 有可能改變值的表示

> T(x)：將x值轉換成T型別

### Pointer 轉型

```go

*string(v) // 先將v值轉成string在取得指標

(*string) (v) // 把變數v值轉成指標欸型*string值

```
