# go 语言中的注意事项

+ Go 程序是通过 package 进行组织的。（与 python 类似）
+ 只有 package 为成为 main 的包可以包含 `main()` 函数。否则在编译的时候，会报错提示不止一个 main 函数。
+ 一个可执行程序**有且仅有**一个 main 包， main 包中**有且仅有**一个 main 函数。


+ 通过 `import` 关键字来导入其他非 main 包
+ 通过 `const` 关键字来定义常量。
+ 通过 `var` 关键字进行变量的声明或赋值。在函数体外声明为全局变量；在函数体内声明为局部变量。
+ 通过 `type` 关键字进行结构( `struct` )或接口( `interface` ) 的声明。
+ 通过 `func` 关键字进行函数的声明。


## 基本代码结构示例

```go
// 当前程序的包名
// 可以与文件名不一致
// package 必须放在『非注释或非空行的第一行』
//
// package 包名
pakcage main

// 导入其他包
// 导入的包名需要使用『双引号』括起来
import "fmt"

// 常量的声明
const PI=3.14

// 全局变量的声明
var author="octowhale"

// 一般类型的声明
type newType int

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

// 由 main 函数作为程序入口点启动
func main(){
    fmt.Println("Hello world！ 你好，世界！")
}

```

