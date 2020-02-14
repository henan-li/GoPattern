package main

import "fmt"

type iphoneCharger struct {
}

func (*iphoneCharger) charge() {
	fmt.Println("charge iphone")
}

type androidCharger struct {
}

func (*androidCharger) charge() {
	fmt.Println("charge android")
}
