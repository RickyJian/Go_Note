# Make 與 New 的比較

資料初始化

## new 

* 建立一T型別的不具名變數，以T零值初始化
* 用來分配到可用的儲存空間，參數為類型，回傳為指標

> new (T)

```go
	
func new(Type) *Type

```

## make

用來分配到可用的儲存空間並初始化，且只用在 map、slice、channel，參數為類型及長度，回傳為類型。

```go

func make(Type, size IntegerType) Type

```