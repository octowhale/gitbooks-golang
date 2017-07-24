package main

import (
	"fmt"
)

func main() {

	/* 格式:  map[KeyType]ValueType */

	// 01
	var m1 map[int]string
	m1 = map[int]string{}
	fmt.Println(m1)

	// 02
	var m2 map[int]string
	m2 = make(map[int]string)
	fmt.Println(m2)

	// 03
	var m3 map[int]string = map[int]string{}
	var m4 map[int]string = make(map[int]string)

	fmt.Println(m3, m4)

	// 05
	m5 := map[int]string{}
	m6 := make(map[int]string)

	fmt.Println(m5, m6)
}
