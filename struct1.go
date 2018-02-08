package main

import (
	"fmt"
)

type person struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

func main() {
	a := person{Name: "shang", Age: 17}
	a.Contact.City = "aq"
	a.Contact.Phone = "1234567"
	fmt.Println(a)
}
