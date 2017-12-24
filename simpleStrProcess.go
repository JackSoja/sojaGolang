package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	/* 字符串基本操作--strings */
	str := "sojaGolang"

	// 是否包含 : true false
	fmt.Println(strings.Contains(str, "soja"), strings.Contains(str, "nana"))

	// 字符串长度 : 10
	fmt.Println(len(str))

	// 字符在字符串中的位置，从0开始，不存在返回-1 : 4 -1
	fmt.Println(strings.Index(str, "G"), strings.Index(str, "i"))

	// 判断字符串是否以 xx 开头 : true false
	fmt.Println(strings.HasPrefix(str, "soja"), strings.HasPrefix(str, "Golang"))

	// 判断字符串是否以 xx 结尾 : false true
	fmt.Println(strings.HasSuffix(str, "soja"), strings.HasSuffix(str, "Golang"))

	// 判断2个字符串大小，相等0，左边大于右边-1，其他1 : 0 1 -1
	str2 := "nanana"
	fmt.Println(strings.Compare(str, str),strings.Compare(str, str2),strings.Compare(str2, str))

	// 分割字符串 : 返回数组 [a b c d e f g f]
	str3 := "a,b,c,d,e,f,g,f"
	fmt.Println(strings.Split(str3, ","))

	// 拼接字符串 : a-b-c-d-e-f
	str4 := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(strings.Join(str4, "-"))

	// 去除字符串两端空格 : 13 5 world
	str5 := "    world    "
	str6 := strings.Trim(str5, " ")
	fmt.Println(len(str5), len(str6), str6)

	// 大小写转换 : iphone IPHONE
	str7 := "iPhone"
	fmt.Println(strings.ToLower(str7), strings.ToUpper(str7))

	// 字符串替换 : iphone iPhone
	str8 := "iphone"
	str9 := strings.Replace(str8, "ip", "iP", 2)
	fmt.Println(str8, str9)

	/* 字符串类型转换 strconv */
	// 整型 转 字符串

	int1 := 12
	str10 := strconv.Itoa(int1)
	fmt.Println(str10)

	// 字符串 转 整型
	int2, err := strconv.Atoi(str10)
	if err == nil {
		fmt.Println(int2)
	} else {
		fmt.Println(err)
	}

	// bool 转 字符串
	bool1 := true
	str11 := strconv.FormatBool(bool1)
	fmt.Println(str11)

	// 字符串 转 bool
	bool2, err := strconv.ParseBool(str11)
	if err == nil {
		fmt.Println(bool2)
	} else {
		fmt.Println(err)
	}

	// FormatFloat 将浮点数 f 转换为字符串值
	// f：要转换的浮点数
	// fmt：格式标记（b、e、E、f、g、G）
	// prec：精度（数字部分的长度，不包括指数部分）
	// bitSize：指定浮点类型（32:float32、64:float64）
	//
	// 格式标记：
	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)
	//
	// 如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
	// 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）
	float1 := 3.14159265
	str12 := strconv.FormatFloat(float1, 'f',5,64)
	fmt.Println(str12)


}