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


