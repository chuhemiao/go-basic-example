package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
)

var c, php, golang bool

var i, j int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// 声明常量

const sc string = "hello chuhemiao"
const pi float32 = 3.1415926

// 定义一个结构体

type Sex struct {
	X int
	Y int
}

// 映射

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

// 函数之 方法函数

type VertexFunc struct {
	Xv, Yv float64
}

func (vv VertexFunc) Abs() float64 {
	return math.Sqrt(vv.Xv*vv.Xv + vv.Yv*vv.Yv)
}

func main() {

	k := 3
	erlang, js, python := true, false, "no!"

	// 零值
	var i int
	var f float64
	var b bool
	var s string

	// for 循环

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	// 简化版for循环
	sums := 0
	for i := 0; i < 99; i++ {
		sums += i

	}

	var kk bool

	if sums > 4850 {
		kk = true
	} else {
		kk = false
	}

	// 声明指针
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值

	// switch语法

	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// defer 的使用 ，延迟调用
	defer fmt.Println("end world defer")

	fmt.Println("hello")

	fmt.Println("for循环的值：", sum, sums, kk)

	// 声明变量的常用方式

	fmt.Println(c, php, golang)

	fmt.Println("变量的初始化i的值：", i, "变量初始化j的值", j)

	fmt.Println("短变量声明方式：", k, erlang, js, python)

	// 声明布尔类型、uint64类型、复合类型
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	// 零值
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	// 常量
	fmt.Println("常量声明", sc, pi)

	// 结构体

	fmt.Println(Sex{0, 1})

	v := Sex{18, 20}
	v.X = 4
	fmt.Println(v.X)
	p1 := &v
	p1.X = 1e9
	fmt.Println(v.X)
	// 映射

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println("映射的值：", m["Bell Labs"])

	// 映射的操作  赋值，删除，检测key是否存在
	mdk := make(map[string]int)

	mdk["Answer"] = 42
	fmt.Println("The value:", mdk["Answer"])

	mdk["Answer"] = 48
	fmt.Println("The value:", mdk["Answer"])

	delete(mdk, "Answer")
	fmt.Println("The value:", mdk["Answer"])

	vmdk, ok := mdk["Answer"]
	fmt.Println("The value:", vmdk, "Present?", ok)

	// 数组

	var a [2]string
	a[0] = "Hello"
	a[1] = "Chuhemiao"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [10]int{2, 3, 5, 7, 11, 13, 15, 17, 19, 21}
	// 如果位数不够则补0
	prime := [10]int{2, 3, 5, 7, 11, 13}
	//array index 5 out of bounds [0:5]
	//primeNum := [5]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Println(prime)
	//fmt.Println(primeNum)

	// 切片
	primeSlice := []int{2, 3, 5, 7, 11, 13}

	var sslice []int = primeSlice[1:4]
	// 简写从0开始
	var slicestart []int = primeSlice[:4]
	// 简写 切到最后一个元素
	var sliceend []int = primeSlice[1:]
	// [3,5,7]
	fmt.Println(sslice)
	fmt.Println(slicestart)
	fmt.Println(sliceend)

	// 截取切片使其长度为 0
	scap := primeSlice[:0]
	printSlice(scap)

	// 拓展其长度
	scap = primeSlice[:4]
	printSlice(scap)

	// 舍弃前两个值
	scap = primeSlice[2:]
	printSlice(scap)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	an := names[0:2]
	bn := names[1:3]
	fmt.Println(a, b)

	bn[0] = "XXX"
	fmt.Println(an, bn)
	fmt.Println(names)
	// 切片为nil 的情况
	var snil []int
	fmt.Println(s, len(s), cap(snil))
	if snil == nil {
		fmt.Println("nil!")
	}

	// 创建切片
	amake := make([]int, 5)
	printSliceMake("amake", amake)

	// 给切片增加元素

	var sappend []int

	printSlice(sappend)

	// 添加一个空切片
	sappend = append(sappend, 0)
	printSlice(sappend)

	// 这个切片会按需增长
	sappend = append(sappend, 1)
	printSlice(sappend)

	// 可以一次性添加多个元素
	sappend = append(sappend, 2, 3, 4)
	printSlice(sappend)

	//  循环切片中的值
	for i, v := range sappend {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// 省略下标的模式 ，如果不需要取下标，则可使用 _ 下划线直接过滤

	for _, v := range sappend {
		fmt.Println("当前需要的值：", v)
	}

	// 函数

	fmt.Println("加法函数：", add(100, 22))
	// 函数之 方法

	vv := VertexFunc{3, 4}
	fmt.Println("函数之方法：", vv.Abs())

	// 函数之变长参数

	xargs := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", xargs)
	slice := []int{7, 9, 3, 5, 1}
	xargs = min(slice...)
	fmt.Printf("The minimum in the slice is: %d", xargs)

	//递归实现斐波那契数列
	var ii int
	for ii = 0; ii < 10; ii++ {
		fmt.Printf("%d\t", fibonacci(ii))
	}

	// 类型转换

	var sumType int = 17
	var count int = 5
	var mean float32

	mean = float32(sumType) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)

	fmt.Println()

}

// 函数的定义
/*
	1.必须以func 开头，然后是函数名（传入值类型）返回类型{
		函数体，必须有返回值
	}
*/
func add(a int, b int) int {
	return a + b
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceMake(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// 变长参数

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

// 递归实现斐波那契数列

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
