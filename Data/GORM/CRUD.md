# CRUD

## Create

新增，但零值並不會寫入 DB，但零值可以當作 default

```go

account := Account{Name: "test", Age: 18}

db.Create(&account)

```

### NewRecord

判斷 PK 是否為 nil ， PK 為 nil 回傳 true


```go

account := Account{Name: "test", Age: 18}

db.NewRecord(account)

```

### Callback

```go

func (a *Account) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("ID", uuid.New())
    return nil
}

```
