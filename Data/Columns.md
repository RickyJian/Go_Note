# 未知 Column

`rows.Columns()` 回傳 column name

```go

func main (){
    cols , err := rows.Columns()
    if err != nil {

    } else {
        dest := []interface{}{ 
            new(uint64), 
            new(string), 
            new(string), 
        }
    }
    vals := make([]interface {] , len(cols)})
    for i, _ := range cols {
        //  sql.RawBytes：若不知 column type 時，可以用此
        vals[i] = new(sql.RawBytes)
    }
    for rows.Next() {
        err = rows.Scan(vals...)
    } 
}


```