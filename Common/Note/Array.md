# 組合型別-陣列(Array)

* 零或多個特定的固定長度序列，若無給值初值為零值初始化
* 大小是型別之一，因此 `[3]int` 與 `[4]int` 不同
* 傳址不傳值，通常會將陣列指標當作是函式參數

基本
```go

var arr [5]int // 宣告五個整數序列
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

symbol2 := [6]string{2:"Go",1:"Python",5:"JAVA",3:"C",4:"PHP",0:"C++"}

symbol3 := [6]string{2:"Go",1:"Python","JAVA","C",4:"PHP","C++"}

r := [...]int{99:-1} // 總共有 100 個元素的陣列，只有最後一個元素為 -1 其餘會被零值初始

```
