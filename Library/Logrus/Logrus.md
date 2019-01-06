# Logrus

* 結構化 log 框架
* log 層級(由低到高)
  * Info：提示訊息
  * Warning：警告訊息，不會造成程式中斷執行
  * Error：錯誤產生，不會造成程式中斷執行
  * Fatal：重大錯誤產生，會造成程式中斷

`<  >`：替代成使用者的設定

## 套件下載

```cmd

go get https://github.com/Sirupsen/logrus

```

## log 設定

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

### 輸出格式設定

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