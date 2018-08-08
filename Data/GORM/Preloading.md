# Preloading(Eager Loading)

```go

db.Preload("order").find(&user)
// select * from user
// select * from order where user_id IN [1,2,3,4]

db.Preload("order", "state NOT IN (?)", "cancelled").find(&user)
// select * from user
// select * from order where user_id IN [1,2,3,4] and state NOT IN (cancelled)

db.Where("state = ?", "active").Preload("order", "state NOT IN (?)", "cancelled").find(&user)
// select * from user  WHERE state = 'active'
// select * from order where user_id IN [1,2,3,4] and state NOT IN (cancelled)

```

## Nested Preload

```go

db.Preload("order").Preload("profile").Preload("role").Find(&users)
// select * from user
// select * from order where user_id IN [1,2,3,4]
// select * from profile where user_id IN [1,2,3,4]
// select * from role where user_id IN [1,2,3,4]

```

## Custom Preload

自訂 Preload

```go

db.Preload("order", func(db *gorm.DB) *gorm.DB {
    return db.Order("orders.amount DESC")
}).Find(&user)
// SELECT * FROM user;
// SELECT * FROM order WHERE user_id IN (1,2,3,4) order by orders.amount DESC;

```

## Auto Preload

自動 preload

```go

type User struct {
    gorm.Model
    Name       string
    CompanyID  uint
    Company    Company `gorm:"auto_preload"`
}

```