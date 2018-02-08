package main

import (
	"fmt"
)

func main() {
	a := &struct {
		Name string
		Age  int

		Contact struct {
			Phone, City string
		}
	}{
		Name: "shang",
		Age:  17,
	}

	fmt.Println(*a)
}
