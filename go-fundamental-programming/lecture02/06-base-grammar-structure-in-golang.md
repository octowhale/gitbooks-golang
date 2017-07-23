
# 基本代码结构示例

```go
// 当前程序的包名
// 可以与文件名不一致
// package 必须放在『非注释或非空行的第一行』
//
// package 包名
package main

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

