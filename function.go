package main

import "fmt"

/*
Go函数不支持嵌套、重载和默认参数
支持以下特性：
  无需声明原型、不定长度变参、多返回值、命名返回值参数、匿名函数、闭包
  函数也可以作为一种类型使用
defer
  执行方式类似其他语言中的析构函数，在函数体执行结束后，按照调用顺序的相反顺序逐个执行
  即使函数发生严重错误也会执行
  支持匿名函数的调用
  常用语资源清理、文件关闭、解锁以及记录时间等操作
  通过与匿名函数配合可在return之后修改函数计算结果
  如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer时即已经获得了拷贝，否则则是引用某个变量的地址
  Go没有异常机制，但有panic/recover模式来处理错误
  Panic可以在任何地方引发，但recover只有在defer调用的函数中有效
*/
func test() {
	A()
	//Go语言中一切皆是类型，函数也是类型
	b := B
	b()

	//匿名函数
	c := func() {
		fmt.Println("Func C")
	}
	c()

	closure := Closure(10)
	fmt.Println("闭包函数result：", closure(12))

	/**
	//defer中的Println函数被逆序调用，输出2,1,0
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	//这个defer的执行方式类似其他语言中的析构函数，在函数体执行结束后按照调用顺序的相反顺序逐个执行
	//所以最后输出结果也是------------在defer这两个函数的输出之上
	//按照顺序此处defer输出的是3,3,3。主要是因为闭包的问题，在上面循环中，fmt.Println()中，i是传值的
	//而在下面的循环中，defer中的匿名函数，使用父函数中的main的i值，那么此处是传址。在这里Println输出的是i最后循环结束后的值，为3
	fmt.Println("-----------------------------")
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	*/
	D()
	E()
	G()
}

//此种情况为命名返回值参数
func A() (a, b, c int) {
	fmt.Println(a, b, c)
	return a, b, c
}

func B() {
	fmt.Println("Func B")
}

//闭包的使用
func Closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

//以下函数都是为panice和recover做演示
func D() {
	fmt.Println("Func D")
}

func E() {
	//这里defer在函数运行结束后，不论是否发生错误都会执行，通过recover()函数，判断是否发生了panic，如果err!=nil，那么说明确实发生了Panic
	//那么执行对应的操作，并恢复程序的正常执行。
	//虽然Go没有异常捕捉处理机制，但是此处的panic和recover非常类似，panic如同Java的抛出异常throw Exception，而recover为Java中的catch，捕获Exception异常
	//此处特别需要注意的是，defer一定要写在panic之前，因为程序发生panic之后，后面的代码段无效，不会被继续执行，那么defer放在panic之前才会被注册，最后函数return时才会被调用
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in E")
		}
	}()
	panic("Panic in E")
}

func G() {
	fmt.Println("Func G")
}
