# Panic and Recover

用於報告和執行期間的恐慌

## panic

* 檢查執行期間的錯誤，例：陣列溢位
* 包含 panic 值，panic 發生時每個作用中函式呼叫的堆疊紀錄
* 執行到 panic 會停止程式
* 只接受一個參數，此參數可以是任意型別的值
* 實際參數的類型常常是介面類型 error 的某個實現類型

> 預期中的錯誤常用error，例：I/O

```go

func Reset (x *Buffer){
    if x == nil {
        panic("x is nil")
    }
    x.elements = nil 
}

```


## recover

* 恢復 panic
* 不接受任何參數，但是傳回一個interface{}類型的結果值

```go 

func Parse(input string)(s *Syntax , err error){
    defer func (){
        if p := recover();p != nil{
            err = fmt.Errorf("error")
        }
    }()
}

```