# Cobra

cli 程式用套件

`<  >`：替代成使用者的設定

## 套件下載

```cmd

go get https://github.com/spf13/cobra

```

## 建立專案

他會在 Go 工作目錄下新建新的專案

```cmd

cobra init <projectname or path>

```

project
|--- cmd
|    |___ root.go // 
|--- LICENSE
|___ main.go  // 主要執行檔

## Command Struct

Crobra 核心 Struct，用來設定 Command

### 描述

```go

// 位置：root.go

var rootCmd = &cobra.Command{
	Use:   "<command name>", // Command 名稱
	Short: "<short description>", // 簡短敘述
	Long: `<long description>`  // 詳細敘述
}


```

### Flag 新增

```go

// 位置：root.go

var local string

var global string

func init() {
	// 參數 1：綁定，參數 2：長命令，參數 3：簡短命令，參數 4：預設值，參數 5：描述
	// Global Flag
	rootCmd.PersistentFlags().StringVarP(&global, "global", "g", "", "Global Flag")
	// Local Flag
	rootCmd.Flags().StringVarP(&local, "local", "l", "", "Local Flag")
}

```

### Flag 綁定

```go

// 位置：root.go

var rootCmd = &cobra.Command{
	Args: func(cmd *cobra.Command, args []string) error {
		// 參數處理
  	},
	Run: func(cmd *cobra.Command, args []string) {
		// command 邏輯處理
  	},
}


```

## 執行 Command

```go

// 位置：main.go

import "<project>/cmd"

func main() {
	cmd.Execute()
}

```