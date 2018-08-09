# 懶初始

`sync.Once`：一次性初始化

```go

var (
    loadIconsOnce   sync.Once
    icons map[string]image.Image
)

func loadIcons(){
    icons = map[string]image.Image{
        "1.png":loadIcon("1.png")
        "2.png":loadIcon("2.png")
        "3.png":loadIcon("3.png")
    }
}

func Icon(name string) image.Image{
    // Do：接受初始化函式作為他參數
    loadIconsOnce.Do(loadIcons)
    return icons[name]
}

```