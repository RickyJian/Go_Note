# 介面 (Interface)

`Go`型別：基本型態、集合(aggregate)型別、參考型別、**介面型別**

* 撰寫更具彈性與適應能力的函式
* 函式抽象化，並未定義實作細節
* 零值型別與值均為nil

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

* 介面型別的值
* type 與 value 零值 為 nil
* 可以透過 == 、 != 來比較

> 帶有空指標的介面非空

### 動態(具體)型別 type

編譯期間無法得知動態型別會是甚麼，必須等到介面呼叫時的動態分發

### 動態(型別)值 value

動態分發後獲得的值

## 型別判別

x.(T)：x為介面型別運算式，T為要判別的型別

### T 具體型別

檢查 x 的動態型別是否與 T 相同，檢查成功型別判別結果為 x 的動態值

### T 非具體(介面)型別

檢查 x 的動態型別是否符合 T ，檢查成功不會擷取動態值，介面值為原樣，但介面型別為 T 

## 型別交換

動態辨別這些型別並以不同方式處理，x.(type)

```go

switch x.(type){
    case nil:
    case int:
    case bool:
    case string:
    default:
    
}

```



