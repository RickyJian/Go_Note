# 介面 (Interface)

`Go`型別：基本型態、集合(aggregate)型別、參考型別、**介面型別**

* 撰寫更具彈性與適應能力的函式
* 函式抽象化，並未定義實作細節
* 零值型別與值均為nil
* 實現一個介面就是實作所有介面方法
* 一個資料型別可以實現多個介面

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
    // 檢查介面是否時做成功
    _, ok := interface{}(SortableStrings{}).(sort.Interface)
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

不能嵌入自己本身

```go

type Reader interface {
    Read (p []byte)(n int, err error )
}

type ReadWriter interface {
    Reader
}

```
## 空介面(interface{})

* 可以指派任何型別的值給空介面
* Go語言中的資料型別都是他的實現

```go

var any interface{}

func (args ... interface{})

```

## 介面值

* 介面型別的值
* type 與 value 零值 為 nil
* 可以透過 == 、 != 來比較

> 帶有空指標的介面非空

### 動態(具體)型別 type

編譯期間無法得知動態型別會是甚麼，必須等到介面呼叫時的動態分發

### 動態(型別)值 value

動態分發後獲得的值

## 型別判別與指派

* x.(T)：x為介面型別運算式，T為要判別的型別
* 動態辨別這些型別並以不同方式處理，x.(type)

```go

// 判別
switch x.(type){
    case nil:
    case int:
    case bool:
    case string:
    default:
    
}

// 指派
x.(string)

```

### T 具體型別

檢查 x 的動態型別是否與 T 相同，檢查成功型別判別結果為 x 的動態值

### T 非具體(介面)型別

檢查 x 的動態型別是否符合 T ，檢查成功不會擷取動態值，介面值為原樣，但介面型別為 T 

----- 

[reflect](Reflect.md)