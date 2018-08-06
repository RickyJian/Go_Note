# Read

## 基本

### First

根據 PK 排序並獲取第一筆資料

```go

// 不帶參數
db.First(&user)
// select * from users order by id limit 1

// 但參數，但只對 PK 型態為 int 才起作用
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

db.Where("name = ?", "student").Find(&users)
// SELECT * FROM users WHERE name = 'student'

db.Where("id <> ?", "100").Find(&users)
// SELECT * FROM users WHERE id <> '100'

```

### IN

```go

db.Where("name in (?)", []string{"student","student1"}).Find(&users)

```

### Like

```go

db.Where("name like ?", "%tud%").Find(&users)

```

### And

```go

db.Where("name = ? AND id = ?", "student" , "100").Find(&users)

```
