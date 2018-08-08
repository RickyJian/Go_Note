# 查詢

## Query

```go

func main (){
    var id int
    var name string 
    rows , err  := db.Query("select id , name from users where id = ?", 1 )
    if err != nil {

    }
    defer rows.Close()
    // Next：將搜尋出來的資料撈出來做處理
    for rows.Next(){
        // Scan：將資料寫入 id 和 name 
        err := rows.Scan(&id, &name)
        if err != nil {

        }
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
}


```

## QueryRow


```go

func main (){
    var id int
    var name string 
    err = db.QueryRow("select id , name from users where id = ?", 1).Scan(&id, &name)
    if err != nil {

    }
   
}


```
