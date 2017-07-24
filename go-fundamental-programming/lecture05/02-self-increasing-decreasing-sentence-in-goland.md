# 递增递减语句


在 GO 中， `++` 和 `--` 是作为语句而不是表达式。 
即自增自减语句不能在等号的右边

```go
package main

import "fmt"

func main() {

	a := 10
	a++
	//b := a++		/* 这一行是错误的 */
	b := a

	fmt.Println(a)
	fmt.Println(b)

}

```

