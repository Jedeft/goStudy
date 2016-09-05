package main

import "fmt"

/*
接口是一个或多个方法签名的集合
只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需声明实现了哪个接口，这称为Structural Typing
接口只有方法声明，没有数据字段
接口可以匿名嵌入其他接口，或嵌入到结构中
将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也获取指针
只有当接口存储的类型和对象都为nil时，接口才等于nil
接口调用不会做receiver的自动转换
接口同样支持匿名字段方法
接口也可实现类似OOP中的多态
空接口可以作为任何类型数据的容器
*/

//此处empty的空interface就相当于Java的Object
//go语言中没有继承概念，但是相当于任何类型都实现了这个空接口
type empty interface{}

type USB interface {
	Name() string
	Connecter
}

//嵌套接口
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
	fmt.Println("Connect :", pc.name)
}

func test() {
	var a USB
	//这里只能进行字面值初始化，不能a.name=***，因为interface没有name字段
	a = PhoneConnecter{name: "PhoneConnecter"}
	a.Connect()
	DisConnect(a)
	DisConnectTow(a)

	pc := PhoneConnecter{"PhoneConnecter"}
	var c Connecter
	c = Connecter(pc)
	c.Connect()
	/*这里找不到Name方法
	c.Name()
	*/

	//这里完全无视对name的修改，所以c = Connecter(pc)
	//将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针
	pc.name = "pc"
	//这里输出的还是PhoneConnecter
	c.Connect()
	//这里输出的就是pc
	pc.Connect()

	var d interface{}
	//输出true，真的为nil,当存储的类型和对象都为nil时，输出nil
	fmt.Println(d == nil)
	var p *int = nil
	d = p
	//输出false，虽然存储的类型为nil但对象为p，是一个指向nil的指针。此处为false
	fmt.Println(d == nil)
}

//如果这里是空接口，那么这里可以传递进来任何类型
func DisConnect(usb interface{}) {
	//简单的类型调研
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("DisConnected : ", pc.name)
		return
	}
	fmt.Println("UnKnone decive")
}

//类型调研时，用switch会更加简洁
func DisConnectTow(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("DisConnected : ", v.name)
	default:
		fmt.Println("UnKnone decive")
	}
}
