# File I/O

檔案&目錄 寫入/寫出

ioutil：[連結](https://golang.org/pkg/io/ioutil/)

os：[連結](https://golang.org/pkg/os/)

filepath：[連結](https://golang.org/pkg/path/filepath/)

bufio：[連結](https://golang.org/pkg/bufio/)

## 狀態

### 檢視檔案&目錄屬性

```go

// 若 err != nil 代表檔案或屬性不存在
status, err := os.Stat("C:\\tmp\\")
if err == nil {
    // 判斷是否為目錄
    fmt.Println(status.IsDir())
    // 最後異動時間
    fmt.Println(status.ModTime())
    // 檔案&目錄權限
    fmt.Println(status.Mode())
    // 檔案&目錄名稱
    fmt.Println(status.Name())
    // 檔案&目錄大小
    fmt.Println(status.Size())
}


```

### 檢查檔案&目錄是否存在

```go

path := "C:\\tmp\\"
if _, err := os.Stat(path); os.IsNotExist(err) {
    // 檔案&目錄存在為 false
    fmt.Println(os.IsNotExist(err))
    // 檔案&目錄存在為 true
    fmt.Println(os.IsExist(err))
}

```

## 讀取目錄下所有檔案

### 單層目錄走訪

```go

path := "C:\\"
// 走訪目錄檔案
files, err := ioutil.ReadDir(path)
if err != nil {
    log.Fatal(err)
}

for _, file := range files {
    // 列出檔案或目錄的名稱
    fmt.Println(file.Name())
}

```

### 子目錄走訪

```go

import "path/filepath"

func visit (path string, f os.FileInfo, err error) error {
    // 走訪目錄後的邏輯處理
}

func main (){
    root = "C:\\"
    err := filepath.Walk(root, visit)
}


```

## 目錄新建

```go

path := "C:\\tmp\\"
if _, err := os.Stat(path); os.IsNotExist(err) {
    // 參數一：路徑，參數二：目錄權限
    // 建立一層目錄
    os.Mkdir(path, 777)
    // 建立多層層目錄
    os.MkdirAll(path+"a\\b\\c", 777)
}


```

[檔案&目錄權限解說](http://linux.vbird.org/linux_basic/0210filepermission.php#filepermission_dir)

## 檔案開啟

```go

path := "c:\\temp.txt"
f, err := os.Open(path)


```

## 檔案讀取

```go

path := "c:\\temp.txt"
// 方法一
f, err := os.Open(path)
ioutil.ReadAll(f)

// 方法二
content, err := ioutil.ReadFile(path)

```

## 檔案寫出

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

## 檔案重新命名

```go

// 參數一：舊名稱，參數二：新名稱
err := os.Rename("oldname", "newname")

// 特別用法：檔案移動
path := "c:\\tmp\\newpath\\"
err := os.Rename("oldname", path+"newname")

```

## 檔案複製

```go

orig := "C:\\tmp\\orig.txt"
dest := "C:\\tmp\\dest.txt"
origFile, _ := os.Open(orig)
defer origFile.Close()
destFile, _ := os.Create(dest)
defer destFile.Close()
// 參數一：新檔案，參數二：被複製檔案
bytes, err := io.Copy(destFile, origFile)
if err != nil {
    // error process
}
log.Printf("Copied %d bytes.", bytes)


```

## 刪除

```go

// 單檔刪除
os.Remove("C:\\tmp\\tmp.mp4")

// 多檔&路徑刪除
os.RemoveAll("C:\\tmp\\complete")

```