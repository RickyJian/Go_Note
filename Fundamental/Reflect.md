# 反射(Reflection)

提供執行其修改變數與檢查其值、呼叫他的方法，與套用他們的內部操作的機制而無須於編譯期知道他們的型別

## reflect.Type

* Type 代表一個 Go 型別
* 判斷型別

```go

// 獲取 interface 的動態型別
t := reflect.TypeOf(2)
fmt.Println(t.String())
fmt.Print("%T",t)

// 獲取 v 型別
t2 := v.Type()

```

## reflect.Value 

保存任何型別的值

```go

// 獲取 interface 的動態值
v := reflect.ValueOf(2)
fmt.Println(v)
fmt.Print("%v",v)

// 

```