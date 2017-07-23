
# 变量的类型转换

+ 在 go 中，不存在隐式转换，所有类型的转换必须显示声明。
+ 转换只能发生在两种互相兼容的类型之间。
+ 类型转换的格式
  + `ValueA := TypeOfValueA(ValueB)`
    + 当 ValueA 之前并未被声明，必须使用『冒号』
  + `ValueA = TypeOfValueA(ValueB)` 
    + 当 ValueA 之前已经声明过的时候，可以不使用『冒号』
 
```go
package main

import (
	"fmt"
)

func main() {

	var v_int int = 1
	var v_float float32 = 1.1

	fmt.Println(" float32(v_int) = ", float32(v_int))

	/* 使用冒号 */
	v_float_tr := float32(v_int)
	fmt.Println(" v_float_tr = ", v_float_tr)

	/*  不使用冒号 */
	var v_int_tr int
	v_int_tr = int(v_float)
	fmt.Println(" v_int_tr = ", v_int_tr)

	/* 不能转换 */

	//v_bool := bool(v_int)  /* cannot convert v_int (type int) to type bool */
	//fmt.Println("v_bool = ", v_bool)
}

/* 
 float32(v_int) =  1
 v_float_tr =  1
 v_int_tr =  1

*/
```