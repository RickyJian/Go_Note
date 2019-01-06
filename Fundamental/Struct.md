# Struct

* 他將零個或多個任意型別具名值組成一個實體，每個值稱為欄位
* 欄位：大寫字母可以匯出
* 可以混和匯出與不匯出欄位
* 具名S型別不能宣告型別同為S的欄位，可宣告成 `*S` 指標型別的欄位
* 可做參數傳給函式及回傳，通常較大的struct會將指標傳入函式
* 被呼叫的函式只會接收到參數的拷貝而不是原始資料參數的參考
* 若所有欄位可以比較，則struct可以比較
* 只有嵌入沒有繼承

```go

type Employee struct {
    ID int
    Name string
    Address string
}

type Employee struct {
    ID int
    Name , Address string
}


var dilbert Employee

// 變數

dilbert.Name = "Name"

// 指標

position := &dilbert.Address
*position = "Road:" +  *position

var employeeOfMonth *Employee = &dilbert  // &dilbert 為指標型別所以 Employee 要改為 *Employee 
employeeOfMonth.Address = "Road"
(*employeeOfMonth).Address = "Road"

pp := &Employee{1,"name","road"}  // 建構及初始struct並取得位置
pp := new (Employee)
*pp = Employee{1,"name","road"}  

// 實字宣告方法

position := Employee{1,"name","road"} 

position := Employee{ID : 1,Name : "name", Address : "road"} // 若省略宣告欄位會對此欄位零值初始

// 函式

func Bonus (e *Employee) int{
    /**/
}


```

## 欄位嵌入

```go

type Point struct {
    X , Y int
}

type Circle struct{
    Center Point
    R int 
}


var cir Circle

cir.Center.X = 1

// 等同

cir.X = 1

```

## 方法嵌入

```go

type Point struct {
    X , Y float64
}

type ColoredPoint struct {
    Point 
    Color color.RGBA
}

func (p Point) Distance (q Point) float64{

} 

func (p *Point) ScaleBy(factor float64){

}


func main (){
    var cp ColoredPoint
    cp.X = 1
    red := color.RGBA{255,0,0,255}
    var p = ColoredPoint{Point{1,2},red}
}

```

## 不具名(匿名)嵌入

* 可以有多個具名型別
* 欄位：宣告沒有名字的欄位，該欄位必須是具名型別或不具名型別的指標，但不能用在實字語法，也不可以有兩個同型別的不具名欄位
* 方法：接受器會自動提升

```go

// 欄位

type Point struct {
    X , Y int
}

type Circle struct{
    Point
    R int 
}

type Circle2 struct{
    *Point
    R int 
}

type Wheel struct {
    Circle
    Spokes int
}

var w Wheel

func main (){
    w.X=5 // 與 w.Circle.Point.X 相同

    w = Wheel{8,8,5,20} //編譯錯誤，未知欄位
    w = Wheel{X:8,Y:8,R:9,Spokes:20}//編譯錯誤，未知欄位

    w = Wheel{Circle{Point{8,8},5},20} //編譯成功    
}

```


```go

// 方法

type Point struct {
    X , Y float64
}

type ColoredPoint struct {
    Point 
    Color color.RGBA
}

func (p Point) Distance (q Point) float64{

} 

func (p *Point) ScaleBy(factor float64){

}


func main (){
    var cp ColoredPoint
    cp.X = 1
    red := color.RGBA{255,0,0,255}
    var p = ColoredPoint{Point{1,2},red}

    p.Distance(cp) // 編譯器會將 接受器((p Point)) 提升成 (p ColoredPoint)

}
```

## 匿名結構

* 一次性
* 用於臨時資料存取或資料傳遞

```go

var anymous struct {
    a int 
    b string 
}

anymous2 := struct {
    a int 
    b string
}{0 , "string"}

```

## 字面常數標籤

* 對應欄位的附加屬性
* 標準函數程式套件 `reflect` 中提供的函數檢視到結構類型中的欄位標籤

```go

type Person struct {
    Name string `json:"name"`
    Age uint8 `json:"age"`
}

```

## 參數初始化

若要讓 struct 參數初始化常用 New  或 NewStructName 開頭命名，回傳型態為 struct 指標型態

```go

type Person struct {
    Name string
    Age int
}

func New(name string , age int) *Person {
    return &Pereson{
        Name : name,
        Age : age,
    }
}

func NewPerson(name string , age int) *Person {
    return &Pereson{
        Name : name,
        Age : age,
    }
}


```