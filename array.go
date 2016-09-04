package main

import "fmt"

//在go语言中，函数间直接的数组传递，是值传递，并不是址传递
func test() {

	/*这里是会报错的，在go中，数组长度不同的[2]int和[1]int相当于是两个类型，不能相互赋值，必须使用循环
	var a [2]int
	var b [1]int
	b = a
	*/

	//此处相当于给第一个元素赋值为1，其他元素为默认值0
	a := [2]int{1}
	fmt.Println(a)

	//此处相当于给下标为19的元素，即第20个元素赋值1，其他元素默认值为0
	b := [20]int{19: 1}
	fmt.Println(b)

	//这里相当于由系统来判断数值的大小。
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println(c)

	//这里系统知道最少有20个元素，那么数组长度为20
	d := [20]int{19: 1}
	fmt.Println(d)

	//此处为指向数组的指针
	var pd *[20]int = &d
	fmt.Println(pd)

	//此处为指针数组，一个数组里放的是指针
	x, y := 1, 2
	e := [...]*int{&x, &y}
	fmt.Println(e)

	//此处返回true，可以直接进行!=和==比较
	f := [2]int{1, 2}
	g := [2]int{1, 2}
	fmt.Println(f == g)

	/*此处不能进行比较，因为这相当于是两种类型之间的比较，go语言并不支持
	h := [1]int{1}
	g := [2]int{1, 2}
	fmt.Println(h == g)
	*/

	//如果使用new关键字，则创建完int数组后返回的是一个指向数组的指针
	//两种方式创建的数组，都可以用下标的方式进行直接赋值
	i := [10]int{}
	i[1] = 2
	fmt.Println(i)
	p := new([10]int)
	p[1] = 2
	fmt.Println(p)

	//go语言支持多维数组
	j := [2][3]int{
		{1, 1, 1},
		{2, 2, 2}}
	fmt.Println(j)

	//很多对于一维数组的特性，也适用于二维数组
	l := [2][3]int{
		{1: 1},
		{2: 2}}
	fmt.Println(l)

	sort()
}

//go语言版的冒泡排序
func sort() {
	a := [...]int{5, 2, 6, 3, 9}
	fmt.Println(a)

	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}
