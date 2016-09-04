package main

import "fmt"

/*
go语言的map类似其他语言中的哈希表或者字典，以key-value形式存储数据
Key必须是支持==或!=比较运算的类型，不可以是函数，map或slice
Map查找比线性搜索快很多，但比索引访问数据的类型慢100倍（可以使用slice和数组的时候尽量使用slice和数组，在不得已的情况下使用map）
Map使用make()创建，支持:=这种简写方式


*/
func test() {
	//创建map方式1
	var ma map[int]string
	ma = map[int]string{}
	fmt.Println(ma)

	//创建map方式2
	var mb map[int]string = make(map[int]string)
	fmt.Println(mb)

	//创建map方式3
	mc := make(map[int]string)
	fmt.Println(mc)
	mc[1] = "OK"
	a := mc[1]
	fmt.Println(mc)
	fmt.Println(a)
	//如果赋值给b的这个键没有值，那么b是等于空字符串的
	b := mc[2]
	fmt.Println(b == "")
	//这样为删除mc这个map的key为1的元素
	delete(mc, 1)
	fmt.Println(mc)

	//这种复杂的复合map，如果要 给m[1][1]进行赋值的话，那么必须得给一级map的value进行一个map初始化，否则会出现运行时错误
	md := make(map[int]map[int]string)
	md[1] = make(map[int]string)
	md[1][1] = "OK"
	d := md[1][1]
	fmt.Println(d)

	//多返回值解决未初始化问题，第二个返回值ok是用来判断是否me[1]的map为空，如果为空那么进行初始化操作，如果不为空那么直接进行赋值
	//每次进行赋值操作之前都需要进行一次检查
	me := make(map[int]map[int]string)
	e, ok := me[1][1]
	if !ok {
		me[1] = make(map[int]string)
	}
	me[1][1] = "Good"
	fmt.Println(e, ok)

	for k, v := range md {
		//这里同slice的使用相同，如果修改v值，那么不会影响到md的值，此处的v只是一个md元素的copy
		maptemp := make(map[int]string)
		maptemp[1] = "SUCCESS"
		v = maptemp
		fmt.Println(k, v)
		fmt.Println(k, md[k])
	}

	//按照这样的方式，并没有给sm成功赋值，因为v是sm元素的copy
	sm := make([]map[int]string, 5)
	for _, v := range sm {
		v = make(map[int]string, 1)
		v[1] = "OK"
		fmt.Println(v)
	}
	fmt.Println(sm)

	//按照如下方式的话，那么是可行的，sma会被赋值
	sma := make([]map[int]string, 5)
	//此处for之后写为i := range和i, _ := range都是可以的，但是后者会跑报警告
	for i := range sma {
		sma[i] = make(map[int]string, 1)
		sma[i][1] = "OK"
		fmt.Println(sma[i])
	}
	fmt.Println(sma)

	//课后练习，将map[int]string，翻转存为map[string]int类型。另外注意，每次运行map的顺序一般都不相同，从这也说明了map的无序性
	mi := make(map[int]string)
	mi[1] = "a"
	mi[2] = "b"
	mi[3] = "c"
	mj := make(map[string]int)
	for k, v := range mi {
		mj[v] = k
	}
	fmt.Println(mi)
	fmt.Println(mj)
}
