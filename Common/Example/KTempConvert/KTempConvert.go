package main

import (
	"fmt"
)

type Kelvin float64

func (k Kelvin) String() string {
	return fmt.Sprintf("%g K \n", k)
}

func CToK(c float64) {
	k := c + 273.15
	fmt.Printf("%g K \n", k)
}

func (k Kelvin) KToC() {
	c := k - 273.15
	fmt.Printf("%g C \n", c)
}

var KTemp Kelvin

func main() {
	KTemp = 5
	fmt.Printf("%g K\n", KTemp)
	CToK(100)
	KTemp.KToC()
}
