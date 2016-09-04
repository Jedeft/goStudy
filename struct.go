package main

import "fmt"

/*
Go中的struct与C中的struct非常相似，并且Go没有class
使用type<Name> struct{}定义结构，名称遵循可见性规则
支持指向自身的指针类型成员
支持匿名结构，可用作成员或定义成员变量
匿名结构也可以用于map的值
可以使用字面值对结构进行初始化
允许直接通过指针来读写结构成员
相同类型的成员可进行直接拷贝赋值
支持==与!=比较预算符，但不支持>或<
支持匿名字段，本质上是定义了以某个类型名为名称的字段
嵌入结构作为匿名字段看起来像继承，但不是继承
可以使用匿名字段指针
*/

//struct初探
type person struct {
	Name string
	Age  int
}

//匿名结构体嵌套
type consumer struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

//结构体匿名字段
type student struct {
	string
	int
}

//以下是struct的组合使用
type human struct {
	Sex int
}

type doctor struct {
	human
	Name string
	Age  int
}

type patient struct {
	human
	Name string
	Age  int
}

func test() {
	a := person{}
	a.Name = "Jedeft"
	a.Age = 22
	fmt.Println(a)

	//简便初始化
	b := person{
		Name: "Jerry",
		Age:  13,
	}
	fmt.Println(b)

	//此处是进行一个值拷贝
	A(b)
	fmt.Println(b)

	//此处进行一个址拷贝
	B(&b)
	fmt.Println(b)

	//go语言中，推荐结构初始化时进行取地址符，那么在函数中都为传址操作。并且在go语言中的struct结构体是否取地址符，对于结构体中属性的改变。变量操作都是相同的
	c := &person{
		Name: "Tom",
		Age:  43,
	}
	c.Age = 88
	//这里是输出，指向{Tom,88}这个结构体的地址，所以前面会有&
	fmt.Println(c)
	//这里是输出{Tom,88}这个结构体
	fmt.Println(*c)

	//匿名struct
	d := &struct {
		Name string
		Age  int
	}{
		Name: "Joe",
		Age:  99,
	}
	fmt.Println(d)

	//匿名结构体支持嵌套
	e := consumer{
		Name: "consumer-Wang",
		Age:  55,
	}
	e.Contact.Phone = "123456789"
	e.Contact.City = "雁塔区"
	fmt.Println(e)

	//当使用匿名字段结构体时，要严格遵守结构体中变量类型的顺序
	f := student{"student", 10}
	fmt.Println(f)

	//结构体之间是可以相互赋值和比较的，但是比较必须是相同类型之间才能发生比较
	var g student
	g = f
	fmt.Println(g)
	fmt.Println(g == f)

	//结构体嵌入实现struct组合
	h := doctor{Name: "doctor", Age: 55, human: human{Sex: 0}}
	i := patient{Name: "patient", Age: 28, human: human{Sex: 1}}
	fmt.Println(h, i)
	h.Name = "doctor2"
	h.Age = 13
	/*这样写是完全可以的，这种写法是为了解决struct嵌套时命名冲突问题的
	  h.human.Sex = 99
	*/
	//这样写也对，相当于human中的属性已经嵌套到doctor中了，可以直接对其进行操作
	h.Sex = 99
	fmt.Println(h)
}

//值拷贝函数
func A(per person) {
	per.Age = 19
	fmt.Println("B", per)
}

//址拷贝函数
func B(per *person) {
	per.Age = 19
	fmt.Println("B", per)
}
