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

[連結](Catalog.md)

### Command

常用`Go`指令

* run
* build
* install
* get
* version
* env

[連結](Command.md)

## Doc

Go 文件註解

[連結](Doc.md)

## 程式結構

### PackageAndImport

* 套件、檔案
  * package  
  * 匯出(Export)
  * 匯入(Import)

> go 套件搜尋：http://godoc.org

[連結](PackageAndImport.md)

### Identifier

識別符號

* 語言符號
  * 識別符號
    * 限定識別符號


[連結](Identifier.md)

### Keyword

關鍵字

[連結](Keyword.md)

### Operator

運算式

* 算術運算子
* 比較運算子
* 位元運算子
* 地址運算子
* 接收運算子

[連結](Operator.md)

### Literal 

字面常數

* 基礎資料類型值
* 建置複合資料類型值
* 表示複合資料類型值

[連結](Literal.md)

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

[連結](VariableAndDeclaration.md)

## 資料型態

### BasicDatatype

基本資料型別

* 整數
* 浮點數
* 複數
* 布林
* 字串
  * 跳脫序列
* 常數
  * iota 常數產生器

[連結](BasicDatatype.md)

### Pointer

指標

[連結](Pointer.md)

## Aggreate_Datatype

組合型別

### 陣列(Array)

陣列

[連結](Array.md)

### Slice

切片

* append
* copy

[連結](Slice.md)

### Map

key/value

* delete

[連結](Map.md)

### Struct

* 欄位嵌入
* 方法遷入
* 匿名遷入
* 匿名結構
* 字面常數標籤
* 參數初始化

[連結](Struct.md)

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

[連結](Interface.md)

## 流程控制

### If

if 判斷式使用方式

[連結](If.md)

### Switch

switch 判斷式使用方式

* 運算式判斷
* 型別判斷

[連結](Switch.md)

### For

for 迴圈使用方式

[連結](For.md)

### Defer

延遲

[連結](Defer.md)

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

[連結](Function.md)

### Range

迭代 array , slice , map

[連結](Range.md)

### Init

套件初始化

[連結](Init.md)

### Method

方法

* 方法宣告
  * 方法值
    * 方法運算式
  * Receive
  * Point Receive

[連結](Method.md)

## 錯誤處理

### ERROR 介面

* 建構 error package

[連結](ERROR.md)

### PanicAndRecovery

* panic
* recover

[連結](PanicAndRecovery.md)

## 平行處理

### goroutine

並行

[連結](Goroutine.md)

### channel

goroutine 之間連線

* 建立即關閉
* 通訊
* 管道
* 單向 channel 類型
* 無緩衝 channel
* 有緩衝 channel

[連結](channel.md)

### pipe

### select

判定 channel 狀態並選擇該執行的邏輯

[連結](Select.md)

### 互斥鎖

* mutex
* rwmutex

[連結](Mutex.md)

### 初始化

[連結](Once.md)

## 測試

Go 測試程式撰寫

### UniTest

單元測試撰寫

[連結](UniTest.md)

### Benchmark

基準測試

[連結](Benchmark.md)

## Log

[連結](Log.md)

-----

## 其他

### ZeroValue

零值初始

[連結](ZeroValue.md)

### JSON

JSON 常用方法

[連結](JSON.md)

### 字串格式化

fmt.printf

[連結](Fmt.md)

### CLI

自訂命令列

[連結](CLI.md)

### MakeAndNew

資料初始化，了解`new`與`make`的用法

[連結](MakeAndNew.md)

### File I/O

檔案 寫入/寫出

[連結](FileIO.md)

### 正規表示式(Regex)

* 函式&方法
  * Compile
  * MustCompile
  * Match
  * MatchString
  * FindString
  * FindStringIndex
  * FindAllString
  * FindAllStringIndex
  * ReplaceAllLiteralString
* 表示式符號

[原文網址](https://golang.org/pkg/regexp/#example_Regexp_MatchString)

[連結](Regex.md)

### Strings

常用字串工具

* Contains
* Split
* Join
* Trim
* TrimSpace
* TrimLeft
* TrimPrefix
* TrimRight
* TrimSuffix
* Title
* ToLower
* ToUpper

[原文網址](https://golang.org/pkg/strings/)

[連結](Strings.md)

### Exec

執行外部 Command

[連結](Exec.md)

### Datetime

日期時間處理

* 函式&方法
  * Now
  * Format
  * 時區設定
  * 抓取日期&時間
  * 日期運算
  * 日期比較

[連結](Datetime.md)