# Benchmark 

* 在固定負債下評估程式效能
* 函式名稱以 Benchmark 開頭
* 預設狀態不會執行基準測試

```go

package main

import "testing"

func BenchmarkSum(b *testing.B) {
    // 以 N 次做基準測試
	for i := 0; i < b.N; i++ {
		Sum(1, 2)
	}
}

func Sum(x, y int) int {
	return x + y
}

```

```shell

執行所有基準測試
$ go test -bench='.'

執行函式 Sum 基準測試
$ go test -bench='Sum'

```