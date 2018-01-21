# 基本資料型別

`Go`型別：**基本型態**、集合(aggregate)型別、參考型別、介面型別

## 整數

* 正負
* 有號、無號，不能混用要用須轉型

```go

// 有號

int     
int8    
int16   
int32   
int64   
rune // 指示該值為 Unicode 同 int32

 // 無號

uint    
uint8   // 同 byte 
uint16  
uint32  
uint64  
uintptr // 大小未指定，足夠保存指標值的所有位元，用於低階程式設計，例：與 C 函式庫介接 

```

> int、uint ： 根據編譯器決定大小 <br>
> 有號數範圍： -(2^(n-1)) ~ (2^(n-1))-1 <br>
> 以`0`開頭為八進，以`0x`開頭為十六進位數字 

## 浮點數

 float32 ，float64

## 複數

* complex64 ， complex128
* 內建 `real` 、 `imag` 函式

```go

var x complex128 = complex(1,2) // 1+2i
real(x) // 1
imag(x) // 2

```

## 布林

 true ， false 

## 字串

* UTF-8編碼
* 內建 `len`函式
* 值(字串實字)是不可變 (imutable)
* 共用底層記憶體，讓字串複製成本下降
* 一個位元組的序列


```go

s := "hello world"
len(s) // 11 ，字串長度
s[0] // 104 ，讀取字串s的第1個位元組，字串第 i 個位元組並不是第 i 個字元
s[0:5] // "hello" ，子字串
s[0] = 'L' // 編譯錯誤，不可改變字串值 

```

> s[i:j] ： 複製字串從 i 到 j 但不函 j <br>
> 短路：`&&`、`||` 左運算元決定後，則右運算元不會決定

### 字串實字

包在 `""` 中一系列位元組

#### 跳脫序列


| 跳脫序列 | 用法 |
| ---------|---------- |
| \a  | 警告或響鈴 |
| \b  | 後退 |
| \f  | 跳頁 |
| \n  | 換行 | 
| \r  | 回車 |
| \t  | tab |
| \v  | 垂直 tab |
| \'  | 單引號(只用於`rune`實字"\") |
| \"  | 雙引號(僅限"..."實字) |
| \\  | 反斜線 |
| \xh | 十六進位跳脫 |
| \o  | 八進位跳脫 |

> 原始字串實字：反引號(``)取代雙引號，跳脫序列不會被處理，可以在程式碼分成很多行，唯一會處理的是刪除回車。常用於正規表示式。

## 常數

值與計算在編譯期完成，可以是任何基本型別

```go

const c int = 1000

```

### iota 常數產生器

相關值無須明確指定，`iota`值從零開始遞增

```go

type Weekday int

const{
    Sunday Weekday = iota
    Monday 
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
}

```

> Sunday 為 0 ，Monday 為 1 ，以此類推

### 無型別常數

字面常數表示的常數，例：true、false、"A"、iota`

```go

const c = 1.0

```