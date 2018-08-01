# Primary Key

主鍵(Primary Key)設定

## 預設 PK

當 struct 裡有 ID 欄位，則會視為預設欄位

```go

type Student struct {
    ID int //  預設 PK
    Name string
}

```

## 自定義 PK


```go

// 非複合 PK
type Student struct {
    Student_ID int `gorm:"primary_key"`
    Name string
}

// 複合 PK
type Student struct {
    Student_ID int `gorm:"primary_key"`
    Name string    `gorm:"primary_key"`
}

```
