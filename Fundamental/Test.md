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
func TestProduct(t *testing.T) {
	product := Sum(5, 7)
	if product != 30 {
		t.Errorf("Product was incorrect, got: %d, want: %d.", product, 30)
	}
}

func Sum(x, y int) int {
	return x + y
}

func Product(x, y int) int {
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

### 常用測試方法

* 表格方式測試
* t.Errof 並不會中斷測試，測試完後會拋出完整堆疊紀錄
* t.Fatalf 會停止測試

```go

package main

import "testing"

var testsSum = []struct {
	x    int
	y    int
	want int
}{
	{1, 1, 2},
	{1, 2, 2},
	{1, 3, 2},
	{1, 4, 5},
	{1, 5, 6},
	{1, 6, 10},
}

var testsProduct = []struct {
	x    int
	y    int
	want int
}{
	{1, 1, 2},
	{1, 2, 2},
	{1, 3, 2},
	{1, 4, 4},
	{1, 5, 5},
	{1, 6, 6},
}

func TestSum(t *testing.T) {
	for _, test := range testsSum {
		if got := Sum(test.x, test.y); got != test.want {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", test, test.want)
		}
	}
}
func TestProduct(t *testing.T) {
	for _, test := range testsProduct {
		if got := Product(test.x, test.y); got != test.want {
			t.Errorf("Product was incorrect, got: %d, want: %d.", test, test.want)
		}
	}
}

func Sum(x, y int) int {
	return x + y
}

func Product(x, y int) int {
	return x * y
}


```

### 隨機化測試

* 建構隨機輸入進行廣泛的測試
* 會記錄測試失敗資料

```go

package main

import (
	"math/rand"
	"testing"
	"time"
)

func randSum(rng *rand.Rand) (x, y, want int) {
	x = rng.Intn(10)
	y = rng.Intn(10)
	want = rng.Intn(20)
	return
}
func TestRandSum(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random Seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 100; i++ {
		x, y, want := randSum(rng)
		if got := Sum(x, y); got != want {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
		}
	}
}

func Sum(x, y int) int {
	return x + y
}


```

### 覆蓋率

套件受測試的程度稱為測試的覆蓋率

```shell

$ go test -run=Sum -coverprofile='c.out' folder

$ go tool cover -html=c.out

```

`-coverprofile`：指定輸出文件
`-html`：指定輸出格式

## 基準測試

函式名稱以 Benchmark 開頭，評估函式效能

## 範例

函式名稱以 Example 開頭，提供機器檢查文件

-----

[Command](Command.md)