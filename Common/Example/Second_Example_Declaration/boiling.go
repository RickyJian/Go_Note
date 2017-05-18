package main

import (
	"fmt"
)

const boilingF = 212.0 //套件層級宣告

func main() {
	var f = boilingF                                   // 區域變數
	var c = (f - 32) * 5 / 9                           // 區域變數
	fmt.Printf("boiling point = %g F or %g C\n", f, c) // %g 浮點樹輸出
}
