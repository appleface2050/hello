package main

import "fmt"
import (
// "sort"
)

type TZ int

func (tz *TZ) Increase(num int) {
	*tz += TZ(num)
}

func main() {
	// a := TZ
	var a TZ
	a.Increase(100)
	fmt.Println(a)
}
