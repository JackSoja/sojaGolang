package main

import "fmt"





/*
结构体

结构体定义需要使用 type 和 struct 语句。
struct 语句定义一个新的数据类型，结构体有中一个或多个成员。type 语句设定了结构体的名称。
 */
type Car struct {
	id int
}



/* 定义接口 */
type ICar interface {
	beep()
	drive(dirver string)
}
/* 函数方法定义

Go 语言中同时有函数和方法。
一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
所有给定类型的方法属于该类型的方法集。
s
func (variable_name variable_data_type) function_name() (return_types){
	// 函数体
}
*/

/* 实现接口方法 */
func (car Car) beep() {
	// %v  相应值的默认格式。 Printf("%v", people) {zhangsan}，
	fmt.Printf("Car %v beeps, the %v time\n", car.id)

}
/* 实现接口方法 */
func (car Car) drive(driver string)  {
	fmt.Printf("%v drives car %v", driver, car.id)
}

func Method()  {
	c1 := Car{id:1001}
	c1.beep()
}

func Interface()  {
	var c2 = Car{id:1002}
	// 接口赋值
	var ic2 ICar = c2
	ic2.beep()
	ic2.drive("bob")
}

func main()  {
	Method()
	Interface()
}