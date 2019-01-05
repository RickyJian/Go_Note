# Model Attribute

## default

```go

type Account struct {
    ID      int
    Name    string 
    Gender  string    `gorm:"default:'male'"`
}

```

### default 為 null

若設定 default 但必須設定為 null 時，可以用 pointer 或 scaner/valuer

```go

// Use pointer value
type User struct {
    gorm.Model
    Name string
    Age  *int   `gorm:"default:18"`
}

// Use scanner/valuer
type User struct {
    gorm.Model
    Name string
    Age  sql.NullInt64  `gorm:"default:18"`
}

```