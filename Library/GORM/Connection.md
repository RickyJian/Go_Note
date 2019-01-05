# Connection

SQL DRIVER：[連結](https://github.com/golang/go/wiki/SQLDrivers)

## MSSQL

```go

package main 

import (
    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// username: 帳號
// password: 密碼
// dbname: database 名稱
func main (){
    db, err = gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=dbname")
    defer db.Close()
}

```

## MySQL

```go

package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"    
)

// user: 帳號
// password: 密碼
// dbname: database 名稱
func main(){
    db , err := gorm.Open("mysql","user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
    defer db.Close()
}

```

## PostgreSQL

```go

package main 

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"    
)


func main(){
    db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
    defer db.Close()
}

```

## Sqlite

```go

package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"   
)

func main(){
    db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
    defer db.Close()    
}

```