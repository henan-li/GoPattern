package main

type theme struct {
	color themeColor
	size  themeSize
}

// change color
type themeColor interface {
	changeColor() string
}

type redColor struct{}
func (*redColor) changeColor() string {
	color := "red"
	return "color is " + color
}

type greenColor struct{}
func (*greenColor) changeColor() string {
	color := "green"
	return "color is " + color
}

// change size
type themeSize interface {
	changeSize() string
}

type smallSize struct{}
func (*smallSize) changeSize() string {
	return "size is changed to small"
}

type bigSize struct{}
func (*bigSize) changeSize() string {

	return "size is changed to big"
}