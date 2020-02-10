package main

import "fmt"

type operator struct {
	theme
}

func (this *operator) choseColorAndSize(t *theme) {
	this.theme.color = t.color
	this.theme.size = t.size
}
func (this *operator) changeTheme() string {
	color := this.theme.color.changeColor()
	size := this.theme.size.changeSize()
	return color + " , " + size
}

func main() {

	p := &operator{}

	p.choseColorAndSize(&theme{&greenColor{}, &smallSize{}})
	res := p.changeTheme()
	fmt.Println(res)

	p.choseColorAndSize(&theme{&redColor{}, &bigSize{}})
	res = p.changeTheme()
	fmt.Println(res)
}
