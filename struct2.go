package main

import (
	"fmt"
)

type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age  int
}

func main() {
	a := teacher{Name: "shang", Age: 17, human: human{Sex: 11}}
	a.Name = "qq"
	// a.human = human{Sex: 99}
	a.human.Sex = 99
	a.Sex = 99
	fmt.Println(a)

}
