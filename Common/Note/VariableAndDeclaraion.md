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

## 範圍
- [宣告](#%E5%AE%A3%E5%91%8A)
    - [區域 (local)](#%E5%8D%80%E5%9F%9F-local)
    - [全域](#%E5%85%A8%E5%9F%9F)
    - [變數 (Variable) 宣告](#%E8%AE%8A%E6%95%B8-variable-%E5%AE%A3%E5%91%8A)
        - [一般](#%E4%B8%80%E8%88%AC)
        - [簡短宣告](#%E7%B0%A1%E7%9F%AD%E5%AE%A3%E5%91%8A)
            - [type](#type)
            - [expression](#expression)
            - [變數生命週期](#%E8%AE%8A%E6%95%B8%E7%94%9F%E5%91%BD%E9%80%B1%E6%9C%9F)
    - [型別宣告](#%E5%9E%8B%E5%88%A5%E5%AE%A3%E5%91%8A)
        - [name](#name)
        - [underlying-type](#underlying-type)
    - [指派 `=`](#%E6%8C%87%E6%B4%BE)
        - [指派運算](#%E6%8C%87%E6%B4%BE%E9%81%8B%E7%AE%97)
        - [資料組指派](#%E8%B3%87%E6%96%99%E7%B5%84%E6%8C%87%E6%B4%BE)
    - [範圍](#%E7%AF%84%E5%9C%8D)
        - [語句(syntactic)區塊](#%E8%AA%9E%E5%8F%A5syntactic%E5%8D%80%E5%A1%8A)
        - [詞彙(lexical)區塊](#%E8%A9%9E%E5%BD%99lexical%E5%8D%80%E5%A1%8A)
            - [全通(universe)區塊](#%E5%85%A8%E9%80%9Auniverse%E5%8D%80%E5%A1%8A)
        - [套件層級](#%E5%A5%97%E4%BB%B6%E5%B1%A4%E7%B4%9A)
        - [檔案層級](#%E6%AA%94%E6%A1%88%E5%B1%A4%E7%B4%9A)
    - [轉型](#%E8%BD%89%E5%9E%8B)
        - [Pointer 轉型](#pointer-%E8%BD%89%E5%9E%8B)

```go

func f (){}

var g = "g"

func main (){

    f := "f"
    fmt.Println(f) // 區域變數 遮蔽 套件層級函式
    fmt.Println(g) // 套件層級變數
    fmt.Println(h) // 編譯錯誤， h 未定義
}

```

### 語句(syntactic)區塊

函式內容或迴圈等包圍在括弧內的一連串陳述式，宣告在區塊外不可見

### 詞彙(lexical)區塊

語句區塊中引入程式中其他沒有北括弧直接包圍的宣告群組，可以任意`套疊`，但不建議`套疊`。例：套件、檔案的詞彙區塊

```go

func main (){

   x := "hello"
   for i := 0 ; i < len(x) ; i++ {
       x := x[i]
       if x != '!' {
           x := x + 'A' - 'a'
           fmt.Printf("%c",x) // "HELLO"
       }
   }
}

```

#### 全通(universe)區塊

全程示範圍，`函示`與`常式`所處地

### 套件層級

可被同套件下的任何檔案參考

### 檔案層級

可從同一檔案內參考

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
