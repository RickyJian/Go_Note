# Viper

## 套件下載

```cmd

go get https://github.com/spf13/viper

```

## 設定

### 外部檔案讀取

```go

func init(){
	// config 檔案格式
	viper.SetConfigType("yaml")
	//  config 檔案名稱
	viper.SetConfigName("config")
	// config 放置位置，可新增多個地方
	viper.AddConfigPath("/")
	viper.AddConfigPath("config")
	//	讀取設定檔
	err := viper.ReadInConfig()
	if err != nil{
		fmt.Println(err)
	}
}

```

### io.Reader 讀取

```go

func init() {
    // 內部設定檔值
    configBytes := []byte(`project: viper example`)
    // 讀取設定值
	err := viper.ReadConfig(bytes.NewBuffer(configBytes))
	if err != nil {
		fmt.Println(err)
	}
}

```

## 取值

```go

    // 讀取設定檔值，回傳型態為 interface{}
	viper.Get("project")
	// 讀取設定檔值，回傳型態為  string
	viper.GetString("project")
	// 讀取設定檔值，回傳型態為  map[string][]string
	viper.GetStringMapStringSlice("project")
	// 讀取設定檔值，回傳型態為  map[string]string
	viper.GetStringMapString("project")
	// 讀取設定檔值，回傳型態為  []string
	viper.GetStringSlice("project")
	// 讀取設定檔值，回傳型態為  map[string]interface{}
	viper.GetStringMap("project")
	// 讀取設定檔值，回傳型態為  bool
	viper.GetBool("project")
	// 讀取設定檔值，回傳型態為  int
	viper.GetInt("project")

```