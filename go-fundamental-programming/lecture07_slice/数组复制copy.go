package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{1, 2, 3}

	fmt.Println(s1, s2)

	s3 := s1

	fmt.Println(s3)
	fmt.Printf("%p\n%p\n", s1, s3)

	s4 := s1[:]
	fmt.Printf("%p\n%p\n", s1, s4)

	/* 结果
	都指向同一个内存地址
	[1 2 3 4 5 6] [1 2 3]
	[1 2 3 4 5 6]
	0xc042037f50
	0xc042037f50
	0xc042037f50
	0xc042037f50

	*/

}
