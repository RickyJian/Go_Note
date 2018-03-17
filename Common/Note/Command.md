# Go Command

基本 `Go` 指令應用

## go run 

執行 `go` 程式

```go

go run xx.go

```

## go build 

編譯 `.go` 檔案，並在目錄中產生執行檔


```go

go build 

```

## go install

編譯 `.go` 檔案，並在目前工作區產生執行檔

> 當 `GOPATH`包含多個環境目錄，此時指令就會失敗，必須設定環境變數`GOBIN`

## go get

抓取程式碼丟進相對應的目錄下

```go

go get url 

```

## go version

檢查go版本

## go env

查看系統變數