# Go 目錄

路徑查看 `go env`

## GOROOT

go 安裝路徑

## GOPATH

* 工作區，可以有多個工作區，但都必須增加至GOPATH中
* 不包含環境變數 GOROOT 的值(Go的安裝目錄路徑)

### 檔案結構

#### src

* 儲存原始程式碼
* go run 或 go install 所執行的路徑

> 副檔名：.go 、 .c 、 .h 、 .s 

##### Go 函式庫原始程式檔案

透過 `go install` 安裝的套件

##### Go 指令原始程式檔案

宣告為 main 的程式套件，可透過 go run 執行。同一套件中的所有原始程式檔案和函數原始程式檔案處於同一個程式套件中，那麼該套件中就無法執行 `go build` 和 `go install`

```go

package main 

```

##### Go 測試原始程式檔案 

透過 `go test` 執行目前套件下的測試檔案

* 檔案須以 `_test.go` 結尾

#### pkg

編譯後產生的`.a`壓縮檔

> 副檔名：.a

#### bin

編譯後產生的可執行檔

## GOBIN

* 編譯後執行檔放置的路徑
* 不允許多個路徑
* 為空時，會將編譯檔案放置在各 GOPATH 的 bin 資料夾中

