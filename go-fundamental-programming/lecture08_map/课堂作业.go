package main

import "fmt"

func main() {

	//fmt.Println("")

	m1 := map[int]string{0: "j", 1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f", 7: "g", 8: "h"}
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k

	}

	fmt.Println("m1 ", m1)
	fmt.Println("m2 ", m2)
}
