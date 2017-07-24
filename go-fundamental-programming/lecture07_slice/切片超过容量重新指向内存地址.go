package main

import "fmt"

func main() {

	s1 := make([]int, 3, 6)
	fmt.Printf("%p\n", s1, )

	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 4, 5, 6)
	fmt.Printf("%p\n", s1)

	/* 结果

	0xc042037f50
	0xc042037f50
	0xc042032060

	*/

}
