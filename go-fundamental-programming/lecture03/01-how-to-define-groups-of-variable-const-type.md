# 批量声明类型、常量、变量

与之前说的到[包批量引入](../lecture02/04-how-to-import-package-in-golang.md)一样， 在 golang 中，同样可以批量声明 类型、常量、变量等。
方法都是在『关键字后面接一个圆括号』。


## 声明常量
```go

// 声明常量
const (
    PI = 3.14
    const1 = "1"
    const2 = 2
    const3 = 'a'
)
```

## 声明全局变量
```go
// 声明全局变量与赋值
var (
    name  = "gopher"
    name1 = "1"
    name2 = 2
    name3 = 'a'
)
```

需要注意的是，在使用 `var` 关键字声明组变量的时候，只能在函数体外声明全局变量。这种方式不能在函数体内声明局部变量。


## 声明一般类型
```go
// 声明一般类型
type (
    newType int
    type1   float32
    type2   string
    type3   byte
)
```

