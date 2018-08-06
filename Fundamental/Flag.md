# flag

* 基本 command line 程式撰寫
* integer、string、boolean 為選用方法
* 回傳型態為指標


[原文網址](https://gobyexample.com/command-line-flags)


```go

func main(){
    // 參數一：command 指令，參數二：預設值，參數三：描述
    demo := flag.String("url", "default", "video url")
    flag.Parse()
    
}

```
