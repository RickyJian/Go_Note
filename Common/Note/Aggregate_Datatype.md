# 組合型別

`Go`型別：基本型態、**集合(aggregate)型別**、參考型別、介面型別

* 陣列(Array)：靜態
* struct：靜態
* slice：動態
* map：動態

## 陣列(Array)

* 零或多個特定的固定長度序列，若無給值初值為零值初始化
* 大小是型別之一，因此 `[3]int` 與 `[4]int` 不同
* 傳址不傳值，通常會將陣列指標當作是函式參數

基本
```go

var arr [5]int // 宣告三個整數序列
var arr2 [5]int = [5]int {1,2,3,4,5} // 陣列實字初始化
var arr3 [5]int = [5]int {1,2} // 陣列實字初始化
arr4 := [...]int {1,2} // 陣列長度根據初始化樹浪決定
arr[0] // 第一個元素
len(arr) // 陣列長度
arr[len(arr)-1] // 陣列最後元素

// 輸出索引及元素
for i, v := range arr {
    fmt.Printf("%d %d\n",i,v)
}

```

索引與值

```go

type Currnecy int

const (
    USD Currency = iota
    EUR 
    GBP
    RMB
    NTD
)

symbol := [...]string{USD:"$",EUR:"E",GBP:"G",RMB:"R",NTD:"N"}
fmt.Printf(RMB,symbol[RMB]) // 會印出 3 R

r := [...]int{99:-1} // 總共有 100 個元素的陣列，只有最後一個元素為 -1 其餘會被零值初始

```

## slice

* 可變長度序列，其元素型別皆要相同。
* 三元件：長度、指標、容量
* 多個slice 會共用一個底層陣列
* 切割不可超過cap()但可超過len()
* 不能比較，但`nil`可比較


```go

months := []string{1:"Jan.",2:"Feb" ... , 12: "Dec"} // slice 宣告 ，在此索引值0會零值初始 ""
summer := months[6:9]   // len : 3 , cap : 7
summer[:20] // 會導致 panic
summer[:5] // 擴張slice , [June,..., Oct]

a := []int {1,2,3,4,5,6,7,8}

```

### append 函式

slice 加入元素

```go 

runes := []rune

for _, r := range "hello" {
    runes = append (runes,r)
}

```

## map

* 無排序鍵(K)/值(V)群，所有鍵獨一無二
* 鍵值必須可比較
* 若有鍵無值則值為零值初始
* 不可取得位址
    1. map非變數
    2. 會一直成長會導致現有元素重新計算雜湊而放置新位置
* 順序隨機，可用 `sort` 套件排序
* 零值為 `nil`

```go

// 宣告map 

args := make (map[string]int)
args := map[string]int{
    "a":1,
    "b":2,
}
args["a"] = 1

delete(args,"a") // 內鍵刪除函式

// 迭代

for name , age := range args {

}

// ok 型態為 boolean，用來判斷 k 是否存在

if age ,ok := ages["a"]; !ok{/* a is not a key in this map */ }

```

### delete 函式

* 刪除 map 元素
* delete(map,k)

## struct

他將零個或多個任意型別具名值組成一個實體，每個值稱為欄位


```go

type Employee struct {
    ID int
    Name string
    Address string
}

var dilbert Employee

dilbert.Name = "Name"

```