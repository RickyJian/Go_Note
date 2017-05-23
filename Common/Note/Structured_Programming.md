# 程式結構

以 `Package` 開頭，後面接著 `import`

## 宣告

* 以字母底線開頭
* 駝峰字(camel case)
* 原縮寫字維持原大小

> 四種宣告 ： var 、 const 、 type 、 func

### 區域 (local)

* 函式內宣告
* 傾向短名稱

### 全域

* 函式外宣告
* 所屬的套件的全部檔案皆能看見
* 第一個字母決定他的可見度
* 名稱以大寫字母開頭為`匯出(exported)`
* 套件名稱永遠是小寫
* 範圍愈大名稱愈長

## 變數 (Variable) 宣告

值的儲存體，又稱有位置的值

* 一般

    保留給須明確設定參數與型別，或之後被指派值且初值不重要的變數

    > var name type = expression <br>
    > var name type 

    ```go

    var i, j, k int  // 三個變數都是 int 型態

    var b, f, s = true , 2.3 , "four"  // boolean , float , string

    ```

* 簡短宣告

    用於初始化大部分的區域變數，且不能用在全域變數，在使用時至少要宣告一個新變數
    

    >  name := expression

    ```go

    i := 0
    i , j := 0,1

    ```

    > `:=` 宣告 <br>
    > `=`  指派


### type

可以省略，若省略會運算式自行決定它的型態

### expression

可以省略，若省略初值會是該`type`的零值

> 零值機制 ： 確保變數總是保持正確的值，`Go`沒有未定義的值。

## 指標 (Pointer)

變數的位置，指標零值為 `nil`。取

```go

x := 1 
p := &x // 型別為 *int 的 p 指向 x ，
fmt.Printf("%d \n",p) // 印出位址
fmt.Printf("%d \n",&p) // 印出p位址
fmt.Printf("%d \n",&x) // 印出x位址
*p = 2 // 等同 x =2 ， 
fmt.Printf("%d \n",*p) // 印出值
p = 2 // 程式會出錯誤，由於值型態不能指派給位址型態

```

> `*` 不能放在 `=` 右邊，轉成`變數` <br>
> `&` 不能放在 `=` 左邊，轉成`位址`

## new 

建立一T型別的不具名變數，以T零值初始化

> new (T)

## 變數生命週期

* 套件層級：整個執行期間
* 區域變數：動態，當不用就會被回收

### heap

### stack

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

