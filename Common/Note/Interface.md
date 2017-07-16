# 介面 (Interface)

* 撰寫更具彈性與適應能力的函式
* 函式抽象化，並未定義實作細節

> 具體型別；知道它`是`什麼，並且知道可以對它`做`什麼 <br>
> 介面型別：不知道它`是`什麼，只知道可以對它`做`什麼 

介面
```go

type Writer interface {
    Write(p []byte)(n int ,err error)
}

```

介面實作
```go
type ByteCounter int

func Write(p []byte)(n int ,err error){
     //實作過程 
}

func main (){
    c.Write([]byte("hello") )
}

```

## 介面型別

指定具體型別

```go

type Reader interface {
    Read (p []byte)(n int, err error )
}

```

### 介面嵌入

```go

type Reader interface {
    Read (p []byte)(n int, err error )
}

type ReadWriter interface {
    Reader
}

```
## 空介面(interface{})

可以指派任何型別的值給空介面

```go

var any interface{}

func (args ... interface{})

```

## flag.Value

定義新的命令列旗標



```go
// period 新的命令
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Seeping for %v... ", *period)
	time.Sleep(*period)
	fmt.Println()
}

```

> . / sleep -period 50ms

## 介面值

介面型別的值

> 型別描述器

### 動態型別

### 動態值