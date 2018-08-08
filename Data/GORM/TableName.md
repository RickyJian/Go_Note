# TableName

指定 TableName

```go

// 以 struct `User` 新建 table User 
db.Table("User").CreateTable(&User{})

var users []User
db.Table("User").Find(&users)
//// SELECT * FROM User;

db.Table("User").Where("name = ?", "student").Delete()
//// DELETE FROM User WHERE name = 'student';

```