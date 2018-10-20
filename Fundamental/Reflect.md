# 反射(Reflection)

提供執行其修改變數與檢查其值、呼叫他的方法，與套用他們的內部操作的機制而無須於編譯期知道他們的型別

## Type

* Type 代表一個 Go 型別
* 判斷型別

```go

v := 3

// 獲取 v 的型別
t := v.Type()

```

### TypeOf()

* 接受任意 interface{} 型別
* 獲取動態型別

```go

// 獲取 interface 的動態型別
t := reflect.TypeOf(2)
fmt.Println(t.String())
fmt.Print("%T",t)

```

## Value

保存任何型別的值

### ValueOf()

* 接受任意 interface{} 型別
* 獲取動態值

```go

// 獲取 interface 的動態值
v := reflect.ValueOf(2)
fmt.Println(v)
fmt.Print("%v",v)

```

### Kind()


#### NumField()

#### Filed()