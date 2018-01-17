# 組合型別-map

* 無排序鍵(K)/值(V)群，所有鍵獨一無二
* 鍵值必須可比較
* 若有鍵無值則值為零值初始
* 不可取得位址
    1. map非變數
    2. 會一直成長會導致現有元素重新計算雜湊而放置新位置
* 順序隨機，可用 `sort` 套件排序
* 零值為 `nil`
* 並非平行處理安全

```go

// 宣告map 

args := make (map[string]int)
args := map[string]int{
    "a":1,
    "b":2,
}
args["a"] = 1

delete(args,"a") // 內鍵刪除函式

// 迭代

for name , age := range args {

}

// ok 型態為 boolean，用來判斷 k 是否存在

if age ,ok := ages["a"]; !ok{/* a is not a key in this map */ }

```

## delete 函式

* 刪除 map 元素
* delete(map,k)
