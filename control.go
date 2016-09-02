package main

import "fmt"

func main() {
	//指针，和C语言用法相同。但是不支持指针运算
	a := 1
	var p *int = &a
	fmt.Println(*p)

	/*这里是有问题的，在go中++和--并不是一个表达式运算，不能进行运算赋值
	var c int = 1--;
	*/
	//单纯做一个运算是可以的
	a++
	fmt.Println(a)

	//if 语句，此处是可以不加括号的，加上括号也可以，在格式化时会被去掉
	if a < 1 {

	}

	//此处b的作用于只在if语句块中，这里总算用到了分号，我还以为这个语言没有分号的。哈哈哈哈
	if b := 2; b < 3 {
		fmt.Println(b)
	}

	if a := -1; a < 0 {
		fmt.Println("此处的a是if语句块中的a，外部的a被隐藏掉了")
	}
	fmt.Println("此处的a是if语句块外的a")

	//这样也是可以的
	if a, b := 1, 2; a < 3 {
		fmt.Println(b)
	}

	//支持无条件for循环，无限循环，相当于while(true)
	c := 1
	for {
		c++
		if c > 3 {
			fmt.Println("over")
			break
		}
	}

	/*相当于while(condition)
	for c < 5 {
	}
	*/

	//for循环的经典形式
	for index := 0; index < 10; index++ {
	}

	//这样做是不合适的，因为for循环在每次检查条件时都会再运行一次len(str)表达式，影响循环速度
	//合理的做法是，lenght := len(str);先计算出lenght后，再放入循环条件中
	str := "test"
	for index := 0; index < len(str); index++ {
	}

	i := 3
	//这里的i支持任何基本类型，需要注意的是go语言中是可以不写break的，默认含有break，如果在判断完case条件后还进行后续case判断
	//那么使用fallthrough关键字
	switch i {
	case 1:
		fmt.Println("switch1")
		fallthrough
	case 2:
		fmt.Println("switch2")
	case 3:
		fmt.Println("switch3")
	default:
		fmt.Println("无匹配结果")
	}

	//switch后面可以不跟变量，直接在case中进行条件判断，其结构有点类似if else if
	b := 3
	switch {
	case a < 10:
		fmt.Println("a < 10")
	case b < 10:
		fmt.Println("b < 10")
	default:
		fmt.Println("所有条件不通过时default")
	}

	//这样做也是可以的，十分灵活
	switch d := 3; {
	case d < 3:
		fmt.Println("test")
	}

	//跳转语句break, continue的用法和Java相同，可以break和continue到指定的标签位置
LABEL:
	for {
		for {
			break LABEL
		}
	}

	fmt.Println("---------------------")
	//此处的Continue不能换为goto，goto为调整程序的执行位置，如果换为goto，那么将会出现死循环，goto的标签应该放在goto语句之后
LABEL2:
	for index := 0; index < 3; index++ {
		for {
			fmt.Println(index)
			continue LABEL2
		}
	}

}
