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

## log 基本設定

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