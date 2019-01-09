# CLI

## 套件下載

```cmd

// 穩定版
go get gopkg.in/urfave/cli.v1

// 測試版
go get gopkg.in/urfave/cli.v2

```

## 設定

```go

package main

import (
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	// CLI 名稱
	app.Name = "cliName"
	// CLI 敘述
	app.Usage = "Learn urfave/cli"

}

```

## Flag 新增 & 綁定

```go

var name string

// 新增 & 綁定 Flag
app.Flags = []cli.Flag{
	cli.StringFlag{
		// Flag 名稱，長名稱 name，短名稱 n
		Name: "name,n",
		// 預設值
		Value: "John",
		// Flag 描述
		Usage: "Name description",
		// 綁定 Flag
		Destination: &name,
	},
}

```

## 參數及邏輯處理

```go

	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello", name)
		return nil
	}

```

## 執行

```go

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

```