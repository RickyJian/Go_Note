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

    log.Fatalf("os.Getwd failed : %s", cwd) // 

    ```

    >  log.Fatalf() ： 會呼叫 os.Exit(1)

### type

可以省略，若省略會運算式自行決定它的型態

### expression

可以省略，若省略初值會是該`type`的零值

> 零值機制 ： 確保變數總是保持正確的值，`Go`沒有未定義的值。

## 指標 (Pointer)

* 變數的位置，指標零值為 `nil`。
* 取用變數的位址或複製一個指標時就會建構新的別名(alias)

```go

x := 1 
p := &x // 型別為 *int 的 p 指向 x ，
fmt.Printf("%d \n",p) // 印出位址
fmt.Printf("%d \n",&p) // 印出p位址
fmt.Printf("%d \n",&x) // 印出x位址
*p = 2 // 等同 x =2 ， 
fmt.Printf("%d \n",*p) // 印出值
p = 2 // 程式會出錯誤，由於值型態不能指派給位址型態


// 傳遞指標參數給函式
func incr(p *int) int {

}

// 傳遞指標參數給函式並回傳函式
func incr(p *int) *int {

}

```

> `*` 轉成`變數` <br>
> `&` 不能放在 `=` 左邊，轉成`位址`，例： var x int 則 `&x` 這個運算式代表x的位址，x的型別會由 int 變成 *int 

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

#### 轉型

* 底層型別相同
* 指向相同底層型別的變數指標
* 有可能改變值的表示

> T(x)：將x值轉換成T型別

## 套件、檔案

* 原始碼能放在一個或多個 `.go`的檔案中
* 放`package`前為**文件註解**，用來說明套件，每個套件只有一個文件註解
* 大量文件放在`doc.go`中


### 匯出(exported)

名稱以大寫字母開頭為

### 匯入(import)

* 路徑，字串
* 匯入宣告將短名稱與匯入的套件綁定已透過檔案來參考它的內容
* 沒有用到匯入會造成編譯錯誤  

> "gopl.io/ch2/tempconv"

## 套件初始化

* 從套件層級變數開始初始化
* 先解決相依性問題，在已宣告順序進行初始化
* `main`最後一個初始化

### init 函式

* 程式啟動後自動以宣告的順序執行

> func init () {}

## 範圍

* 宣告名稱指向該宣告原始碼程式
* 內層宣告會將外層勳宣告`遮蔽`

```go

func f (){}

var g = "g"

funv main (){

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




