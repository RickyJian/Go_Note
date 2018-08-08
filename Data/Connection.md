# Connection

## MSSQL

```go

func main(){
    db , err := sql.Open("mssql","sqlserver://sa:mypass@localhost:1433?database=master&connection+timeout=30 ")
        if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
}

```

## MySQL

```go

func main() {
    db, err := sql.Open("mysql","user:password@tcp(127.0.0.1:3306)/hello")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
}

```

## 驗證連線

```go

err = db.Ping()
if err != nil {
    
}

```