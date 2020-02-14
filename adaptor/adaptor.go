package main

import "fmt"

type advanceCharger interface {
	chargerOtherDevice(phoneBrand string)
}

// 适配器，在兼容的基础上，还有额外功能,这些额外功能的管理可以交给一个接口
type adaptor struct {
	iphoneCharger
	androidCharger
}

func (this *adaptor) charge(phoneBrand string) {
	switch phoneBrand {
	case "iphone":
		this.iphoneCharger.charge()
	case "android":
		this.androidCharger.charge()
	default:
		this.chargerOtherDevice(phoneBrand)
	}
}

func (this *adaptor) chargerOtherDevice(phoneBrand string) {
	fmt.Println("adaptor can charge other device: ", phoneBrand)
}
