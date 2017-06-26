# Make 與 New 的比較

## new 

用來分配到可用的儲存空間，參數為類型，回傳為指標，會將其零值初始

```go
	
func new(Type) *Type

```

## make

用來分配到可用的儲存空間並初始化，且只用在 map、slice、channel，參數為類型及長度，回傳為類型。

```go

func make(Type, size IntegerType) Type

```