// package main

package main

import (
	"fmt"
)

func main() {

	break_func()
	goto_func()
	continue_func()
}

func break_func() {
	fmt.Println("in break func")
LEBEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LEBEL1
			}
		}
	}
	fmt.Println("OK")
}

func goto_func() {
	fmt.Println("in goto func")
	a := 1
LEBEL1:
	for i := 1; i < 10; i++ {
		a++
		//fmt.Println("a = ", a)
		fmt.Println("i = ", i)
		if a == 4 {
			break
		} else {
			goto LEBEL1
		}
	}

	/* 直接结果： i 的置每次都被重置了
	in goto func
	i =  1
	i =  1
	i =  1

	*/
}

func continue_func() {
	fmt.Println("in continue func")
	a := 1
LEBEL1:
	for i := 1; i < 10; i++ {
		a++
		//fmt.Println("a = ", a)
		fmt.Println("i = ", i)
		if a == 4 {
			//continue LEBEL2
			break
		} else {
			continue LEBEL1
		}
	}

	/* 直接结果： i 的值，没有被重置，依旧依次递增
	in continue func mode
	i =  1
	i =  2
	i =  3
	*/
}

/*
在 go 语言中，提供了三个跳转关键字以及一个『标签』功能。
+ goto: 调整执行位置
+ break: 退出循环到标签，并跳过剩余循环
+ continue: 退出循环到标签，并执行下一次循环

+ 标签名区分大小写
  + 标签声明后，若不使用会造成编译错误
+ 三个语法都可以配合标签使用
  + 如果没有指定标签， `break` 和 `continue` 只会跳出当前循环
  + 如果指定了标签，`break` 、 `continue` 和 `goto` 会跳转到标签位置
+ `break` 和 `continue` 配合标签可用于多层循环的跳出
  + 跳出循环后，保留循环的当前状态。 `continue` 可继续之前没有执行完的循环
+ `goto` 是调整执行位置，与其他 2 个语句配合标签的结果并不相同
  + 重新运行到循环时，循环重置。标签之后执行过的内容还会重新被执行
  + 在循环中使用 goto ，建议将跳出标签放在循环外面

*/
