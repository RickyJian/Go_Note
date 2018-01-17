# 組合型別-slice

* 可變長度序列，其元素型別皆要相同。
* 三元件：長度、指標、容量
* 容量是由指標指向的那個元素開始直到最後一個元素
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

## append 函式

* 第一個參數將要被擴充的 slice 值綁定，第二及後續參數作為擴充內容或多個元素值綁定
* 類型必須與第一格參數元素型別相同
* 共用底層陣列
* 建立新的 Slice
* 當大於 Array 容量會建立新 Array 並初始化

slice 加入元素

```go 

// 基本

runes := []rune{}

for _, r := range "hello" {
    runes = append (runes,r)
}

// 進階 限定存取範圍，當存取權限大於容量則會建立新的底層陣列
slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
slice2 := slice[2:4:5]

```

## copy

* 將第二個參數複製到第一個參數
* 類型必須一致
* 會將 slice 長度較長的複製到 slice 長度較短的
* 回傳長度(int)為長度較短的slice長度

```go

slice := []int {0,1,2,3,4,5}
slice2 := []int {9,8,7}
slice3 := copy(slice,slice2)

// result：9 8 7 3 4 5

```