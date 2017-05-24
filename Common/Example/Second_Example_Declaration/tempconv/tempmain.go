package main

import (
	"gopl.io/ch2/tempconv"
	"fmt"
)

func main(){
	c := tempconv.FToC(212.0)
	fmt.Printf(c.String())
}