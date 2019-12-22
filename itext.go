package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "ssssppllssdsdljjklljsd"
	// 字符串长度
	l1 := len([]rune(str))
	l2 := bytes.Count([]byte(str), nil) - 1
	l3 := strings.Count(str, "") - 1
	l4 := utf8.RuneCountInString(str)
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println(l3)
	fmt.Println(l4)

	// 字符串中是否存在某个字符  返回值 true或false
	fmt.Println(strings.ContainsAny(str, "i"))
	// 字符串出现的次数
	fmt.Println(strings.Count(str, "ss"))
	// 字符串分割
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))

	fmt.Printf("%q\n", strings.SplitN("foo,bar,baz", ",", 2))
	// 字符串以某某开头
	sstr := strings.HasPrefix(str, "ss")
	fmt.Println(sstr)
	// 字符串以某某结尾
	send := strings.HasSuffix(str, "dddd")
	fmt.Println(send)
	// 字符串替换
	/*
		用 new 替换 s 中的 old，一共替换 n 个。
		如果 n < 0，则不限制替换次数，即全部替换
		func Replace(s, old, new string, n int) string
	*/

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2)) //oinky oinky oink

	/*
		字符串大小写函数

		1、func Title(s string) string
		将字符串s每个单词首字母大写返回

		2、func ToLower(s string) string
		将字符串s转换成小写返回

		3、func ToLowerSpecial(_case unicode.SpecialCase, s string) string
		将字符串s中所有字符按_case指定的映射转换成小写返回

		4、func ToTitle(s string) string
		将字符串s转换成大写返回

		5、func ToTitleSpecial(_case unicode.SpecialCase, s string) string
		将字符串s中所有字符按_case指定的映射转换成大写返回

		6、func ToUpper(s string) string
		将字符串s转换成大写返回

		7、func ToUpperSpecial(_case unicode.SpecialCase, s string) string
		将字符串s中所有字符按_case指定的映射转换成大写返回
	*/
	// 转换成大写返回
	TestToTitle()
	// 去除字符串左右的无效字符

	fmt.Println(strings.Trim("sssdklklsd !!", "!"))

	// 类型之间转换

	// 字符串转整型
	fmt.Println(strconv.ParseInt("-12", 10, 0)) // -12 <nil>

	// 进制转换

	n, err := strconv.ParseInt("120", 10, 8)

	if err != nil {
		panic("进制转换失败")
	}
	fmt.Println("进制转换之后的数:", n)

	// 将字符串转为浮点数

	v := "3.14159267823535"

	if s, err := strconv.ParseFloat(v, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s) // float64, 3.1415927410125732
	}

	if s, err := strconv.ParseFloat(v, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s) // float64, 3.14159267823535
	}

	// 字符串转布尔值

	fmt.Println(strconv.ParseBool("1")) // true

}

func TestToTitle() {
	fmt.Println(strings.ToTitle("chuhe miao"))
}
