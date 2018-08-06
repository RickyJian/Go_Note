# range

* 迭代出 array 、 slice 的值及 index
* 迭代出 map 的 key 與 value

```go

ints := []int{1,2,3,4,5}

for i , d := range ints {
    fmt.Printf("%d: %d\n" ,i ,d)
}

```