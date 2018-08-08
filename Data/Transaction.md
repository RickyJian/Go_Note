# 交易管理

```go

func main(){
    tx , err := db.Begin()
    if err != nil {

    }
    defer tx.Rollback()
    stmt  err := tx.Prepare("insert into user values(?)")
    if err != nil {
        
    }
    result , err = stmt.Exec(1)
    defer stmt.Close() 
    err = tx.Commit()
    if err != nil {
        
    }
}

```