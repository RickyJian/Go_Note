# Model Attribute

## default

```go

type Account struct {
    ID      int
    Name    string 
    Gender  string    `gorm:"default:'male'"`
}

```