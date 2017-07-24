package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(a)

	for i, v := range a {
		fmt.Printf("%d -> %d\n", i, v)
	}

	/*
	注意： i 为 array 中的位置， v 为对应位置的值的一个拷贝。
	对 v 的任何操作都不会影响到 array。
	如果需要对 array 进行修改，可以使用 a[i] 的方式
	*/

}
