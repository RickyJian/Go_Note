# 關聯模式

用 Method 做關聯設定

## 開啟

```go

var account Account 

db.Model(&account).Association("Roles")

```

## 尋找關聯

```go

var account Account 
var role Role 

db.Model(&account).Association("Roles").Find(&role)

```

## Append 

新增新的 `Has Many`、`Many To Many` 來取代 `belongs to`、`Has One`

```go

var account Account 

db.Model(&account).Association("Roles").Append([]Role{adimin, supervisor})
db.Model(&account).Association("Roles").Append(Role{Name: "guest"})

```

## Replace 

取代舊關聯

```go

var account Account 

db.Model(&account).Association("Roles").Replace([]Role{adimin, supervisor})
db.Model(&account).Association("Roles").Replace(Role{Name: "guest"})

```

## Delete  

刪除 Model 與引數(argument)間的關聯

```go

var account Account 

db.Model(&account).Association("Roles").Delete([]Role{adimin, supervisor})
db.Model(&account).Association("Roles").Delete(Role{Name: "guest"})

```

## Clear  

刪除Model關聯

```go

var account Account 

db.Model(&account).Association("Roles").Clear()
db.Model(&account).Association("Roles").Clear()

```

## Count  

計算關聯

```go

var account Account 

db.Model(&account).Association("Roles").Count()
db.Model(&account).Association("Roles").Count()

```