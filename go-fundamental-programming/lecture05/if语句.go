// package main

package main

import (
	"fmt"
)

func main() {
	fmt.Println("")

	a := 1
	if a := 2; a > 1 {
		fmt.Println("Inner: ", a)
	}
	fmt.Println("Outer: ", a)

	/*
	Inner:  2
	Outer:  1
	*/

	if_else_func()
}

func if_else_func() {

	if a := -1; a > 0 {
		fmt.Println(" a > 0")
	} else if a == 0 {
		fmt.Println("a == 0")
	} else {
		fmt.Println("a < 0 ")
	}
}

/*

+ 在 GO 语言中，条件表达式没有括号。
+ 左大括号『必须』在 条件表达式末尾。
+ 支持在 if 条件语句中初始化表达式。在 if 中初始化的变量，仅在 if 作用域中有效。
+ 如果 if 初始化表达式中的变量名称已经在上层使用，那么在 if 的作用域中，上层的同名变量将会被隐藏。
+ `else` 左边必须为『右大括号』
+ 在 GO 语言中，不支持关键字 `elif`


*/
