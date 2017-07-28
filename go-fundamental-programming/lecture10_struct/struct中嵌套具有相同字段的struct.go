// package lecture10_struct

package main

import (
	"fmt"
)

type teacher struct {
	salary int
	gender int
}

type female struct {
	gender string
}

type master struct {
	good string
}

type female_teacher struct {
	name string
	teacher
	female
	master
}

func main() {

	fmt.Println("")

	ft := female_teacher{
		name: "joe",
	}

	ft.salary = 1000
	ft.teacher.gender = 1
	ft.female.gender = "female"
	//ft.gender=1  /* ambiguous selector ft.gender : 歧义字段*/
	ft.good = "ok "
	fmt.Println(ft)

}

/*
+ 嵌套 struct
  + 嵌套的 struct 中如果没有歧义字段，可以直接引用。 例如 ft.salary
  + 如果有歧义，需要写明上层 struct 名称。例如 ft.gender
*/