package main

import "fmt"

/*
执行结果分析：
先输出非defer的，即增强for循环中的f()。因为在第一个for循环中，通过匿名函数对fs[i]进行了赋值，所以这里使用的是传址，那么i的值最后为4，输出结果也为4
然后按照defer的先进后出，后进先出的顺序输出，那么最先输出defer_closure i = 4，这个匿名函数是最后进去的，此处也是传址，为4
而后输出fmt.Println中的defer i = 3;此处不是匿名函数，为传值。此后非匿名函数和匿名函数交叉输出
*/
func test() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() {
			fmt.Println("defer_closure i = ", i)
		}()
		fs[i] = func() {
			fmt.Println("closure i = ", i)
		}
	}

	for _, f := range fs {
		f()
	}
}
