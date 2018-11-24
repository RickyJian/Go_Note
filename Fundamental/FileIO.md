# File I/O

檔案 寫入/寫出

ioutil：[連結](https://golang.org/pkg/io/ioutil/)

os：[連結](https://golang.org/pkg/os/)

bufio：[連結](https://golang.org/pkg/bufio/)

## 開啟

```go

path := "c:\\temp.txt"
f, err := os.Open(path)


```

## 讀取

```go

path := "c:\\temp.txt"
// 方法一
f, err := os.Open(path)
ioutil.ReadAll(f)

// 方法二
content, err := ioutil.ReadFile(path)

```

## 寫出

```go

path := "c:\\temp.txt"
// 方法一
d1 := []byte("hello\ngo\n")
// param1：路徑、param2：內容、param3：檔案權限
err := ioutil.WriteFile(path, d1, 0644)

// 方法二
f, err := os.Create(path)
defer f.Close()
d2 := []byte{115, 111, 109, 101, 10}
n2, err := f.Write(d2)
f.Sync()

// 方法三
f, err := os.Create(path)
defer f.Close()
n3, err := f.WriteString("writes\n")
f.Sync()

// 方法四
f, err := os.Create(path)
defer f.Close()
w := bufio.NewWriter(f)
n4, err := w.WriteString("buffered\n")
w.Flush()

```

## 刪除

刪除檔案或路徑

```go

// 單檔刪除
os.Remove("C:\\tmp\\tmp.mp4")

// 多檔&路徑刪除
os.RemoveAll("C:\\tmp\\complete")

```