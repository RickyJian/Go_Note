# 增改查

## Exec

```go

func main(){
    // Prepare：創建 statement 待之後 Exec 使用
    stmt , err := db.Prepare("Insert Into User(id,name) values(?,?)")
    if err != nil {

    }
    // 執行 SQL
    result , err := stmt.Exec(1,"name")
    if err != nil {

    }
}

```