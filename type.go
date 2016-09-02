package main

import (
	"fmt"
	"strconv"
)

//Go语言类型
func testa() {
	//冒号相当于替代了var关键字，并且其后的变量是可推断赋值
	b, c := 1*1.2, "测试"
	d := 36 * 1.2

	fmt.Println(b)
	fmt.Println(c)

	//这里_表示忽略掉8这个返回值。在多返回值函数的时候会用上
	e, _, f, g := 7, 8, 9, 10

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	//Go语言不存在隐式类型转换，是一门类型安全的语言
	//转换只能发生在两种相互兼容的类型之间
	var testFloat float32 = 3.14
	testInt := int(testFloat)
	fmt.Println(testInt)

	/*此处编译不通过
	  var testbool bool = true
	  var testInt2 int = int(testbool)
	  fmt.Println(testInt2);
	*/

	//计算机中存储任何东西本质上都是数字，因此此函数自然认为我们需要的是数字65表示字符A
	var testInt3 int = 65
	var testString string = string(testInt3)
	fmt.Println(testString)

	//在Stirng和int类型之间相互转换的时候，要用到strconv包中的函数
	var testString2 string = strconv.Itoa(testInt3)
	fmt.Println(testString2)
	//此处需要注意的点为，Atoi函数返回多个值，int, error。那么第二个error值需要用_来进行忽略
	var testInt4, _ = strconv.Atoi(testString2)
	fmt.Println(testInt4)
}
