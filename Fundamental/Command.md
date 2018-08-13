# Go Command

基本 `Go` 指令應用

## go run 

執行 `go` 程式

```shell

$ go run xx.go

```
### 競爭檢查器

```shell

$ go build -race

```

## go build 

* 編譯 `.go` 檔案，並在目錄中產生執行檔
* 建構套件與其相依檔案，然後拋棄所有編譯過的碼，只留下可執行檔案


```shell

$ go build 

```

### 競爭檢查器

```shell

$ go build -race

```

## go install

* 編譯 `.go` 檔案，並在目前工作區產生執行檔
* 建構套件與其相依檔案，會儲存每個套件與命令編譯過的程式碼而不是將它拋棄

> 當 `GOPATH`包含多個環境目錄，此時指令就會失敗，必須設定環境變數`GOBIN`

## go get

抓取程式碼丟進相對應的目錄下

```shell

$ go get url 

```

## go version

檢查go版本

## go env

查看系統變數

## go test

* Go 套件測試工具
* 套件以 _test.go 做結尾，並以 go test 建構