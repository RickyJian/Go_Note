# 零值問題

當 value 為零值(ZeroValue)時，並不會帶入 Query 做搜尋條件

```go

db.Where(&Student{Name: "student", Gender: ""}).Find(&Student)
// select * from students WHERE name = "student"

```

## 解決辦法

* 使用指標
* 使用`scanner/valuer`

```go

// Pointer
type Student struct {
    ID        int
    Name      string
    Gender    *string
}

// scanner/valuer
type User struct {
    ID        int
    Name      string
    Gender    sql.NullString
}


```

-----

[零值](/Fundamental/ZeroValue.md)