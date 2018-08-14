# Test

套件以 _test.go 做結尾

## 單元測試

函式名稱以 Test 開頭，對程式邏輯做檢查

```go
package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 6)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
func TestMultiple(t *testing.T) {
	product := Sum(5, 7)
	if product != 30 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", product, 30)
	}
}

func Sum(x, y int) int {
	return x + y
}

func Multiple(x, y int) int {
	return x * y
}


```

```shell

$ go test

// 輸出套件中每個測試的名稱與執行時間
$ go test -v

// 指定執行測試
$ go test -run="Mutilple"

```

## 基準測試

函式名稱以 Benchmark 開頭，評估函式效能

## 範例

函式名稱以 Example 開頭，提供機器檢查文件

-----

[Command](Command.md)