# switch 分支选择语句

+ 可以使用任何类型或表达式做为条件语句
+ 不需要写 `break`， 一旦符合条件自动终止。
+ 如果希望继续执行下一个 case ，则需要使用 `fallthrough` 关键字。
+ 如果所有条件都不符合，可以使用 `default` 关键字增加默认分支。
  + 如果不需要默认分支， `defualt` 可以省略。
+ 支持一个初始化表达式（可以是并行方式），右侧需要跟分号；
  + 初始化变量的作用域只在 switch 中。，如 mode_03。
+ 左大括号必须与 `switch` 在同一行，并且在条件表达式的末尾。


```go
package main

import (
	"fmt"
)

func main() {

	mode_01()
	mode_02()
	mode_03()
}

func mode_01() {

	/*
		switch 提供一个常量或表达式
		case 后面可以是一个常量或表达式，如果该值与前面 switch 中的常量或表达式结果相同，则结果为真，执行 case 分支中的语句
	*/
	a := 1
	switch a - 1 { /* switch 后面跟的是一个表达式 */
	case 1 - 1: /* case 后面跟的是一个表达式 */
		fmt.Println("a=0")
	case 1: /* case 后面跟的是一个常量 */
		fmt.Println("a=1")
	}

	switch a { /* switch 后面跟的是一个常量 */
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	}
	fmt.Println(a)

	switch a > 1 {   /* switch 后面跟的是一个条件表达式 */
	case true:
		fmt.Println("condition is : true")
	case false:
		fmt.Println("condition is : false")
	}
}

func mode_02() {

	/*
		switch 不提供标准值。
		case 后面可以是一个条件表达式，如果条件表达式为真，则会执行 case 分支中的语句。
	*/
	a := 1
	switch {
	case a >= 0: /* case 后面跟的是一个 条件表达式 */
		fmt.Println("a>=0")
		fallthrough // 使用 fallthrough ，继续后面的判断。
	case a >= 1:
		fmt.Println("a>=1")
	}

}

func mode_03() {
	a := -1
	switch a := 1; { /* 如果不写分号，报错提示：syntax error: a := 1 used as value */
	case a >= 0:
		fmt.Println("inside a>=0")
		fallthrough
	case a >= 1:
		fmt.Println("inside a>=1")
	default:
		fmt.Println("default: nothing matched")

	}

	fmt.Println("Ouside a = ", a)
}

```
