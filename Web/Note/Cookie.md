# cookie

* 儲存在 client 端的 browser 裡
* 時間短暫

```go

cookie := http.Cookie{
    Name:  "_cookie"
    Value: session.Uuid
    HttpOnly: true
}

func auth(w http.ResponseWriter , r *http.Request){
    http.SetCookie(w,&cookie)
}

```

## Name

cookie 名稱，能隨意設置

## Value 

儲存在 browser 裡的唯一ID

## HttpOnly

設為 true 時 cookie 只能通過 HTTP 或 HTTPS 訪問