# for 循环语句
+ 在 GO 语言中，循环语句只有 `for` 关键字。但是 `for` 的不同的条件形式实现了 `while` 循环。
+ 左大括号必须与 `for` 关键字同一行，并且在『条件表达式』的后面。
  + 如果省略条件表达式，则循环条件『恒真』。如 `while_mode_without_condition`。
  + 『条件表达式』可以仅为一个判断语句。如 `while_mode_with_condition`。
  + `for` 循环还有一种经典的表达式。三个语句之间使用 `;` 分割。
+ 尽量不要在条件表达式中放入『需要计算的表达式』，否则循环的每次判断都会去计算该表达式的结果。这样会浪费性能。


```go
// package main

package main

import (
	"fmt"
)

func main() {

	while_mode_without_condition()
	while_mode_with_condition()
	classic_mode()
}

func while_mode_without_condition() {

	a := 1
	/* 这里 for 后面不接任何条件表达式，恒真*/
	for {
		a++
		if a > 3 {
			break
		}
		fmt.Println(a)

	}
	fmt.Println("while_mode_without_condition: over")
}

func while_mode_with_condition() {

	a := 1
	for a <= 3 {
		a++
		fmt.Println(a)
	}
	fmt.Println("while_mode_with_condition: Over")

}

func classic_mode() {
	a := 1
	for i := 0; i < 3; i++ {
		a++
		fmt.Println(a)
	}
	fmt.Print("Classic mode: Over")
}

/*
*/

```
