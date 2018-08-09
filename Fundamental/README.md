# Go 筆記

## 程式撰寫前

### Catalog

`Go`目錄

* GOROOT
* GOPATH
* 檔案結構
    * src
    * pkg
    * bin
* GOBIN

### Command

常用`Go`指令

* run
* build
* install
* get
* version
* env

## 程式結構

### PackageAndImport

* 套件、檔案
    * package  
    * 匯出(Export)
    * 匯入(Import)

> go 套件搜尋：http://godoc.org

### Identifier

識別符號

* 語言符號
    * 識別符號
        * 限定識別符號

### Keyword

關鍵字

### Operator

運算式

* 算術運算子
* 比較運算子
* 位元運算子
* 地址運算子
* 接收運算子

### Literal 

字面常數

* 基礎資料類型值
* 建置複合資料類型值
* 表示複合資料類型值

### VariableAndDeclaration

宣告

* 區域
* 全域
* 變數
    * 一般
    * 簡短
* 型別
* 指派
    * 指派運算
    * 資料組指派
* 範圍
* 轉型
    * 指標轉型

## 資料型態

### Basic_Datatype

基本資料型別

* 整數
* 浮點數
* 複數
* 布林
* 字串
    * 跳脫序列
* 常數
    * iota 常數產生器

### Pointer

指標

## Aggreate_Datatype

組合型別

### 陣列(Array)

陣列

### slice

切片

* append
* copy

### map

key/value

* delete

### Struct

* 欄位嵌入
* 方法遷入
* 匿名遷入
* 匿名結構
* 字面常數標籤

## Interface

介面

* 介面型別
    * 嵌入
* 空介面
* flag.value
* 介面值
    * 動態型別
    * 動態值
* 型別判別
    * 具體型別
    * 非具體型別
* 型別判斷

## 流程控制

### If

if 判斷式使用方式

### Switch

switch 判斷式使用方式

* 運算式判斷
* 型別判斷

### For

for 迴圈使用方式

#### range

迭代 array , slice , map 

### Defer

延遲

## 方法及函式

### Function

函式

* 不具名函式
* 參數列
    * 參數
    * 型別
* 回傳值
    * 多回傳值
    * 資料類型
* 遞迴
* defer 

###  Method

方法

* 方法宣告
    * 方法值
        * 方法運算式
    * Receive
    * Point Receive

## 錯誤處理

### ERROR 介面

* 建構 error package

### PanicAndRecovery

* panic
* recover

## 平行處理

### goroutine

並行

### channel

goroutine 之間連線

* 建立即關閉
* 通訊
* 管道
* 單向 channel 類型
* 無緩衝 channel
* 有緩衝 channel

### select

### 互斥鎖

* mutex
* rwmutex

### 初始化

once

## 測試

-----
## 其他

### ZeroValue

零值初始

### JSON

JSON 常用方法

### 字串格式化

fmt.printf

### flag

自訂命令列

### Init

套件初始化

### MakeAndNew

資料初始化，了解`new`與`make`的用法

### File I/O

檔案 寫入/寫出
