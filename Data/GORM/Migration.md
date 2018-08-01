# Migration

db migrate

## Auto Migration

只會新建尚未新建的表，並不會修改異動的表、欄位或是資料

```go

db.AutoMigrate(&Account{})

db.AutoMigrate(&Account{},&Role)

```