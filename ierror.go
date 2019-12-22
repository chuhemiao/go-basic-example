package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// error 返回两个参数，判断err是否为nil，并处理错误信息
	open, err := os.OpenFile("./opentest.txt", os.O_RDONLY, 0644)

	if err != nil {
		errors.New("math: square root of negative number")
	}

	fmt.Println("函数之方法：", open)

}
