package main

import "fmt"

/*
Go中没有class，但依旧有method
通过显示说明receiver来实现与某个类型的组合
只能为同一个包中的类型定义方法
Receiver可以是类型的值或者指针
不存在方法重载
可以使用值或指针来调用方法，编译器会自动完成转换
从某种意义上来说，方法是函数的语法糖，因为receiver其实就是方法所接收的第一个参数
如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
类型别名不会拥有底层类型所附带的方法
方法可以调用结构中的非公开字段
*/

//对于同一个包来说，私有字段也可以被访问到，但是包外只能访问到公有字段
type testa struct {
	name string
}

type testb struct {
	Name string
}

type TZ int

//类型方法绑定必须是同一个包下的，否则找不到对应的类型
func test() {
	a := testa{}
	a.Print()
	fmt.Println(a.name)

	b := testb{}
	b.Print()
	fmt.Println(b.Name)

	//非常灵活，一个底层类型被type为自定义类型以后，一样可以通过进行方法绑定
	//这里也说明了为什么一定要进行强制类型转换
	var c TZ
	//这样写是method value
	c.Print()
	//这样写 也是可以的，method expression
	(*TZ).Print(&c)
}

//这里说明，如果是指针的话，那么是传址，如果是变量那么是传值
func (a *testa) Print() {
	//method有对struct所有字段的访问权限，包括开头小写的私有字段
	a.name = "AA"
	fmt.Println("Func A")
}

func (b testb) Print() {
	b.Name = "BB"
	fmt.Println("Func B")
}

func (c *TZ) Print() {
	fmt.Println("Func TZ")
}
