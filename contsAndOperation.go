package main

import (
	"fmt"
)

const a, b, c = 1, 2, 3
const (
	d, e, f = 4, 5, 6
)

//此处常量test2和test3取test1的值。
const (
	test1 = 1
	test2
	test3
)

//常量也可以根据已有值和内置函数来得到新的值，但是不可以根据运行时才能获得的值进行赋值
const (
	test4 = "abc"
	test5 = len(test4)
	test6
)

/*这里这样写是无法通过编译的
const (
	a1, b1 = 1, "2"
	d1
)
*/

//这样是可行的，已被赋值的常量和未赋值的常量要保持数量一致
const (
	a1, b1 = 1, "2"
	c1, d1
)

//iota关键字是做自增长使用，同样也可以以iota + 1这种方式实现从1开始自增长
const (
	Sunday = iota
	Monday
	Toesday
)

const (
	cMAXCOUNT = iota + 1
	cMINCOUNT
)

//常量及运算符号
func testb() {
	fmt.Println("123")
	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)

	fmt.Println(test4)
	fmt.Println(test5)
	fmt.Println(test6)

	fmt.Println(cMAXCOUNT)
	fmt.Println(cMINCOUNT)

	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Toesday)

	//输出结果为1024，将1左移10位，二进制
	fmt.Println(1 << 10)

	//输出结果为0，将1右移10位，最末位为0，已是最低位
	fmt.Println(1 >> 10)

	test := 1
	if test > 0 && 10/test > 0 {
		fmt.Println("OK")
	}

}
