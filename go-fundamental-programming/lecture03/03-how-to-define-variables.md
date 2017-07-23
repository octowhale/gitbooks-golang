# 变量的声明与赋值

## 单个变量的声明与赋值

变量的声明有多种方式

+ 变量的类型声明格式: `var 变量名称 变量类型` , `var v_str string`
  + 变量的赋值格式: `变量名称 = 表达式` ,  `v_str = "hello, my love"`
  + 注意：这种方式赋值只能在函数体内进行
+ 变量的类型声明与赋值: `var 变量名称 [变量类型] = 表达式`
  + 这种情况下， go 会根据声明语句，自动为变量设置一个类型。这种类型可能会与你的预期不同。
  + 例如， `var v_int = 3 `。 程序可能给 v_int 分配的是一个 `int8` 类型，但程序可能之后会用到 `int32`。
+ 省略 `var` 的声明方式: `变量名 := 表达式` , ` v_str := "hello, my love"
  + 注意：该声明与赋值方式只能在函数体内进行


```go

package main

import "fmt"

// 声明与赋值
var var1 int
/* var1 = 123  单独赋值只放在函数体内*/

// 声明并赋值
var var2 int = 456

// 省略类型的声明与赋值
var var3 = "hello, my love"

//var4 := 789  /* 省略 var 声明与赋值只能在函数体内 */

func main() {

	var1 = 123

	// 省略 var 的声明与赋值
	var4 := 789  /* 省略 var 声明与赋值只能在函数体内 */

	fmt.Println("var1 = ", var1)
	fmt.Println("var2 = ", var2)
	fmt.Println("var3 = ", var3)
	fmt.Println("var4 = ", var4)
}

/* 结果

var1 =  123
var2 =  456
var3 =  hello, my love
var4 =  789

```


## 多个变量的声明与赋值

+ 所有变量都可以使用 `var()` 组的方式进行声明
+ 所有变量都可以使用并行方式进行声明
+ 所有变量都可以使用类型推断
+ 全局变量的声明不可以省略 `var` 关键字
+ 局部变量可以省略 `var` 关键字进行声明


```go
package main

import "fmt"

var a, b, c, d int = 1, 2, 3, 4

var (
	//x, y, z := "x", "y", "z"		/* 在 var() 体内，不能在使用 := 了。因为 : 就是为了代替 var 的 */
	x, y, z = "x", "y", "z"
)

func main() {

	//var a, b, c, d int = 1, 2, 3, 4
	var a, b, c, d int
	a, b, c, d = 5, 6, 7, 8

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)

	var e, f, g, h = 1, 'a', "b", true

	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
	fmt.Println("g = ", g)
	fmt.Println("h = ", h)


	i, _, k, l := "a", false, a+1, 10	  /* _ 下划线， 占位符，赋值省略 */
	fmt.Println("i = ", i)
	//fmt.Println("_ = ", _)   /* 不能直接调用 _ 作为变量  -> cannot use _ as value */
	fmt.Println("k = ", k)
	fmt.Println("l = ", l)


	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	fmt.Println("z = ", z)

}


/* 结果

a =  5
b =  6
c =  7
d =  8
e =  1
f =  97
g =  b
h =  true
i =  a
k =  6
l =  10
x =  x
y =  y
z =  z

*/

```
