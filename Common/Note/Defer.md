# Defer 

* 延遲函數(程式讀取：由上而下再由下而上)
* 當函式本體正常執行完後，則會執行 defer 敘述，此函示才真正執行完
* 當函式本體執行 return 時，只有在 defer 敘述全部執行完畢後才真正回傳
* 當函式本體執行時發生恐慌時，只有在 defer 敘述全部執行完畢後才真正擴散至涵式呼叫方

```go

defer fmt.Println("defer")

```

## 延遲函式

```go

func Begin (funcName string) string{
    fmt.Printf("Enter Begin Func %s. \n ",funcName)
    return funcName
}

func End (funcName string) string{
    fmt.Printf("Exit End Func %s. \n ",funcName)
    return funcName
}

func main(){
    defer End(Begin("Record"))
    fmt.Println("In Main Func.")
}

// result
// Enter Begin Func Record.
// In Main Func
// Exit End Func Record.

```

延遲傳值

```go

func main() {
	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}


```

### 匿名函式

```go

defer func () {
    fmt.Println("defer")
}()

```