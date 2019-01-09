# Logrus

* 結構化 log 框架
* log 層級(由低到高)
  * Debug：除錯模式
  * Info：提示訊息
  * Warning：警告訊息，不會造成程式中斷執行
  * Error：錯誤產生，不會造成程式中斷執行
  * Fatal：重大錯誤產生，會造成程式中斷，會調用 os.Exit(1)
  * Panic： 程式恐慌 panic 產生

`<  >`：替代成使用者的設定

## 套件下載

```cmd

go get https://github.com/Sirupsen/logrus

```

## 設定

### 基本設定

```go

func init() {
	// 輸出格式，預設為 logrus.TextFormatter
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 輸出位置，參數必須實作 io.Writer，預設為 os.Stdout(輸出到終端機)
	logrus.SetOutput(os.Stdout)

	// 輸出層級，會輸出指定層級以上的訊息
	logrus.SetLevel(logrus.WarnLevel)
}

```

#### 日期時間格式

```go

func init(){
    // 輸出格式，預設為 logrus.TextFormatter
	formatter := new(logrus.TextFormatter)
	// 時間格式輸出
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	logrus.SetFormatter(formatter)
}

```

#### 檔案輸出

```go

	file, err := os.OpenFile("C:\\tmp\\log.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		fmt.Println(err)
	}
    
```

## 輸出

### 基本

輸出內容只有時間、錯誤層級、錯誤訊息

```go

func main (){
    logrus.Info("log info.")
    logrus.Warning("log warning.")
    logrus.Error("log error.")
    logrus.Fatal("log fatal.")
}


```

### 自訂

將自訂 struct 輸出

```go

func main (){
    logStruct := logrus.WithFields(
	logrus.Fields{
		"count": 1,
		"place": "School",
	})
	logStruct.Error("Warning")
}

```

## 進階使用

在此必須使用 hook 這個 interface 並實作 level 及 Fire 方法

```go

type Hook interface {
    Levels() []Level
    Fire(*Entry) error
}

```

Level：指定處理層級

Fire：邏輯處理地方

[Hook 文件](https://godoc.org/github.com/sirupsen/logrus#Hook)

### 指定層級處理

```go

// 定義一個 struct 
type ErrorLog struct {
	UserName string
	Gender   string
}

var users = []ErrorLog{
	{"John", "Male"},
	{"Sandy", "Female"},
}

// Fire：邏輯處理地方
func (el *ErrorLog) Fire(entry *logrus.Entry) error {
	// 將 Key users 加入在輸出中
	entry.Data["users"] = users
	return nil
}

// 指定輸出層級
func (el *ErrorLog) Levels() []logrus.Level {
	// 指定輸出在 Error 層級，若要輸出在所有層級則用 logrus.AllLevels 
	return []logrus.Level{logrus.ErrorLevel}
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.AddHook(&ErrorLog{})
	logrus.Info("log info.")
	logrus.Warning("log warning.")
	logrus.Error("log error.")
	logrus.Fatal("log fatal.")
	// 結果顯示
	// {"level":"info","msg":"log info.","time":"2019-01-09T16:20:50+08:00"}
	// {"level":"warning","msg":"log warning.","time":"2019-01-09T16:20:50+08:00"}
	// {"level":"error","msg":"log error.","time":"2019-01-09T16:20:50+08:00","users":[{"UserName":"John","Gender":"Male"},{"UserName":"Sandy","Gender":"Female"}]}
	// {"level":"fatal","msg":"log fatal.","time":"2019-01-09T16:20:50+08:00"}
}

```

### 多位置輸出

```go

type DefaultLog struct {
	InfoPath  string
	WarnPath  string
	ErrorPath string
}

var paths = DefaultLog{"C:\\tmp\\infolog.txt", "C:\\tmp\\warnlog.txt", "C:\\tmp\\errorlog.txt"}

func (dl *DefaultLog) Fire(entry *logrus.Entry) error {
	// entry.String() 獲得 log 資訊
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}
	// entry.Level 獲取錯誤層級
	switch entry.Level {
	case logrus.ErrorLevel:
		file, _ := os.OpenFile(paths.ErrorPath, os.O_WRONLY|os.O_CREATE, 0666)
		defer file.Close()
		file.WriteString(line)
	case logrus.WarnLevel:
		file, _ := os.OpenFile(paths.WarnPath, os.O_WRONLY|os.O_CREATE, 0666)
		defer file.Close()
		file.WriteString(line)
	case logrus.InfoLevel:
		file, _ := os.OpenFile(paths.InfoPath, os.O_WRONLY|os.O_CREATE, 0666)
		defer file.Close()
		file.WriteString(line)
	}
	return nil
}

func (dl *DefaultLog) Levels() []logrus.Level {
	return logrus.AllLevels
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.AddHook(&DefaultLog{})
	logrus.Info("log info.")
	logrus.Debug("log debug.")
	logrus.Warning("log warning.")
	logrus.Error("log error.")
}

```