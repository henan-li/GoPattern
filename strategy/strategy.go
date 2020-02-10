//// 策略模式: 条条大路通罗马
package main
//
//import "fmt"
//
//// 目的地
//type Rome interface {
//	GotoRome() string
//}
//
//// 用什么策略去目的地
//type car struct{}
//
//func (*car) GotoRome() string {
//	method := "take car to rome"
//	return method
//}
//
//// 用什么策略去目的地
//type walk struct{}
//
//func (*walk) GotoRome() string {
//	method := "walk to rome"
//	return method
//}
//
//type Person struct {
//	gotoRomeMethod Rome
//}
//
//// 参数是接口类型，go中接口的实现是隐式，即实现接口定义的方法，就实现了该接口
//// 选择策略
//func (this *Person) setStrategy(romeMethod Rome) {
//	this.gotoRomeMethod = romeMethod
//}
//
//// 执行策略
//func (this *Person) goWithMethod() string {
//	return this.gotoRomeMethod.GotoRome()
//}
//
//func main() {
//	p := &Person{}
//
//	p.setStrategy(&car{})
//	res := p.goWithMethod()
//	fmt.Println(res)
//
//	p.setStrategy(&walk{})
//	res = p.goWithMethod()
//	fmt.Println(res)
//}
//
//// comment out this whole code if you want to run theme example
//
//
