package main

import "fmt"

//Slice本身不是数组，它指向底层的数组
//作为变长数组的替代方案，可以关联底层数组的局部或全部
//为一个引用类型
//可以直接创建或从底层数组获取生成
//一般使用make()创建Slice
//如果多个slice指向相同底层数组，那么一个改变，全部改变
func main() {
	//空的slice
	var sone []int
	fmt.Println(sone)

	//这里的stow是slice，只取了数组最后一个元素
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	stow := a[9]
	fmt.Println(stow)
	//这里相当于取[5 , 10)区间内的元素，包含5但不包含10
	sthree := a[5:10]
	fmt.Println(sthree)
	//如果不知道数组的全部长度，那么可以用以下两种方法
	sfour := a[5:]
	fmt.Println(sfour)
	sfive := a[5:len(a)]
	fmt.Println(sfive)
	//取数组的前5位
	ssix := a[:5]
	fmt.Println(ssix)

	//分配一个初始化3长度的数组，并预留空间长度为10，当元素个数超过3个时，不会重新分配地址
	//当元素个数超过10的时候，那么会重新分配连续内存地址
	sseven := make([]int, 3, 10)
	//此时得到的结果为3, 10，len获取的是元素个数，cap获取的是数组容量
	fmt.Println(len(sseven), cap(sseven))
	fmt.Println(sseven)

	//如果不给第三个参数，那么默认容量为数组当前长度，即输出结果为3，3
	seight := make([]int, 3)
	fmt.Println(len(seight), cap(seight))
	fmt.Println(seight)

	// sa切片，指向temp数组的[2, 5)号元素
	temp := [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa := temp[2:5]
	fmt.Println(string(sa))
	//此处元素个数为3，容量为9，其容量为c元素之后所有元素之和
	fmt.Println(len(sa), cap(sa))

	//下面这段代码输出zg，说明了slice切片实际上就是个指向数组的指针，当原数组改变时，slice切片指向的元素也会改变
	temp[5] = 'z'
	sb := sa[3:5]
	fmt.Println(string(sb))
	//若想sb切片，从sa切片中获取到de元素，那么如下所写，此切片是以sa切片为标准进行的
	sc := sa[1:3]
	fmt.Println(string(sc))
	/*该行代码发生了slice索引越界，不会导致底层数组的重新分配，而是引发错误，无法运行
	fmt.Println(sa[4])
	*/
	//这里给sa切片加了一个元素o，这个sa切片指针指向temp数组，而temp数组的第6位元素此时为z，那么最后sa切片增加了o元素
	//而temp数组的第6位元素被替换为了o
	sa = append(sa, 'o')
	fmt.Println(string(sa))
	fmt.Println(string(temp[5]))

	//以下代码说明了如果slice增加的元素个数小于等于它的容量个数，那么返回原地址，如果大于了它的容量个数，那么返回一个重新分配的地址
	sf := make([]int, 3, 6)
	fmt.Printf("%p", sf)
	sf = append(sf, 1, 2, 3)
	fmt.Printf("%v %p", sf, sf)
	sf = append(sf, 1, 2, 3)
	fmt.Printf("%v %p", sf, sf)
	fmt.Println("")

	//此处si和sj共享一个数组位置，如果改变sj[1]那么si[0]的值也会跟着改变
	tempb := []int{1, 2, 3, 4, 5}
	si := tempb[2:5]
	sj := tempb[1:3]
	fmt.Println(si, sj)
	sj[1] = 9
	fmt.Println(si, sj)

	//短数组copy到长数组中，那么头三位被覆盖
	tempc := []int{1, 2, 3, 4, 5, 6}
	tempd := []int{7, 8, 9}
	copy(tempc, tempd)
	fmt.Println(tempc)

	//长数组copy到短数组中，也只copy了头三位
	tempe := []int{1, 2, 3, 4, 5, 6}
	tempf := []int{7, 8, 9}
	copy(tempf, tempe)
	fmt.Println(tempf)

	//这样写也是可以的，只copy数组中的部分数据过去
	tempg := []int{1, 2, 3, 4, 5, 6}
	temph := []int{7, 8, 9}
	copy(temph[2:3], tempg[4:5])
	fmt.Println(temph)
}
