package main

import (
	"fmt"
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


func main(){

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

	var kk  bool

	if(sums > 4850){
		kk =  true
	}else{
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


	fmt.Println( "for循环的值：",sum,sums,kk)



	// 声明变量的常用方式

	fmt.Println( c, php, golang)

	fmt.Println("变量的初始化i的值：",i,"变量初始化j的值",j)

	fmt.Println("短变量声明方式：",k, erlang, js, python)

	// 声明布尔类型、uint64类型、复合类型
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	// 零值
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	// 常量
	fmt.Println("常量声明",sc,pi)

	// 结构体

	fmt.Println(Sex{0, 1})

	v := Sex{18, 20}
	v.X = 4
	fmt.Println(v.X)
	p1 := &v
	p1.X = 1e9
	fmt.Println(v.X)



}