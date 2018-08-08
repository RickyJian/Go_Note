# Read

搜尋

## Select

```go

db.Select("name, age").Find(&users)
// SELECT name, age FROM users;

db.Select([]string{"name", "age"}).Find(&users)
// SELECT name, age FROM users;

```

## FROM

### First

根據 PK 排序並獲取第一筆資料

```go

// 不帶參數
db.First(&user)
// select * from users order by id limit 1

// 帶參數，但只對 PK 型態為 int 才起作用
db.First(&user, 10)
// SELECT * FROM users WHERE id = 10;

```

### Last

根據 PK 排序獲取最後一筆資料

```go

db.Last(&user)
// select * from users order by id DESC limit 1

```

### Find

全撈

```go

db.Find(&user)
// select * from users

```

## Where

### 一般

```go

// inline 
db.Where("name = ?", "student").Find(&users)
// SELECT * FROM users WHERE name = 'student'

// inline
db.Where("id <> ?", 100).Find(&users)
// SELECT * FROM users WHERE id <> '100'

// struct
db.Where(&User{Name: "student", id: 100}).Find(&user)
// Select * from users where name = 'student' and id = 100

// Map
db.Where(map[string]interface{}{"name": "student", "id": 100}).Find(&users)
// Select * from users where name = 'student' and id = 100


```

### IN

```go

// 以 name 做 in
db.Where("name in (?)", []string{"student","student1"}).Find(&users)

// 以 PK 做 in
db.Where([]int64{20, 21, 22}).Find(&users)

```

### Like

```go

db.Where("name like ?", "%tud%").Find(&users)

```

### And

```go

db.Where("name = ? AND id = ?", "student" , "100").Find(&users)

```

### Or

```go

db.Where("name = ? " , "student").Or("id = ? ", 100).Find(&users)
// select * from users where name = "student" or id = 100

```

### Not

```go

db.Not("name", "student").Find(&user)
// select * from users where name <> "student"

```

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