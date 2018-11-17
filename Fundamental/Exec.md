# 外部指令執行

```go

匯入套件

import "os/exec"

```

## 函式&方法

### Command

設置 Command，回傳指標型別

```go

// 參數一：外部 Command 名稱，參數二：外部 Command 參數的參數

cmd := exec.Command("ls","-li")

```

### Run

執行 Command，並等待他完成

```go

cmd := exec.Command("ls","-li")

// 執行
err := cmd.Run()
if err != nil {
    log.Fatal(err)
}

// 外部 Command 錯誤處理
var stdout, stderr bytes.Buffer
cmd.Stdout = &stdout
cmd.Stderr = &stderr
fmt.Println(string(stdout.Bytes()))
fmt.Println(string(stderr.Bytes()))

```

### Start

執行 Command，並**不**等待他完成

```go

cmd := exec.Command("ls","-li")

// 執行

err := cmd.Start()

```

### Output

執行並回傳外部 Command 執行結果

```go

// 參數一：外部 Command 回傳結果，型態為 []byte，參數二：錯誤回傳
out , err := exec.Command("ls","-li").Output()

if err != nil {
    log.Println(err)
}
fmt.Println(string(out))

```

### CombinedOutput()

執行並回傳外部 Command 執行結果(合併stdout&stderr)

```go

out, err := exec.Command("ls","-li").CombinedOutput()
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(out))

```