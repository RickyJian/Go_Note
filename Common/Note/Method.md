# 方法

特殊的function，此function只能在特定型別調用下才能做使用

## 方法宣告

一般function的函式名稱之前加一個額外的參數，此function會依附在該型別上

```go

type Point struct {
    X , Y int
}

func Distance (p , q Point) float64 {
    return math.Hypot(p.X - q.X , p.Y - q.Y)
}

func (p Point) Distance (q Point) float64 {
    return math.Hypot(p.X - q.X , p.Y - q.Y)
}

func main (){
    p := Point{1,2}
    q := Point{1,2}
    p.Distence(q)
}
```

> p 參數稱為方法的`接受器`，這是對呼叫"傳送訊息給物件"的方法的一種物件導向語言的傳統說法 <br>
>  p.Distence 運算式稱為 `selector` <br>

```go
// 同上範例

type Path []Pint 

func (path Path) Distance() float64{
    sum := 0.0
    for i := range path {
        if i > 0 {
            sum += path[i-1].Distance(path[i])
        }
    }
    return sum
}


func main (){
    perim := path {
        {1,2},
        {5,1}
    }
    fmt.Println(perim.Distance())
}


```

> path[i-1]型別為 Point ，因此呼叫 Point.Distance<br>
> perim型別為 Path，呼叫 Path.Distance


### 接受器(Receive)

名稱選擇短一點，且在方法間一致的名稱較好


