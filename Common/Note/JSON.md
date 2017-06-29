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

* struct 欄位相關的編譯其字串資訊
* 字串實字
