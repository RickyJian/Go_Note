# Model

與 database table 對應的 struct

## gorm.Model

預設 struct

```go

type Model struct {
    ID        uint `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time
}


```

### 嵌入預設 Model

```go

type Account struct {
    gorm.Model
    Password stringsss
}


```

----

[struct](/Common/Note/Struct.md)