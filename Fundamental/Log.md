# Log

## Log 

紀錄正常資訊

```go

t.Log("Test...")

t.Logf("Test...%d",1)

```

## Error

記錄錯誤訊息

```go

t.Error("Error...")

t.Errorf("Error...%d",1)

```

## Fatal

紀錄致命錯誤

> 致命錯誤：錯誤導致程式無法執行

```go

t.Fatal("Fatal...")

t.Fatalf("Fatal...%d",1)

```

## Fail

標記測試是失敗的，但不會終止目前測試函示執行

```go

t.Fail("Fail...")

```

## FailNow

標記測試是失敗的，立即終止目前測試函示執行

```go

t.FailNow("FailNow...")

```

## Failed

判斷測試函式是否被標記為失敗

```go

t.Failed()

```

## SkipNow

忽略測試函式

```go

t.SkipNow()

```

## Skipped

判斷測試函式是否被標記為忽略

```go

t.Skipped()

```
