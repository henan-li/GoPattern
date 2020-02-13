package main

import "fmt"

func main() {

	// original sources
	c := &Coffee{}
	s := &Soda{}

	fmt.Println("***** original coffee ******")
	c.drinksDesc()
	price := c.drinksPrice()
	fmt.Println(price)

	fmt.Println("***** original soda ******")
	s.drinksDesc()
	price = s.drinksPrice()
	fmt.Println(price)

	// put sources into cup
	coffeeCup := Decorator{c}
	sodaCup := Decorator{s}

	fmt.Println("***** coffee in cup ******")
	coffeeCup.drinksDesc()
	price = coffeeCup.drinksPrice()
	fmt.Println(price)

	fmt.Println("***** soda in cup ******")
	sodaCup.drinksDesc()
	price = sodaCup.drinksPrice()
	fmt.Println(price)

	// add more source into cup
	m := &milk{Decorator{c}}
	l := &lemon{Decorator{s}}

	fmt.Println("***** add milk in coffee cup ******")
	m.drinksDesc()
	price = m.drinksPrice()
	fmt.Println(price)

	fmt.Println("***** add lemon in soda cup ******")
	l.drinksDesc()
	price = l.drinksPrice()
	fmt.Println(price)

}
