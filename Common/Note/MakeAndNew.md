# Make 與 New 的比較

資料初始化

## new 

* 建立一T型別的不具名變數，以T零值初始化
* 用來分配到可用的儲存空間，參數為類型，回傳為指標

> new (T)
> func new(Type) *Type

```go


func main (){
    value := new(int) // 回傳指標位址，並零值初始
    fmt.Println(value) // 印出記憶體位址 
    fmt.Println(*value) // 印出值(0)
}

```

## make

用來分配到可用的儲存空間，且只用在 map、slice、channel，參數為類型及長度，回傳為已被初始化類型。

> func make(Type, size IntegerType) Type

```go

func main (){
    slice := make([]int,10,100) //建立一個 int 的 slice，參數二 為長度，參數三 為容量
}

```