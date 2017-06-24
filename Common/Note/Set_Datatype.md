# 組合型別

`Go`型別：基本型態、**集合(aggregate)型別**、參考型別、介面型別

* 陣列(Array)：靜態
* struct：靜態
* slice：動態
* map：動態

## 陣列(Array)

* 零或多個特定的固定長度序列，若無給值初值為零值初始化
* 大小是型別之一，因此 `[3]int` 與 `[4]int` 不同


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