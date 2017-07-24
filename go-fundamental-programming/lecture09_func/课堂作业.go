package main

import "fmt"

func main() {

	var fs = [4]func(){}	// 定义一个长度为 4 的数组，数据类型为 func


	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)	// 倒叙执行： 打印 defer i = 1
		defer func() { fmt.Println("defer_closure i = ", i) }()  //倒叙执行： 声明一个匿名数，并打印 defer_closure
		fs[i] = func() { fmt.Println("closure i = ", i) }  // 修改 fs 数组的元素。

	}

	for _, f := range fs {
		f()   // 调用 fs 中的元素。
	}

	/* 猜测结果：
	先执行所有 f(), 然后在倒叙循环执行 defer

	closure i = 1
	closure i = 2
	closure i = 3
	closure i = 4
	defer_closure = 4
	defer i = 4
	defer_closure = 3
	defer i = 3
	defer_closure = 2
	defer i = 2
	defer_closure = 1
	defer i = 1
	*/


	/* 实际结果

	closure i =  4
	closure i =  4
	closure i =  4
	closure i =  4
	defer_closure i =  4
	defer i =  3
	defer_closure i =  4
	defer i =  2
	defer_closure i =  4
	defer i =  1
	defer_closure i =  4
	defer i =  0
	 */
}


