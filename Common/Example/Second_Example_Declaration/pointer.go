package main

import (
	"fmt"
)

func main (){
	x := 1
	p := &x
	fmt.Printf("%d \n",p)
	*p = 4
	fmt.Printf("%d", *p)

}