# switch

## 運算式判斷

```go

switch content {
    default :
        // 實作
    case "Python" :
        // 實作
    case "Go" :
        // 實作
}

switch content := "Go" ; content {
    default :
        // 實作
    case "Python" :
        // 實作
    case "Go" :
        // 實作
}

switch content := "Go" ; content {
    default :
        // 實作
    case "Python" , "Ruby" :
        // 實作
    case "Go" , "JAVA" :
        // 實作
}

switch content := "Go" ; content {
    default :
        // 實作
    case "Python" :
        // 實作
        fallthrough
    case  "Ruby"  :
        // 實作
    case "Go" , "JAVA" :
        // 實作
}

switch content {
    default :
        // 實作
    case "Python" :
        // 實作
        break
    case "Go" :
        // 實作
}

```

> default 只能有一個 <br>
> fallthrough 會將流程控制權交給下一個 case 處理 <br>
> break 直接結束 switch 

## 型態判斷

用於介面型態

```go

switch x.(type){
    case nil:
    case int:
    case bool:
    case string:
    default:
    
}

switch i := x.(type){
    case nil:
    case int:
    case bool:
    case string:
    default:
    
}

```

> 若 x 為 nil 則結果也會為 nil <br>
> fallthrough 不允許出現在此