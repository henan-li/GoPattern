package main

import "fmt"

// 承上启下的装饰器
type Decorator struct {
	drinks Drinks
}

func (this *Decorator) drinksPrice() int {
	return this.drinks.drinksPrice()
}

func (this *Decorator) drinksDesc() {
	this.drinks.drinksDesc()
}

// 具体装饰
type milk struct {
	Decorator
}

func (this *milk) drinksPrice() int {

	return this.Decorator.drinksPrice() + 10
}
func (this *milk) drinksDesc() {
	fmt.Println("this is milkCoffee")
}

// 具体装饰
type lemon struct {
	Decorator
}

func (this *lemon) drinksPrice() int {
	return this.Decorator.drinksPrice() + 10
}
func (this *lemon) drinksDesc() {
	fmt.Println("this is lemonSoda")
}
