package main

import "fmt"

type USB interface {
	Name() string
	Connect
}

type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("connect:", pc.Name())

}

func main() {
	var a USB
	a = PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	str, ok := a.(PhoneConnecter)
	fmt.Println(str, ok)
	Disconnect(a)
}

func Disconnect(usb USB) {
	fmt.Println("Disconnect")
}
