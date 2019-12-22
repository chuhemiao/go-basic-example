package main

import "fmt"

// 定义接口Phone，里面仅有call()一个方法
type Phone interface {
	call()
}

type HwPhone struct {
}

func (hwPhone HwPhone) call() {
	fmt.Println("I am HwPhone, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

// nil接口值

type I interface {
	M()
}

// 定义一个最常用的接口 Stringer

type Person struct {
	Name string
	Age  int
}

func main() {
	// 定义一个phone的变量
	var phone Phone

	phone = new(HwPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	// 空接口

	var inil interface{}

	describenil(inil)

	inil = 42
	describenil(inil)

	inil = "hello"
	describenil(inil)

	// 类型断言

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)
	//  判断 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
	// 报错(panic，interface conversion: interface {} is string, not float64)
	//f = i.(float64)
	//fmt.Println(f)

	//类型选择 按照接口接受的类型，然后case判断当前类型，在返回对应情况
	do(21)
	do("hello")
	do(true)
	// 定义一个最常用的接口 Stringer,Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。

	/*
		type Stringer interface {
			String() string
		}
	*/

	a := Person{"kk ", 42}
	z := Person{"chuhe miao", 9001}
	fmt.Println(a, z)

	// nil接口值
	//var i I
	//describe(i)
	//i.M()

}

//定义一个最常用的接口 Stringer

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describenil(inil interface{}) {
	fmt.Printf("(%v, %T)\n", inil, inil)
}

// 类型选择

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
