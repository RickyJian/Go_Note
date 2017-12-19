# 指標 (Pointer)

* 變數的位置，指標零值為 `nil`。
* 取用變數的位址或複製一個指標時就會建構新的別名(alias)

```go

x := 1 
p := &x // 型別為 *int 的 p 指向 x ，
fmt.Printf("%d \n",p) // 印出位址
fmt.Printf("%d \n",&p) // 印出p位址
fmt.Printf("%d \n",&x) // 印出x位址
*p = 2 // 等同 x =2 ， 
fmt.Printf("%d \n",*p) // 印出值
p = 2 // 程式會出錯誤，由於值型態不能指派給位址型態


// 傳遞指標參數給函式
func incr(p *int) int {

}

// 傳遞指標參數給函式並回傳函式
func incr(p *int) *int {

}

```

> `*` 放在`=`左邊用來設定指標`變數` <br>
> `&` 不能放在 `=` 左邊，轉成`位址`，例： var x int 則 `&x` 這個運算式代表x的位址，x的型別會由 int 變成 *int 
