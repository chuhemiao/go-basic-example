package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

// 通过反射调用方法
func (u User) HelloName(name string) {
	fmt.Println("Hello", name, ", my name is", u.Name)
}

func main() {

	u := User{1, "ok", 18}
	Info(u)
	// 通过反射方法的调用
	uname := User{1, "ok", 18}
	v := reflect.ValueOf(uname)
	mv := v.MethodByName("HelloName")

	args := []reflect.Value{reflect.ValueOf("chuhemiao")}
	mv.Call(args)

}

// 传递一个空接口
func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields:", v)

	// 获取方法字段
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Println("%6s: %v = %v\n", f.Name, f.Type, val)

	}

	// 获取方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println("%s: %v\n", m.Name, m.Type)
	}

}
