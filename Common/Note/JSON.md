# JSON

JavaScript Object Notation，用來傳送與接收結構化資料的標準記號法，在`Go`裡由 `encoding/json` 提供
<br>
Unicode 編碼
<br>

```javascript

{"name":value,"arr":[1,2,3],"obj":{"name":value,"name":value}}

```

## Marshal 函式

將`Go`的資料結構轉成json<br>
將struct欄位(匯出)名稱傳成json name

## Unmarshal 函式

將json的資料結構轉成`Go`

## MarshalIndent 函式

將`Go`的資料結構轉成json，並整齊縮排

## 欄位標籤

* struct 轉成 json 所替代的名稱
* 字串實字

```go

// Year 會轉成 released
// Color 會轉成 color

type Movie struct {
    Title string 
    Year int `json:released`
    Color bool `json:color,omitempty`
    Actors []string
}

```

> omitempty：若蘭為為其型別的零值或空則不應輸出
