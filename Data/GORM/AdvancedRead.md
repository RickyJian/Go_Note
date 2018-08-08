
# AdvancedRead

## Order

```go

db.Order("id desc, name").Find(&users)
// select * from users ORDER BY id desc, name;

db.Order("id desc").Order("name").Find(&users)
// select * from users ORDER BY id desc, name;

```

## Limit

```go

db.Limit(5).Find(&users)
// select * from users limit 5

```

## Offset

```go

db.Offset(5).Find(&users)
// select * from users Offset 5

```

## Count

在 Query Chain 中只能放在最後

```go

db.Find(&users).Count(&count)
// select count(*) from users 

```

## Group & Having

```go

rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()
for rows.Next() {
    ...
}

rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
for rows.Next() {
    ...
}

type Result struct {
    Date  time.Time
    Total int64
}
db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)

```

## Query Chains

```go

db.Where("name = ?", "student" ).Not("id", 100 ).Find(&user)
// select * from users where name = "student" and id != 100 

```

## SubQuery

```go

db.Where("amount > ?", DB.Table("orders").Select("AVG(amount)").Where("state = ?", "paid").QueryExpr()).Find(&orders)

```

## FirstOrInit

* 若找不到資料就初始化 struct
* 只有 struct、map 中使用

```go

// struct
db.FirstOrInit(&user, User{Name: "初始化"})

// map
db.FirstOrInit(&user, map[string]interface{}{"name": "初始化"})

```

### Attr

若找不到資料就初始化 field

```go

db.Where(User{Name: "初始化"}).Attrs(User{Age: 20}).FirstOrInit(&user)

```

### Assign

不管搜尋出來的資料與否，就直接指派值給 field

```go

db.Where(User{Name: "初始化"}).Assign(User{Age: 20}).FirstOrInit(&user)

```

## FirstOrCreate

* 若找不到資料就將資料 insert 進 DB
* 只有 struct、map 中使用

```go

// struct
db.FirstOrCreate(&user, User{Name: "初始化"})

```

### Attr

若找不到資料就初始化 field ，並新增進 DB

```go

db.Where(User{Name: "初始化"}).Attrs(User{Age: 20}).FirstOrCreate(&user)

```

### Assign

* 若有資料，則直接指派值給 field 並更新 DB 資料
* 沒有資料，則新增 DB 資料
不管搜尋出來的資料與否，就直接指派值給 field

```go

db.Where(User{Name: "初始化"}).Assign(User{Age: 20}).FirstOrCreate(&user)

```

## Join

```go

rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
for rows.Next() {

}

db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

```

## Scan

將搜尋出的資料丟進 struct

```go

type User struct {
    ID int
    Name string
}

var user User
db.Table("User").Select("ID , Name").Where("ID = ?",3).Scan(&user)

db.Row("SELECT ID , Name FROM User Where ID = ?",3).Scan(&user)

```

## Scopes

預設搜尋

```go

func BetweenDate(db *gorm.DB) *gorm.DB{
    return db.Where("year < ? && year > ?", 1998,1996)
}

db.Scopes(BetweenDate).Find(&user)

```