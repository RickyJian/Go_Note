# 正規表示式(Regex)

```go

// 匯入套件

import "regexp"

```

## 函式&方法

### Compile

編譯正規表示式(regular expression)，若編譯正確則之後程式可以繼續使用

```go

r , err := regexp.Compile("regexp")  // regexp 放置 表示式符號

```

### MustCompile

與 Compile 函式類似，但只會回傳單一值，若失敗則直接 panic

```go

r = regexp.MustCompile("regexp")   // regexp 放置 表示式符號

```

### Match

判斷字串是否有符合表示式，以[]byte型態為參數，回傳型態為布林(bool)

```go

r := regexp.MustCompile("hap*y")
fmt.Println(r.Match([]byte("happy")))   // true

```

### MatchString

與 match 使用方式相仿，以字串型態為參數

```go

r := regexp.MustCompile("hap*y")
fmt.Println(r.MatchString("happy"))   // true

```

### FindString

找出第一個符合表示式的字串，回傳型態為字串

```go

r := regexp.MustCompile("o{2}")
fmt.Println(r.FindString("google doodle noodle bool")) // 回傳 oo  

```

### FindStringIndex

找出第一個符合表示式的字串，回傳型態為整數 slice 型態

```go

r := regexp.MustCompile("o{2}")
fmt.Println(r.FindStringIndex("google doodle noodle bool")) // 回傳 [1 3]

```

### FindAllString

找出第一個符合表示式的字串，回傳型態為字串 slice 型態，參數 1 為要找出的字串，參數 2 回傳 slice 數量，-1 為全部回傳

```go

r := regexp.MustCompile("o{2}")
fmt.Println(r.FindAllString("google doodle noodle bool",-1)) // 回傳 [oo oo oo oo]

```

### FindAllStringIndex

找出第一個符合表示式的字串，回傳型態為整數二維 slice 型態，參數 1 為要找出的字串，參數 2 回傳 slice 數量，-1 為全部回傳

```go

r := regexp.MustCompile("o{2}")
fmt.Println(r.FindAllString("google doodle noodle bool",-1)) // 回傳 [[1 3] [8 10] [15 17] [22 24]]

```

### ReplaceAllLiteralString

當表示式符合時，以字串實字方式替換原字串中的字，參數 1 被替換的字串，參數 2 替換用的字串

```go

r := regexp.MustCompile("o{2}")
fmt.Println(r.ReplaceAllLiteralString("google doodle noodle bool","00"))  // 回傳 g00gle d00dle n00dle b00l

```

## 表示式符號

| 語法 | 說明 |
| ----- | ----- |
| ^ | 匹配輸入字串的開始位置 |
| $ | 匹配輸入字串的結束位置 |
| . | 匹配除換行(\n)外的任一字元 |
| * | 匹配前一字元零次或多次等於{0,} |
| + | 匹配前一字元一次或多次等於{1,} |
| ? | 匹配前一字元零次或一次等於{0,1} |
| \s | 匹配任何空白字元 |
| \S | 匹配任何**非**空白字元 |
| \d | 匹配任何數字 |
| \D | 匹配任何**非**數字 |
| \w | 匹配任何字，等於 [a-zA-Z0-9_] |
| \W | 匹配任何**非**字，等於 [^a-zA-Z0-9_] |
| \B | 匹配字首與字尾間的字元 |
| (a\|b) | 匹配`(a\|b)`中的 a 或 b 字元 |
| [abc] | 匹配`[]`所包含的任一字元 |
| [^abc] | 匹配`[]`**未包含**的任一字元 |
| [a-z] <br> [a-zA-Z] | 匹配`[]`範圍中的任一字元，大小寫有區別 |
| {n} | n為一個非負整數，匹配n次單一字元 |
| {n,} | n為一個非負整數，最少匹配n次單一字元 |
| {n,m} | n、m為一個非負整數，n <= m，最少匹配n次單一字元，最多匹配m次 |

```go

// 範例 ^
r := regexp.MustCompile("^Apple")
fmt.Println(r.MatchString("Apple"))   // true
fmt.Println(r.MatchString("apple"))   // false

// 範例 $
r := regexp.MustCompile("Apple$")
fmt.Println(r.MatchString("Apple"))   // true
fmt.Println(r.MatchString("ApplE"))   // false

// 範例 .
r := regexp.MustCompile("App.e")
fmt.Println(r.MatchString("Apple"))   // true

// 範例 *
r := regexp.MustCompile("hap*y")
fmt.Println(r.MatchString("happy"))   // true
fmt.Println(r.MatchString("hapy"))    // true
fmt.Println(r.MatchString("hapsy"))   // false

// 範例 +
r := regexp.MustCompile("hap+y")
fmt.Println(r.MatchString("happy"))   // true
fmt.Println(r.MatchString("happppy")) // true
fmt.Println(r.MatchString("hay"))     // false

// 範例 ?
r := regexp.MustCompile("hap?y")
fmt.Println(r.MatchString("happy"))   // false
fmt.Println(r.MatchString("hapy"))    // true
fmt.Println(r.MatchString("hay"))     // true

// 範例 \s
r := regexp.MustCompile("hap\\sy")
fmt.Println(r.MatchString("hap y"))   // true
fmt.Println(r.MatchString("happy"))   // false

// 範例 \S
r := regexp.MustCompile("hap\\Sy")
fmt.Println(r.MatchString("happy"))   // true
fmt.Println(r.MatchString("hap y"))   // false

// 範例 \d
r := regexp.MustCompile("hap\\dy")
fmt.Println(r.MatchString("hap9y"))   // true
fmt.Println(r.MatchString("happy"))   // false

// 範例 \D
r := regexp.MustCompile("hap\\Dy")
fmt.Println(r.MatchString("happy"))   // true
fmt.Println(r.MatchString("hap9y"))   // false

// 範例 \w
r := regexp.MustCompile("hap\\wy")
fmt.Println(r.MatchString("happy"))   // true
fmt.Println(r.MatchString("hap y"))   // false

// 範例 \W
r := regexp.MustCompile("hap\\wy")
fmt.Println(r.MatchString("hap y"))   // true
fmt.Println(r.MatchString("happy"))   // false

// 範例 \B
r := regexp.MustCompile("app\\B")
fmt.Println(r.MatchString("happy"))   // true
fmt.Println(r.MatchString("apple"))   // true
fmt.Println(r.MatchString("webapp"))  // false

// 範例 (a|b)
r := regexp.MustCompile("(bbc|cnn)news")
fmt.Println(r.MatchString("bbcnews"))   // true
fmt.Println(r.MatchString("cnnnews"))   // true
fmt.Println(r.MatchString("tvbsnews"))  // false

// 範例 []
r := regexp.MustCompile("[beilmow]app")
fmt.Println(r.MatchString("webapp"))    // true
fmt.Println(r.MatchString("mobileapp")) // true
fmt.Println(r.MatchString("newsapp"))   // false

r := regexp.MustCompile("[^beilmow]app")
fmt.Println(r.MatchString("webapp"))    // false
fmt.Println(r.MatchString("mobileapp")) // false
fmt.Println(r.MatchString("newsapp"))   // true

// 範例 {}
r := regexp.MustCompile("go{2}gle")
fmt.Println(r.MatchString("google"))    // true
fmt.Println(r.MatchString("goooogle"))  // false
fmt.Println(r.MatchString("ggle"))      // false

r := regexp.MustCompile("go{2,}gle")
fmt.Println(r.MatchString("google"))    // true
fmt.Println(r.MatchString("goooogle"))  // true
fmt.Println(r.MatchString("ggle"))      // false

r := regexp.MustCompile("go{1,4}gle")
fmt.Println(r.MatchString("google"))    // true
fmt.Println(r.MatchString("goooogle"))  // true
fmt.Println(r.MatchString("gooooogle")) // false
fmt.Println(r.MatchString("ggle"))      // false

```
