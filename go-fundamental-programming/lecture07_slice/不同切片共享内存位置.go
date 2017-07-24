package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5}
	s1 := a[2:5]
	s2 := a[1:3]

	fmt.Println(s1, s2, a)

	//s1[0] = 9	  /* 指向内存地址未变，修改元素会影响其他切片或原数组 */

	fmt.Println(s1, s2, a)
	//fmt.Println(len(s1), cap(s1))
	//fmt.Printf("%p\n",s1)

	s1 = append(s1, 7, 8, 9, 10, 11, 12, 13, 14)
	//fmt.Println("%p", s1)

	//fmt.Println(len(s1), cap(s1))
	//fmt.Printf("%p\n",s1)

	s1[0] = 9	/* 指向内存地址改变，修改元素不会影响其他切片或原数组 */
	fmt.Println(s1, s2, a)

	/* 结果
	[3 4 5] [2 3] [1 2 3 4 5]
	[9 4 5] [2 9] [1 2 9 4 5]

	改变 slice 的元素位置的时候，可能
	  + 原始数组对应的位置的元素也会改变。
	  + 其他引用该位置的切片也会受到影响。
	这种情况的出现，取决于切片 s1 在改变元素的时候是否超过切片容量 `cap()`而指向了新分配的内存地址， 。如果超过，会重新分配内存地址，则不会发生上诉情况。
	*/

}
