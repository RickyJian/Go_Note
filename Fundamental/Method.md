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

type Path []Point 

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

### 方法值

將方法綁定給特定接受器值的函式

```go

type Point struct {
    X , Y float64
}

func (p *Point) Distance(q Point) float64{

}

func (p *Point) ScaleBy(factor float64){

}

func main (){
    p := Point{1,2}
    q := Point{1,2}
    distanceFromP := p.Distance  // 方法值
    fmt.Print(distanceFromP(q))
    scaleP := p.ScaleBy // 方法值
    scaleP(2)
    scaleP(3)
}

```

#### 方法運算式

* 寫做 T.f 或 (*T).f，T 是型別
* 第一個參數是接受器

```go

type Point struct {
	X, Y float64
}

func (p Point) Distence(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) ScaleBy(float float64) {

}

func main() {
	p := Point{1, 8}
	q := Point{1, 2}
	dis := p.Distence
	dis2 := Point.Distence
    sca := (*Point).ScaleBy
	sca2 := (&p).ScaleBy
	fmt.Printf("%T\n", dis) // 印出 func(main.Point) float64
	fmt.Printf("%T\n", dis2) // 印出 func(main.Point, main.Point) float64
	fmt.Printf("%T\n", dis(q)) // 印出 float64
	fmt.Printf("%T\n", dis2(p, q)) // 印出 float64
	fmt.Printf("%T\n", sca) // 印出 func(*main.Point, float64)
	fmt.Printf("%T\n", sca2) // 印出 func(float64)
}


```

### 接受器(Receive)

名稱選擇短一點，且在方法間一致的名稱較好

```go

func (p Point) Distance (q Point) float64 {
    return math.Hypot(p.X - q.X , p.Y - q.Y)
}

func main (){
    p := Point {1,2}
    q := Point {1,2}
    p.Distance(q)

    pptr := &p
    pptr.Distance(q)
    // or
    (*pptr).Distance(q)
}

```

### 指標接受器

```go

func (p *Point) ScaleBy(factor float64){
    p.X *= factory
    p.Y *= factory
}

func main() {

    r := &Point{1,2}
    r.ScaleBy(2)
    fmt.Println(*r) // "{2,4}"

    // or

    p := Point{1,2}
    (&p).ScaleBy(2)
    fmt.Println(p) // "{2,4}"

    p.ScaleBy(2) // p 會自動轉成 &p

    Point{1,2}.ScaleBy(2) // 編譯錯誤 ， 無法取得Point位置

}

```

> 當 Point 任一方法具有指標接受，則 Point 的所有方法應有指標接受器


#### nil 有效的接收器值

容許 nil 指標作為參數
