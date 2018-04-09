# error

為一個介面型別，常用在可預期錯誤

```go

type error interface {
    Error() string
}

```

## errors package

```go

package errors

func New(text string) error {
    return &errorString{text}
}

type errorString struct {
    text string
}

func (e *errorString) Error() string {
    return e.text
}

```

## fmt.Errorf

並不會在 console 端輸出，而是用來初始化一個 error 型態值並作為該函式的結果回傳給呼叫方

```go

err := fmt.Errorf("%s\n", "ERROR")

```