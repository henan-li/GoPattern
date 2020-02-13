package main

import "fmt"

type Drinks interface {
	drinksDesc()
	drinksPrice() int
}

type Coffee struct {
}

func (*Coffee) drinksDesc() {
	fmt.Println("this is coffee")
}
func (*Coffee) drinksPrice() int {
	return 10
}

type Soda struct {
}

func (*Soda) drinksDesc() {
	fmt.Println("this is soda")
}
func (*Soda) drinksPrice() int {
	return 5
}
