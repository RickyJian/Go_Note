# error 介面

```go

type error interface {
    Error() string
}

```

## 建構 error

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
