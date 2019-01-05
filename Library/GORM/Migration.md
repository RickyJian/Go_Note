# Migration

db migrate

## Auto Migration

只會新建尚未新建的表，並不會修改異動的表、欄位或是資料

```go

db.AutoMigrate(&Account{})

db.AutoMigrate(&Account{},&Role)

```

## Talbe

### HasTalbe

檢查 table 是否存在

```go

// string
db.HasTable("user")

// model
db.HasTable(&User{})

```

### CreateTable

新建 Table

```go

db.CreateTable(&User{})

```

### DropTable

刪除 Table

```go

// model
db.DropTable(&User{})

// string
db.DropTable("user")

```

### DropTableIfExists

若 Table  存在則刪除

```go

db.DropTableIfExists(&User{})

```

## Column

### ModifyColumn

異動 Column 

```go

// 修改 model user `name` 的 data type 改為 `text` 
db.Model(&User{}).ModifyColumn("name", "text")

```

### DropColumn

刪除 Column 

```go

// 刪除 model user `description`
db.Model(&User{}).DropColumn("description")

```

### AddForeignKey

新建 FK

```go

// 新建 FK
// param1 : FK field
// param2 : table column
// param3 : ONDELETE
// param4 : ONUPDATE
db.Model(&User{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")

```

## Indexes

### AddIndex

新建 index

```go

db.Model(&User{}).AddIndex("idx_user_name", "name")
db.Model(&User{}).AddIndex("idx_user_name_age", "name", "age")

```

### AddUniqueIndex

```go

db.Model(&User{}).AddUniqueIndex("idx_user_name", "name")
db.Model(&User{}).AddUniqueIndex("idx_user_name_age", "name", "age")

```

### RemoveIndex

刪除 index

```go

db.Model(&User{}).RemoveIndex("idx_user_name")

```